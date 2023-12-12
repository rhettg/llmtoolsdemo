package main

import (
	"fmt"

	"github.com/vikesh-raj/go-sentencepiece-encoder/sentencepiece"
)

const text = "This is a sample text"

func main() {
	spm, _ := sentencepiece.NewSentencepieceFromFile("spiece.model", false)

	if bos, ok := spm.GetControlWord("<s>"); ok {
		fmt.Println("Start: ", bos)
	}
	if eos, ok := spm.GetControlWord("</s>"); ok {
		fmt.Println("End: ", eos)
	}

	tokens := spm.Tokenize(text)
	fmt.Println(tokens)
}
