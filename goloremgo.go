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
	"github.com/urjaacharya/goloremgo/utils/createcontent"
	"github.com/urjaacharya/goloremgo/utils/helpers"
)

func checkTemplates(templatePath string) {

	//Parse file
	fileNumberPattern := regexp.MustCompile(`LFE_(\d+)`)
	fileNumber, _ := strconv.Atoi((fileNumberPattern.FindStringSubmatch(templatePath))[1])

	nMaxFiles := 10
	if fileNumber >= nMaxFiles {
		panic("The number of files to generate must be less than " + strconv.Itoa(nMaxFiles) + " in " + templatePath)
	}
	templateFile, readError := ioutil.ReadFile(filepath.FromSlash(templatePath))
	helpers.Check(readError)
	templates, err := template.New(templatePath).Funcs(createcontent.MapToFunctions).Parse(string(templateFile))
	helpers.Check(err)

	var temp bytes.Buffer
	executeError := templates.Execute(&temp, "")
	if executeError != nil {
		panic(executeError)
	}
}

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
	seed, dirPath, forceOverwrite, recursive := helpers.ReadArgs()
	rand.Seed(int64(seed))
	var allTemplates []string
	var templateFound bool
	var myErr error
	if recursive {
		allTemplates, templateFound, myErr = helpers.GetAllTemplateNames(dirPath)
	} else {
		allTemplates, templateFound, myErr = helpers.GetTemplateNames(dirPath)
	}
	helpers.Check(myErr)
	if templateFound {
		for _, myTemplatePath := range allTemplates {
			checkTemplates(myTemplatePath)
		}
		for _, myTemplatePath := range allTemplates {
			generateTemplate(myTemplatePath, forceOverwrite)
		}
	} else {
		fmt.Println("Any matching templates not found")
	}
}
