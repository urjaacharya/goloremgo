package main

import (
	"createrandom"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
)

// To recursively read files in a directory
func findFiles() {
	err := filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		// if info.IsDir() && info.Name() == subDirToSkip {
		// 	fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
		// 	return filepath.SkipDir
		// }
		// fmt.Printf("visited file or dir: %q\n", path)
		matchesFormat, matchErr := regexp.MatchString(`loremTemplate.+\.md$`, path)
		if matchErr != nil {
			fmt.Printf("ehy")
		}
		fmt.Print(matchesFormat)
		return nil
	})
	if err != nil {
		fmt.Printf("yay")
	}
}

func generateTemplate(templatePath string) {
	dirName, fileName := filepath.Split(templatePath)
	templateFile, readError := ioutil.ReadFile(filepath.FromSlash(templatePath))
	check(readError)
	templates, err := template.New("todos").Funcs(mapToFunctions).Parse(string(templateFile))
	if err != nil {
		panic(err)
	}
	//Parse file
	// re := regexp.MustCompile(`ln(\d+)`)
	match := (regexp.MustCompile(`ln(\d+)`)).FindStringSubmatch(fileName)
	fmt.Println(match[1])
	fmt.Printf(dirName)
	createdTemplate, createErr := os.Create(filepath.FromSlash("./templates/create_template.md"))
	if createErr != nil {
		log.Println("Create template file: ", err)
		return
	}
	executeError := templates.Execute(createdTemplate, "")
	if executeError != nil {
		panic(err)
	}
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

var mapToFunctions = template.FuncMap{"words": createrandom.Words,
	"sentences":  createrandom.Sentences,
	"paragraphs": createrandom.Paragraphs}

func main() {
	generateTemplate("templates/article/loremTemplate_article_ln2.md")
	// templateFile, readError := ioutil.ReadFile(filepath.FromSlash("./templates/sample_template.md"))
	// check(readError)
	// templates, err := template.New("todos").Funcs(mapToFunctions).Parse(string(templateFile))
	// if err != nil {
	// 	panic(err)
	// }
	// createdTemplate, createErr := os.Create(filepath.FromSlash("./templates/create_template.md"))
	// if createErr != nil {
	// 	log.Println("Create template file: ", err)
	// 	return
	// }
	// executeError := templates.Execute(createdTemplate, "")
	// if executeError != nil {
	// 	panic(err)
	// }
}
