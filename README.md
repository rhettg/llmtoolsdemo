# LLM Tools Demo

Just a demo and code snippets for trying some open source LLM tooling in Go.

## Usage

Download tokenizer model:

```console
wget -O spiece.model https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.1/resolve/main/tokenizer.model?download=true
```

Run examples:

```console
go run ./cmd/cybertron-tok
```

## Evaluation

I expected to get similar answers between these tokenizers. They are using the same model file at least.

Also, they should be matching the python implementation. But they do not.

```python
import transformers
tokenizer = transformers.AutoTokenizer.from_pretrained("mistralai/Mistral-7B-Instruct-v0.1")
tokenizer.tokenize("This is a sample text")
# ['▁This', '▁is', '▁a', '▁sample', '▁text']

tokenizer.encode("This is a sample text")
# [1, 851, 349, 264, 7324, 2245]
```

```console
$ go run ./cmd/sentencepiece-encoder-tok/
Start:  1
End:  2
[{542 ▁Th} {278 is} {349 ▁is} {264 ▁a} {268 ▁s} {314 am} {792 ple} {711 ▁te} {510 xt}]

$ go run ./cmd/cybertron-tok/
[▁Th is ▁is ▁a ▁s am ple ▁te xt]
{645 ▁Th} {381 is} {452 ▁is} {367 ▁a} {371 ▁s} {417 am} {895 ple} {814 ▁te} {613 xt}
[<s> ays <0xF3> <0x9E> ▁sweet ,\]
```

The `go-sentencepiece-encoder` does at least seem to match on the token for
`_a`. They are 264 in both. So at least it seems like it's using a similar vocab.

Finally, if I take the tokens from the python output and decode them with cybertron, I get:

    [<s> ays <0xF3> <0x9E> ▁sweet ,\]

gibberish...
