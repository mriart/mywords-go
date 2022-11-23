package main

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func main() {

	// 1. Create a client with CLOUDANT environment vars
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{})
	if err != nil {
		panic(err)
	}

	// 2. Get server information
	serverInformationResult, _, err := client.GetServerInformation(client.NewGetServerInformationOptions())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server Version: %s\n", *serverInformationResult.Version)

	// 3. Get all docs from database
	dbName := "mywords2"

	postAllDocsOptions := client.NewPostAllDocsOptions(dbName)
	postAllDocsOptions.SetIncludeDocs(true)
	//postAllDocsOptions.SetStartKey("abc")
	//postAllDocsOptions.SetLimit(10)

	allDocsResult, _, err := client.PostAllDocs(postAllDocsOptions)
	if err != nil {
		panic(err)
	}

	b, _ := json.MarshalIndent(allDocsResult, "", "  ")
	fmt.Println(string(b))
}
