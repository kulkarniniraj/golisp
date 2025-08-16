package main

import (
	"strings"

	"github.com/elk-language/go-prompt"
)

func shouldEndMultiLineInput(p *prompt.Prompt, indentSize int) (int, bool) {
	buffer := p.Buffer()
	lines := buffer.Document().Lines()
	if len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		return 0, true
	}

	// Check if the input is valid with closing parenthesis
	text := buffer.Text()
	tokens, err := Scan(text)
	if err != nil {
		return indentSize, false
	}
	_, err = parse(tokens)
	if err != nil {
		return indentSize, false
	}
	return 0, true
}

func setupInputReader() {
	p := prompt.New(
		repl,
		prompt.WithPrefix("byolisp> "),
		prompt.WithIndentSize(2),
		prompt.WithExecuteOnEnterCallback(shouldEndMultiLineInput),
	)
	p.Run()
}
