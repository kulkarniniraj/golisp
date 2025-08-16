package main

import (
	"strings"

	"github.com/elk-language/go-prompt"
)

func setupInputReader() {
	p := prompt.New(
		repl,
		prompt.WithPrefix("byolisp> "),
		prompt.WithIndentSize(2),
		prompt.WithExecuteOnEnterCallback(func(p *prompt.Prompt, indentSize int) (int, bool) {
			buffer := p.Buffer()
			lines := buffer.Document().Lines()
			if len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
				return 0, true
			}
			return indentSize, false
		}),
	)
	p.Run()
}
