package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Parents struct {
		Mother string
		Father string
	}
}

type AllDocs struct {
	Total_rows int
	Rows       []Row
}

type Row struct {
	Id  string
	Doc Doc
}

type Doc struct {
	Cat string
	Eng string
}

func templates(w http.ResponseWriter, req *http.Request) {
	encoded := `{
		"total_rows": 4,
		"rows": [
		  {
			"doc": {
			  "_id": "5b930a291955955f2430b44c289ba563",
			  "_rev": "1-2cac4672db5c4229a1051c7350f5a007",
			  "cat": "taula",
			  "eng": "table"
			},
			"id": "5b930a291955955f2430b44c289ba563",
			"key": "5b930a291955955f2430b44c289ba563",
			"value": {
			  "rev": "1-2cac4672db5c4229a1051c7350f5a007"
			}
		  },
		  {
			"doc": {
			  "_id": "73537d380f30f410784e008800a76ef3",
			  "_rev": "1-0ac9187ecaea0a7cf561db85ceedc814",
			  "cat": "casa",
			  "eng": "house"
			},
			"id": "73537d380f30f410784e008800a76ef3",
			"key": "73537d380f30f410784e008800a76ef3",
			"value": {
			  "rev": "1-0ac9187ecaea0a7cf561db85ceedc814"
			}
		  }
		]
	  }`

	// Decode the json object
	u := AllDocs{}
	err := json.Unmarshal([]byte(encoded), &u)
	if err != nil {
		panic(err)
	}

	fmt.Println(u)
	fmt.Printf("Total_rows: %d\n", u.Total_rows)
	fmt.Printf("Cat: %s\n", u.Rows[0].Doc.Cat)

	res := make([]string, 0)
	for _, v := range u.Rows {
		fmt.Println(v.Doc.Cat + ";" + v.Doc.Eng)
		res = append(res, v.Doc.Cat+";"+v.Doc.Eng)
	}
	fmt.Println(res)

	t2 := template.Must(template.ParseFiles("tmpl/welcome.html"))
	collection := make([]string, 0)
	collection = append(collection, "gosar;Com goses dir-me això?;dare;ˈdeə;How do you dare?")
	collection = append(collection, "deure diners;;owe;ˈoʊ;To need to pay or repay money to a person, bank, business, etc.")
	t2.Execute(w, collection)
}

func main() {

	http.HandleFunc("/templates", templates)

	http.ListenAndServe(":8090", nil)
}
