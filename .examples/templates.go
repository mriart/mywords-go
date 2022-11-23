package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Card struct {
	Cat     string
	CatEx   string
	Eng     string
	EngPron string
	EngEx   string
}

var c1 Card = Card{Cat: "hola", CatEx: "bon dia", Eng: "hi", EngPron: "jai", EngEx: "good morning"}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func templates(w http.ResponseWriter, req *http.Request) {

	c2 := Card{"hola1", "bon dia1", "hi1", "jai1", "good morning1"}

	fmt.Println(c1)
	fmt.Println(c2)

	t1 := template.New("t1")
	t1, err := t1.Parse("My Value is {{.Cat}}\n")
	if err != nil {
		panic(err)
	}
	t1.Execute(os.Stdout, c2)

	t2 := template.Must(template.ParseFiles("tmpl/welcome.html"))
	t2.Execute(w, c2)
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/templates", templates)

	http.ListenAndServe(":8090", nil)
}
