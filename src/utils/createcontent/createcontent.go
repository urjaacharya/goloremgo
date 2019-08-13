package createcontent

import (
	"html/template"
	"math/rand"
	"strings"
	"utils/data"
)

//Words Takes the number from user and Returns the given number of words
func Words(n int, sep string) string {
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

//Sentences Takes the the number from user and Returns the given number of sentences
func Sentences(n int) string {
	finalSentences := make([]string, n)
	for i := 0; i < n; i++ {
		randNum := rand.Intn(7) + 8
		if randNum > 12 {
			n1 := int(randNum / 2)
			n2 := randNum - n1
			finalSentences[i] = strings.Title(Words(1, " ")) + " " + Words(n1, " ") + ", " + Words(n2, " ") + "."
		} else {
			finalSentences[i] = strings.Title(Words(1, " ")) + " " + Words(randNum, " ") + "."
		}
	}
	return strings.Join(finalSentences[:], " ")
}

//Paragraphs Takes the the number from user and Returns the given number of paragraphs
func Paragraphs(numSents, numParas int) string {
	finalParas := make([]string, numParas)
	for i := 0; i < numParas; i++ {
		finalParas[i] = Sentences(numSents)
	}
	return strings.Join(finalParas[:], "\n\n")
}

//MapToFunctions Map actions to functions
var MapToFunctions = template.FuncMap{
	"words":      Words,
	"sentences":  Sentences,
	"paragraphs": Paragraphs}
