package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func check(e error, fu string) {
	if e != nil {
		fmt.Println(fu)
		panic(e)
	}
}

func main() {
	f, err := os.Open("mywords.txt")
	check(err, "readFile")

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{})
	check(err, "NewCloudantV1")

	var s string
	var sArr = make([]string, 5)

	for fileScanner.Scan() {
		s = fileScanner.Text()
		fmt.Println(s)
		sArr = strings.Split(s, ";")
		fmt.Println(sArr)

		eventDoc := cloudantv1.Document{}
		eventDoc.SetProperty("Cat", sArr[0])
		eventDoc.SetProperty("CatEx", sArr[1])
		eventDoc.SetProperty("Eng", sArr[2])
		eventDoc.SetProperty("EngPron", sArr[3])
		eventDoc.SetProperty("EngEx", sArr[4])
		putDocumentOptions := client.NewPutDocumentOptions("mywords", fmt.Sprint(time.Now().UnixMicro()))
		putDocumentOptions.SetDocument(&eventDoc)

		documentResult, _, err := client.PutDocument(putDocumentOptions)
		check(err, "PutDocument")
		b, _ := json.MarshalIndent(documentResult, "", "  ")
		fmt.Println(string(b))

		// Sometimes it collapses, so we need to ingest with calm.
		time.Sleep(3000 * time.Millisecond)
	}

	f.Close()
}
