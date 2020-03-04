package main

import (
	"html/template"
	"os"
)

var homeTemplate *template.Template

func main() {
	var err error
	homeTemplate, err := template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Person map[string]int
	}{map[string]int{"chris": 39}}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
