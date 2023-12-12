package main

import (
	"fmt"
	"os"

	ct_sp "github.com/nlpodyssey/cybertron/pkg/tokenizers/sentencepiece"
	vr_sp "github.com/vikesh-raj/go-sentencepiece-encoder/sentencepiece"
)

const text = "This is a sample text"

func runSentencePiece() {
	spm, _ := vr_sp.NewSentencepieceFromFile("tokenizer.model", false)

	if bos, ok := spm.GetControlWord("<s>"); ok {
		fmt.Println("Start: ", bos)
	}
	if eos, ok := spm.GetControlWord("</s>"); ok {
		fmt.Println("End: ", eos)
	}

	tokens := spm.Tokenize(text)
	fmt.Println(tokens)
}

func runCybertronSentencePiece() {
	sp, err := ct_sp.NewFromModelFolder("./", false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading sentencepiece model: %v\n", err)
		os.Exit(1)
	}

	tokens := sp.Tokenize(text)
	fmt.Println(tokens)
}

func main() {
	runSentencePiece()
}
