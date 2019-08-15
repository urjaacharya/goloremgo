package createcontent

import (
	"fmt"
	"math/rand"
	"strings"
	"text/template"
	"time"
	"utils/data"
)

//Word Takes the number from user and Returns the given number of words
func Word(n int, sep string) string {
	sentences := data.Terms()
	words := strings.Split(sentences, " ")
	inds := rand.Perm(len(words))[:n]
	finalWords := make([]string, n)
	//TO DO: add check to see the user do not enter number greater than the number of words in sentences
	for i := 0; i < n; i++ {
		finalWords[i] = words[inds[i]]
	}
	return strings.Join(finalWords[:], sep)
}

//Sent Takes the the number from user and Returns the given number of sentences
func Sent(n int) string {
	finalSentences := make([]string, n)
	for i := 0; i < n; i++ {
		randNum := rand.Intn(7) + 8
		if randNum > 12 {
			n1 := int(randNum / 2)
			n2 := randNum - n1
			finalSentences[i] = strings.Title(Word(1, " ")) + " " + Word(n1, " ") + ", " + Word(n2, " ") + "."
		} else {
			finalSentences[i] = strings.Title(Word(1, " ")) + " " + Word(randNum, " ") + "."
		}
	}
	return strings.Join(finalSentences[:], " ")
}

//Para Takes the the number from user and Returns the given number of paragraphs
func Para(numSents, numParas int) string {
	finalParas := make([]string, numParas)
	for i := 0; i < numParas; i++ {
		finalParas[i] = Sent(numSents)
	}
	return strings.Join(finalParas[:], "\n\n")
}

// CapitalizeFirst capitalizes first word of the input string
func CapitalizeFirst(x string) string {
	xSplit := strings.Split(x, " ")
	xSplit[0] = strings.Title(xSplit[0])
	return strings.Join(xSplit, " ")
}

// Date returns a random date
func Date(x string, n int, format string) string {
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, x)
	if err != nil {
		fmt.Println(err)
	}
	i := rand.Intn(n)
	dateString := (parsedDate.AddDate(0, 0, i)).Format(format)
	return dateString
}

//MapToFunctions Map actions to functions
var MapToFunctions = template.FuncMap{
	"words":    Word,
	"sents":    Sent,
	"paras":    Para,
	"date":     Date,
	"capFirst": CapitalizeFirst,
	"capAll":   strings.ToUpper,
	"capEach":  strings.Title}
