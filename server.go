package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/CDCgov/concept-dictionary-manager/models"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()

	var codeSystems []models.CodeSystem
	codeSystems = append(codeSystems, models.CodeSystem{Name: "TestSystem", Version: "1"})

	s.GET("/systems", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON(200, codeSystems)
	})

	s.GET("/concepts", conceptHandler)
	s.Run(":8080")
}

func conceptHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	systemString := c.Query("system")
	versionString := c.Query("version")
	searchString := c.Query("search")

	var dictionaryServiceURL string
	var dictionaryServicePort string
	var conceptSearchURL string

	switch systemString {
	case "TestSystem":
		switch versionString {
		case "1":
			if os.Getenv("TEST_SERVICE_HOST") != "" {
				dictionaryServiceURL = os.Getenv("TEST_SERVICE_HOST")
				conceptSearchURL = "http://" + dictionaryServiceURL + "/codes?query=" + searchString
			} else {
				dictionaryServiceURL = "localhost"
				dictionaryServicePort = "8080"
				conceptSearchURL = "http://" + dictionaryServiceURL + ":" + dictionaryServicePort + "/codes?query=" + searchString
			}
		}
	default:
		dictionaryServiceURL = ""
		dictionaryServicePort = ""
	}

	res, err := http.Get(conceptSearchURL)
	if err != nil {
		fmt.Printf("Couldn't get concepts from dictionary service: %v", err)
	}
	defer res.Body.Close()

	var codes [][]string
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unable to read http response body: %v", err)
	}

	err = json.Unmarshal(body, &codes)
	if err != nil {
		fmt.Printf("Unable to unmarshal http response body: %v", err)
	}

	var valueSet models.ValueSet

	for _, code := range codes {
		codeToAdd := models.Code{Code: code[0], Display: code[1], System: "TestSystem"}
		valueSet.Expansion.Contains = append(valueSet.Expansion.Contains, codeToAdd)
	}

	c.JSON(200, valueSet)
}
