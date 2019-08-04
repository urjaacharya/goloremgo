package main

import (
	"createrandom"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Todo comment
type Todo struct {
	Name        string
	Description string
}

//Name comment
func Name() {
	fmt.Printf("hey")
}

//RandomCommands word
type RandomCommands struct {
	RandomWords string
}

var funcs = template.FuncMap{"words": createrandom.WordArray}

func main() {
	templateFile, readError := ioutil.ReadFile("./templates/sample_template.md")
	check(readError)
	fmt.Printf(string(templateFile))
	// fileScanner := bufio.NewScanner(templateFile)

	// var contents string
	// for fileScanner.Scan() {
	// 	fmt.Println(fileScanner.Text())
	// }
	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)

	data := []string{"two", "three"}
	t, err := template.New("todos").Funcs(funcs).Parse(string(templateFile))
	if err != nil {
		fmt.Printf("hey")
		panic(err)
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
