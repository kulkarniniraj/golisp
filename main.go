package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("byolisp> ")
		var input string
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		switch input {
		case "quit":
			return
		default:
			tokens, _ := Scan(input)
			tree, err := parse(tokens)
			if err != nil {
				fmt.Println("Parse Error:", err)
				return
			}
			fmt.Println(tree)
			// "+ 1 2 23 45.5 (/ 1 2 3)"
			// fmt.Printf("Your input: %s\n", input)
			// fmt.Printf("Tokens: %v\n", polishCalc(tokens))
			// fmt.Println(polishCalc(tokens))
			evaluatedTree, err := evaluate(tree)
			if err != nil {
				fmt.Println("Evaluate Error:", err)
				return
			}
			fmt.Println(evaluatedTree)
		}
	}
}
func main() {
	repl()
}
