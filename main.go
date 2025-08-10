package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func print(tree parserToken) {
	switch tree.Type {
	case PARSER_LIST:
		fmt.Println(tree)
	default:
		switch tree.Value.(type) {
		case symbol:
			fmt.Println(tree.Value.(symbol).Value)
		case number:
			fmt.Println(tree.Value.(number).Value)
		}
	}
}

func repl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("byolisp> ")
		var input string
		// read until blank line
		for {
			input1, err := reader.ReadString('\n')
			input1 = strings.TrimSpace(input1)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			input += input1
			if input1 == "" {
				break
			}
		}
		switch input {
		case "quit":
			return
		default:
			tokens, _ := Scan(input)
			tree, err := parse(tokens)
			if err != nil {
				fmt.Println("Parse Error:", err)
				continue
			}
			log.Debug("Tree:", tree)
			evaluatedTree, err := evaluate(tree)
			if err != nil {
				fmt.Println("Evaluate Error:", err)
				continue
			}
			print(evaluatedTree)
		}
	}
}
func main() {
	repl()
}
