package main

import (
	"fmt"
	"os"

	"github.com/nlpodyssey/cybertron/pkg/tokenizers/sentencepiece"
)

const text = "This is a sample text"

func main() {
	sp, err := sentencepiece.NewFromModelFolder("./", false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading sentencepiece model: %v\n", err)
		os.Exit(1)
	}

	tokens := sp.Tokenize(text)
	fmt.Println(tokens)

	encoded := sp.TokensToIDs(tokens)

	for ndx, id := range encoded {
		fmt.Printf("{%d %s} ", id, tokens[ndx])
	}

	fmt.Println("")

	/*
		Decode tokens from Python
		encoded = []int{1, 851, 349, 264, 7324, 2245}
		decoded := sp.IDsToTokens(encoded)
		fmt.Println(decoded)
	*/

}
