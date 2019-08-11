package main

import (
	"createrandom"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"text/template"
)

//ReadArgs Reads user provided arguments
func ReadArgs() (int, string, bool) {
	dirPath := flag.String("p", "", "a string")
	randSeed := flag.Int("s", 42, "an int")
	forceOverwrite := flag.Bool("f", false, "a bool")
	flag.Parse()
	if *dirPath == "" {
		fmt.Println("Error: File path not provided.")
		os.Exit(1)
	}

	rootDir := filepath.FromSlash(*dirPath)
	return *randSeed, rootDir, *forceOverwrite
}

// To recursively read files in a directory
func findFiles(dirPath string, forceOverwrite bool) {
	dirErr := filepath.Walk(dirPath, func(path string, info os.FileInfo, fileErr error) error {
		if fileErr != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, fileErr)
			return fileErr
		}

		matchesFormat, matchErr := regexp.MatchString(`LFS_.+_\d+\.md$`, path)
		if matchErr != nil {
			//TO DO: what to do on match error?
		}
		if matchesFormat {
			generateTemplate(path, forceOverwrite)
			//fmt.Print(path)
		}
		return nil
	})
	if dirErr != nil {
		//TO DO: what to do on this?
	}
}

func generateTemplate(templatePath string, forceOverwrite bool) {
	dirName, fileName := filepath.Split(templatePath)
	templateFile, readError := ioutil.ReadFile(filepath.FromSlash(templatePath))
	check(readError)
	templates, err := template.New("todos").Funcs(mapToFunctions).Parse(string(templateFile))
	if err != nil {
		panic(err)
	}
	//Parse file
	fileNumberPattern := regexp.MustCompile(`LFE_(\d+)`)
	fileNumber, _ := strconv.Atoi((fileNumberPattern.FindStringSubmatch(fileName))[1])

	baseNamePattern := regexp.MustCompile(`LFS_(.+)_LFE`)
	baseName := (baseNamePattern.FindStringSubmatch(fileName))[1]

	//create fileNumInt number of templates
	for idx := 1; idx <= fileNumber; idx++ {
		currentPath := path.Join(dirName, (baseName + "_" + strconv.Itoa(idx) + ".md"))

		_, exitsErr := os.Stat(currentPath)

		if os.IsNotExist(exitsErr) || forceOverwrite {
			if !os.IsNotExist(exitsErr) {
				fmt.Println("'" + currentPath + "' already exists but will be replaced.")
			}
			createdTemplate, createErr := os.Create(currentPath)
			if createErr != nil {
				fmt.Println("Create template file: ", createErr)
				return
			}
			executeError := templates.Execute(createdTemplate, "")
			if executeError != nil {
				panic(executeError)
			}
		} else if !os.IsNotExist(exitsErr) && !forceOverwrite {
			fmt.Println(("'" + currentPath + "' already exists and will not be replaced."))
		}
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
	_, rootDir, forceOverwrite := ReadArgs()
	findFiles(rootDir, forceOverwrite)
	//generateTemplate("templates/article/loremTemplate_article_ln2.md")

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
