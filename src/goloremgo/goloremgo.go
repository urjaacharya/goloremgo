package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"text/template"
	"utils/createcontent"
	"utils/helpers"
)

//Generates random template
func generateTemplate(templatePath string, forceOverwrite bool) {
	dirName, fileName := filepath.Split(templatePath)
	templateFile, readError := ioutil.ReadFile(filepath.FromSlash(templatePath))
	helpers.Check(readError)
	templates, err := template.New(templatePath).Funcs(createcontent.MapToFunctions).Parse(string(templateFile))
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
			var dump bytes.Buffer
			executeError := templates.Execute(&dump, "")
			if executeError != nil {
				panic(executeError)
			}
			createdTemplate, createErr := os.Create(currentPath)
			_, writeErr := createdTemplate.WriteString(dump.String())
			if writeErr != nil {
				fmt.Println("Ha Ha")
			}
			if createErr != nil {
				fmt.Println("Create template file: ", createErr)
				return
			}
			if !os.IsNotExist(exitsErr) {
				fmt.Println("'" + currentPath + "' was replaced.")
			}
		} else if !os.IsNotExist(exitsErr) && !forceOverwrite {
			fmt.Println(("'" + currentPath + "' already exists. Use -f to replace file."))
		}
	}
}

func main() {
	seed, dirPath, forceOverwrite := helpers.ReadArgs()
	//flag.Usage = helpers.Usage
	var templateFound bool
	rand.Seed(int64(seed))
	dirErr := filepath.Walk(dirPath, func(path string, info os.FileInfo, fileErr error) error {
		if fileErr != nil {
			return fileErr
		}

		matchesFormat, matchErr := regexp.MatchString(`LFS_.+_\d+\.md$`, path)
		if matchErr != nil {
			return matchErr
		}
		if matchesFormat && !templateFound {
			templateFound = true
		}
		if matchesFormat {
			generateTemplate(path, forceOverwrite)
		}
		return nil
	})
	helpers.Check(dirErr)

	if !templateFound {
		fmt.Println("Any template not found.")
	}
}
