package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func print(tree ParserToken, indent int) {
	switch tree := tree.(type) {
	case ParserList:
		for _, child := range tree.Children {
			print(child, indent+1)
		}
	case ParserSymbol:
		fmt.Println(tree.Value)
	case ParserNumber:
		fmt.Println(tree.Value)
	case ParserFunc:
		fmt.Println("<func>")
	default:
		fmt.Println(tree)
	}
}

func init() {
	GlobalEnv = initEvaluator(GlobalEnv)
}

func repl(input string) {
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
		log.Debug("Tree:", tree)
		evaluatedTree, err := evaluate(GlobalEnv, tree)
		if err != nil {
			fmt.Println("Evaluate Error:", err)
			return
		}
		print(evaluatedTree, 0)
	}
}

func main() {
	_ = initEvaluator(GlobalEnv)
	setupInputReader()
	fmt.Println("Exiting...")
}
