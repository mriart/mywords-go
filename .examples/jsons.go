package main

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

type Card struct {
	Cat     string
	CatEx   string
	Eng     string
	EngPron string
	EngEx   string
}

func main() {

	// Create a client with CLOUDANT environment vars
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{})
	if err != nil {
		panic(err)
	}

	// Get documents and put in a json
	dbName := "mywords2"

	postAllDocsOptions := client.NewPostAllDocsOptions(dbName)
	postAllDocsOptions.SetIncludeDocs(true)
	postAllDocsOptions.SetLimit(2)

	/*
		streamAllDocs, _, err := client.PostAllDocsAsStream(postAllDocsOptions)
		if err != nil {
			panic(err)
		}
		m := make(map[string]interface{})
		err2 := json.Unmarshal(streamAllDocs, &m)
		if err2 != nil {
			panic(err2)
		}
		fmt.Println(m["total_rows"])


			var allDocs map[string]interface{}
			if err := json.NewDecoder(streamAllDocs).Decode(&allDocs); err != nil {
				panic(err)
			}
			//fmt.Printf("\nJson: %+v", settings)

			for k, v := range allDocs {
				fmt.Println("k:", k)
				fmt.Println("v:", v)
			}
	*/

	allDocsResult, _, _ := client.PostAllDocs(postAllDocsOptions)
	b, _ := json.MarshalIndent(allDocsResult, "", "  ")

	m := make(map[string]interface{})
	err2 := json.Unmarshal(b, &m)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(m["total_rows"])

	fmt.Println("\n\n", string(b))

	//fmt.Println(allDocsResult.TotalRows) no va, torna binari

	/* 1 document
	getDocumentOptions := client.NewGetDocumentOptions(dbName, "5b930a291955955f2430b44c289ba563")

	document, _, err := client.GetDocument(getDocumentOptions)
	if err != nil {
		panic(err)
	}

	b, _ = json.MarshalIndent(document, "", "  ")
	fmt.Println(string(b))

	var c Card
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
	*/
}
