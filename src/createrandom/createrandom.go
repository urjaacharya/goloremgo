package createrandom

import (
	"strings"
)

var sentences = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Imperdiet sed euismod nisi porta lorem mollis aliquam. Aliquet enim tortor at auctor urna nunc id. Libero volutpat sed cras ornare arcu dui vivamus arcu felis. In dictum non consectetur a erat nam. Ut pharetra sit amet aliquam. Adipiscing elit duis tristique sollicitudin nibh sit amet commodo nulla."

//WordArray test
func WordArray(sentences string, n int) string {
	words := strings.Split(sentences, " ")
	finalWords := ""
	for i := 0; i < n; i++ {
		if !(strings.HasSuffix(words[i], ",") || strings.HasSuffix(words[i], ".")) {
			finalWords = finalWords + " " + words[i]
		}
	}
	return finalWords
}

func sentenceArray(n int) []string {
	sentencesArray := strings.Split(sentences, ".")
	var finalSentences []string
	for i := 0; i < n; i++ {
		finalSentences = append(finalSentences, sentencesArray[i])
	}
	return finalSentences
}

func para(n int) [][]string {
	firstPara := sentenceArray(8)
	var finalParas [][]string
	for i := 0; i < n; i++ {
		finalParas = append(finalParas, firstPara)
	}
	return finalParas
}

// func main() {
// 	//word_array(10)
// 	//sentence_array(2)
// 	fmt.Println(para(6))
//}
