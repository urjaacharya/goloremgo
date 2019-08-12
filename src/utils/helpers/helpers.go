package helpers

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
	fmt.Println("goloremgo USAGE")
	fmt.Println("===============")
	fmt.Println("-p   REQUIRED: root directory that contains all the templates to be processed.")
	fmt.Println("-s   OPTIONAL: seed to reproduce randomly generated contents.")
	fmt.Println("-f   OPTIONAL: specify whether to overwrite files if they already exist.")
	fmt.Println("-h   print this usage information")
	os.Exit(1)
}

//ReadArgs Reads user provided arguments
func ReadArgs() (int, string, bool) {
	initialSeed := int(time.Now().UTC().UnixNano())
	dirPath := flag.String("p", "", "REQUIRED: root directory that contains all the templates to be processed.")
	randSeed := flag.Int("s", initialSeed, "OPTIONAL: seed to reproduce randomly generated contents.")
	forceOverwrite := flag.Bool("f", false, "OPTIONAL: specify whether to overwrite files if they already exist.")

	flag.Parse()

	if *dirPath == "" {
		fmt.Println("ERROR: path of the root directory containing the template files is not provided.")
		os.Exit(1)
	}

	rootDir := filepath.FromSlash(*dirPath)
	return *randSeed, rootDir, *forceOverwrite
}
