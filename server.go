package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"golang.org/x/net/context"
)

func main() {
	s := gin.Default()
	var ctx = context.Background()

	client, err := elastic.NewClient()
	if err != nil {
		fmt.Printf("Error creating elasticsearch client: %v", err)
	}

	s.GET("/systems", func(c *gin.Context) {
		systems := getSystems(ctx, client)
		c.JSON(200, systems)
	})

	s.GET("/concepts", func(c *gin.Context) {
		systemString := c.Query("system")
		versionString := c.Query("version")
		searchString := c.Query("search")
		valueSet := searchConcepts(ctx, systemString, versionString, searchString, client)
		c.JSON(200, valueSet)
	})
	s.Run()
}

func getSystems(ctx context.Context, client *elastic.Client) []CodeSystem {
	searchResult, err := client.Search().
		Index("code_systems").             // search in index "code_systems"
		Query(elastic.NewMatchAllQuery()). // specify the query
		From(0).Size(200).                 // get 200 documents
		Pretty(true).                      // pretty print request and response JSON
		Do(ctx)                            // execute
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

func searchConcepts(ctx context.Context, system string, version string, search string, client *elastic.Client) ValueSet {
	codeQuery := elastic.NewQueryStringQuery("*" + search + "*")
	codeQuery.Field("conceptCode")
	codeQuery.Field("definitionText")
	searchResult, err := client.Search().
		Index("codes").
		Type(system).
		Query(codeQuery).
		From(0).Size(1000).
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
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var rawCode ESCode
			err := json.Unmarshal(*hit.Source, &rawCode)
			if err != nil {
				fmt.Printf("Deserialization of result failed: %v", err)
			}

			var code Code
			code.Code = rawCode.ConceptCode
			code.Display = rawCode.DefinitionText
			code.System = system

			vs.Expansion.Contains = append(vs.Expansion.Contains, code)

		}
	} else {
		// No hits
		fmt.Print("Found no codes\n")
	}

	return vs
}
