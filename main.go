package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

// The global variable for the Cloudant DB.
var dbName string = "mywords"

// Mapes the JSON returned by Cloudant.
type AllDocs struct {
	Total_rows int
	Rows       []Row
}

// Mapes a row of the AllDocs struct.
type Row struct {
	Id  string
	Doc Doc
}

// Mapes the fields of a Row, the content.
type Doc struct {
	Cat     string
	CatEx   string
	Eng     string
	EngPron string
	EngEx   string
}

// A hello function, to review the server is up and running.
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// Function to review the server is up and running, and well connected to Cloudant.
// It requires the env variables CLOUDANT_URL and CLOUDANT_APIKEY.
func serverInfo(w http.ResponseWriter, req *http.Request) {
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

// Function that shuffles the positions of a slice.
// It is used to randomize the output of the collection of words.
func shuffle(in []string) {
	for i := range in {
		rand.Seed(time.Now().UnixMicro())
		j := rand.Intn(i + 1)
		in[i], in[j] = in[j], in[i]
	}
}

// It prepares the collection send to the device. This sent collection is formed:
// -the first 3 words (latest introduced), for remembering
// -the shuffled collection (including last 3)
// -the END card
func mix(in []string) []string {
	pre3in := []string{in[0], in[1], in[2]}
	shuffle(in)
	in = append(pre3in, in...)
	endCard := `FI;Fi, felicitats.;END;end;The end, congrats.`
	in = append(in, endCard)
	return in
}

func myWords(w http.ResponseWriter, req *http.Request) {
	// Get the number of docs to recover from Cloudant.
	// This is specified in argument ?n= in url. By default it is 50.
	numDocsURL := req.URL.Query().Get("n")
	numDocs, err := strconv.Atoi(numDocsURL)
	if err != nil {
		numDocs = 50
		//panic(err)
	}
	//Minimum 3 docs.
	if numDocs < 3 {
		numDocs = 3
	}
	//fmt.Println(numDocs)

	// Create a client with CLOUDANT environment vars.
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{})
	if err != nil {
		fmt.Println("Error NewCloudantV1:")
		panic(err)
	}

	// Get documents and put in a json/byte array.
	postAllDocsOptions := client.NewPostAllDocsOptions(dbName)
	postAllDocsOptions.SetIncludeDocs(true)
	postAllDocsOptions.SetLimit(int64(numDocs))
	postAllDocsOptions.SetDescending(true)
	allDocsResult, _, err := client.PostAllDocs(postAllDocsOptions)
	if err != nil {
		fmt.Println("Error postAllDocs:")
		panic(err)
	}

	b, err := json.MarshalIndent(allDocsResult, "", "  ")
	if err != nil {
		fmt.Println("Error Marshal:")
		panic(err)
	}

	// Place the json object into struct.
	u := AllDocs{}
	err = json.Unmarshal(b, &u)
	if err != nil {
		fmt.Println("Error Unmarshal:")
		panic(err)
	}

	collection := make([]string, 0)
	for _, v := range u.Rows {
		collection = append(collection, v.Doc.Cat+";"+v.Doc.CatEx+";"+v.Doc.Eng+";"+v.Doc.EngPron+";"+v.Doc.EngEx)
	}
	//fmt.Println(collection)

	collection = mix(collection)
	//fmt.Println(collection)

	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, collection)
}

// For introducing new words. It shows a form with the fields for the introduction.
func postWord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Create a client with CLOUDANT environment vars
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{})
	if err != nil {
		fmt.Println("Error NewCloudantV1:")
		panic(err)
	}
	eventDoc := cloudantv1.Document{}
	eventDoc.SetProperty("Cat", r.Form["Cat"][0])
	eventDoc.SetProperty("CatEx", r.Form["CatEx"][0])
	eventDoc.SetProperty("Eng", r.Form["Eng"][0])
	eventDoc.SetProperty("EngPron", r.Form["EngPron"][0])
	eventDoc.SetProperty("EngEx", r.Form["EngEx"][0])
	putDocumentOptions := client.NewPutDocumentOptions(dbName, fmt.Sprint(time.Now().UnixMicro()))
	putDocumentOptions.SetDocument(&eventDoc)

	documentResult, _, err := client.PutDocument(putDocumentOptions)
	if err != nil {
		fmt.Println("Error PutDocument:")
		panic(err)
	}
	b, _ := json.MarshalIndent(documentResult, "", "  ")
	fmt.Println(string(b))

	http.Redirect(w, r, "/mywords", http.StatusFound)
}

// Main. Serves port 8090 for regular functioning.
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/serverinfo", serverInfo)
	http.HandleFunc("/mywords", myWords)
	http.HandleFunc("/postword", postWord)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	http.ListenAndServe(":8090", nil)
}
