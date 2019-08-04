package main

import (
	"createrandom"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	sentence, err := ioutil.ReadFile("./sentence.txt")
	check(err)
	converted := createrandom.WordArray(string(sentence), 1)
	fmt.Print(converted)
}
