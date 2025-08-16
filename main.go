package main

import (
	"fmt"

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

// func init() {
// 	GlobalEnv = initEvaluator(GlobalEnv)
// }

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
		print(evaluatedTree)
	}
}
func main() {
	_ = initEvaluator(GlobalEnv)
	setupInputReader()
	fmt.Println("Exiting...")
}
