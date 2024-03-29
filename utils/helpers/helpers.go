package helpers

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

//Check Panics when detects errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

//Usage Displays usage of defined arguments
func Usage() {
	fmt.Println("\ngoloremgo version 0.0 'hang tight' USAGE")
	fmt.Println("========================================")
	fmt.Println("-p   REQUIRED: directory that contains the templates to be processed.")
	fmt.Println("-f   OPTIONAL: overwrites files if they already exist.")
	fmt.Println("-r   OPTIONAL: searches for templates recursively inside all sub-directories")
	fmt.Println("-s   OPTIONAL: seed for randomly creating content. Default is current time.")
	fmt.Println("-h   print this usage information")
	fmt.Println("-v   version information")
	os.Exit(1)
}

// GetAllTemplateNames searches recursively inside 'path' and finds all templates
func GetAllTemplateNames(p string) ([]string, bool, error) {
	templateFound := false
	templateNamePattern := `LFS_.+_\d+\.glg$`
	var allTemplates []string
	dirErr := filepath.Walk(p, func(path string, info os.FileInfo, fileErr error) error {
		if fileErr != nil {
			return fileErr
		}
		matchesFormat, matchErr := regexp.MatchString(templateNamePattern, path)
		if matchErr != nil {
			return matchErr
		}
		if matchesFormat && !templateFound {
			templateFound = true
		}
		if matchesFormat {
			allTemplates = append(allTemplates, path)
		}
		return nil
	})
	return allTemplates, templateFound, dirErr
}

// GetTemplateNames searches only inside files and finds templates
func GetTemplateNames(dirPath string) ([]string, bool, error) {
	templateFound := false
	templateNamePattern := `LFS_.+_\d+\.glg$`
	var allTemplates []string
	files, err := ioutil.ReadDir(dirPath)
	for _, file := range files {
		matchesFormat, matchErr := regexp.MatchString(templateNamePattern, file.Name())
		if matchErr != nil {
			return nil, false, matchErr
		}
		if matchesFormat && !templateFound {
			templateFound = true
		}
		if matchesFormat {
			allTemplates = append(allTemplates, filepath.FromSlash(dirPath+"/"+file.Name()))
		}
	}
	return allTemplates, templateFound, err
}

//ReadArgs Reads user provided arguments
func ReadArgs() (int, string, bool, bool) {
	initialSeed := int(time.Now().UTC().UnixNano())
	dirPath := flag.String("p", "", "REQUIRED: directory containing templates")
	forceOverwrite := flag.Bool("f", false, "OPTIONAL: overwrite files or not")
	recursive := flag.Bool("r", false, "OPTIONAL: search directories recursively")
	randSeed := flag.Int("s", initialSeed, "OPTIONAL: seed")
	versionInfo := flag.Bool("v", false, "version information")

	flag.Usage = Usage
	flag.Parse()

	if *versionInfo {
		fmt.Println("\ngoloremgo version 0.0 'hang tight'")
		os.Exit(1)
	}

	if *dirPath == "" {
		fmt.Println("ERROR: directory path of template files not provided.")
		os.Exit(1)
	}

	rootDir := filepath.FromSlash(*dirPath)
	return *randSeed, rootDir, *forceOverwrite, *recursive
}
