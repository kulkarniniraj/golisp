package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

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
			fmt.Println(evaluatedTree)
		}
	}
}
func main() {
	repl()
}
