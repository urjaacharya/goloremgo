package createcontent

import (
	"html/template"
	"math/rand"
	"strings"
	"utils/data"
)

//Words Takes the number from user and Returns the given number of words
func Words(n int, sep string) string {
	sentences := data.GetSentence()
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
	sentences := data.GetSentence()
	sentencesArray := strings.Split(sentences, ".")
	var finalSentences string
	for i := 0; i < n; i++ {
		if finalSentences == "" {
			finalSentences = finalSentences + sentencesArray[i] + "."
		} else {
			finalSentences = finalSentences + sentencesArray[i]
		}
	}
	return finalSentences + "."
}

//Paragraphs Takes the the number from user and Returns the given number of paragraphs
func Paragraphs(numSents, numParas int) string {
	firstParagraph := Sentences(numSents)
	var finalParas string
	for i := 0; i < numParas; i++ {
		finalParas = finalParas + firstParagraph + "\n\n"
	}
	return finalParas
}

//MapToFunctions Map actions to functions
var MapToFunctions = template.FuncMap{"words": Words,
	"sentences":  Sentences,
	"paragraphs": Paragraphs}
