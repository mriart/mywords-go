package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func cloudant(w http.ResponseWriter, req *http.Request) {

	// Create a client with CLOUDANT environment vars
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{})
	if err != nil {
		panic(err)
	}

	// Get server information
	serverInformationResult, _, err := client.GetServerInformation(client.NewGetServerInformationOptions())
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Server Version: %s\n", *serverInformationResult.Version)
}

func docs(w http.ResponseWriter, req *http.Request) {

	// Create a client with CLOUDANT environment vars
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{})
	if err != nil {
		panic(err)
	}
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
	fmt.Fprintf(w, string(b))
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/cloudant", cloudant)
	http.HandleFunc("/docs", docs)

	http.ListenAndServe(":8090", nil)
}
