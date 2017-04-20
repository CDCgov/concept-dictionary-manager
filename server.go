package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"golang.org/x/net/context"
)

func main() {
	s := gin.Default()
	var ctx = context.Background()
	var client *elastic.Client
	var err error
	host := os.Getenv("ELASTIC_SEARCH_HOST")

	if host != "" {
		client, err = elastic.NewClient(elastic.SetURL("http://" + host))
	} else {
		client, err = elastic.NewClient()
	}
	if err != nil {
		fmt.Printf("Error creating elasticsearch client: %v", err)
	}

	s.GET("/systems", func(c *gin.Context) {
		limitString := c.Query("limit")
		pageString := c.Query("page")
		from, size := resolvePagination(limitString, pageString)
		systems := getSystems(ctx, from, size, client)
		c.JSON(200, systems)
	})

	s.GET("/concepts", func(c *gin.Context) {
		systemString := c.Query("system")
		versionString := c.Query("version")
		searchString := c.Query("search")
		limitString := c.Query("limit")
		pageString := c.Query("page")
		from, size := resolvePagination(limitString, pageString)
		valueSet := searchConcepts(ctx, systemString, versionString, searchString, from, size, client)
		c.JSON(200, valueSet)
	})
	s.Run()
}

func getSystems(ctx context.Context, from int, size int, client *elastic.Client) []CodeSystem {
	searchResult, err := client.Search().
		Index("code_systems").
		Query(elastic.NewMatchAllQuery()).
		From(from).Size(size).
		Pretty(true).
		Do(ctx)
	if err != nil {
		fmt.Printf("Error searching with elasticsearch client: %v", err)
	}

	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	fmt.Printf("Found a total of %d code systems\n", searchResult.Hits.TotalHits)

	var codeSystems []CodeSystem

	var ttyp ESSystem
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if rawSystem, ok := item.(ESSystem); ok {
			var codeSystem CodeSystem
			codeSystem.URL = rawSystem.SourceURL
			codeSystem.Version = rawSystem.Version
			codeSystem.Name = rawSystem.Name
			codeSystem.Status = rawSystem.Status
			codeSystem.Oid = rawSystem.Oid

			codeSystems = append(codeSystems, codeSystem)
		}
	}

	return codeSystems
}

func searchConcepts(ctx context.Context, system string, version string, search string, from int, size int, client *elastic.Client) ValueSet {
	codeQuery := elastic.NewQueryStringQuery("*" + search + "*")
	codeQuery.Field("conceptCode")
	codeQuery.Field("definitionText")

	searchResult, err := client.Search().
		Index("codes").
		Type(system).
		Query(codeQuery).
		From(from).Size(size).
		Pretty(true).
		Do(ctx)

	if err != nil {
		fmt.Printf("Error searching with elasticsearch client: %v", err)
	}

	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	var vs ValueSet

	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d codes\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			var rawCode ESCode
			err := json.Unmarshal(*hit.Source, &rawCode)
			if err != nil {
				fmt.Printf("Deserialization of result failed: %v", err)
			}

			var code Code
			code.Code = rawCode.ConceptCode
			code.Display = rawCode.DefinitionText
			code.System = rawCode.CodeSystemOid

			vs.Expansion.Contains = append(vs.Expansion.Contains, code)

		}
	} else {
		// No hits
		fmt.Print("Found no codes\n")
	}

	return vs
}

func resolvePagination(limitString string, pageString string) (int, int) {
	var limitNum int
	var pageNum int
	var err error

	if limitString != "" {
		limitNum, err = strconv.Atoi(limitString)
		if err != nil {
			fmt.Printf("Error parsing limit parameter: %v", err)
		}
	} else {
		limitNum = 1000
	}

	if pageString != "" {
		pageNum, err = strconv.Atoi(pageString)
		if err != nil {
			fmt.Printf("Error parsing page parameter: %v", err)
		}
	} else {
		pageNum = 0
	}

	from := limitNum * pageNum
	size := limitNum

	return from, size
}
