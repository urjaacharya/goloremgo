package main

import (
	"createrandom"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var mapToFunctions = template.FuncMap{"words": createrandom.Words,
	"sentences":  createrandom.Sentences,
	"paragraphs": createrandom.Paragraphs}

func main() {
	templateFile, readError := ioutil.ReadFile("./templates/sample_template.md")
	check(readError)
	templates, err := template.New("todos").Funcs(mapToFunctions).Parse(string(templateFile))
	if err != nil {
		panic(err)
	}
	createdTemplate, createErr := os.Create("./templates/create_template.md")
	if createErr != nil {
		log.Println("Create template file: ", err)
		return
	}
	executeError = templates.Execute(createdTemplate, "")
	if executeError != nil {
		panic(err)
	}
}
