package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type EnvFunc func(parserToken) (parserToken, error)

// var Env = make(map[string]EnvFunc)
var Env = make(map[string]EnvFunc, 4)

func initEvaluator() {
	Env["+"] = add
	Env["-"] = sub
	Env["*"] = mul
	Env["/"] = div
}

func add(tree parserToken) (parserToken, error) {
	log.SetLevel(log.InfoLevel)
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	args := tree.Children
	eArgs := make([]parserToken, 0, len(args))

	for _, arg := range args {
		evaluatedArg, err := evaluate(arg)
		if err != nil {
			return parserToken{}, err
		}
		eArgs = append(eArgs, evaluatedArg)
	}

	assert(len(eArgs) > 0 && eArgs[0].Value.GetType() == SYMBOL, "invalid argument type")
	for _, arg := range eArgs[1:] {
		log.Debug(arg)
		assert(arg.Value.GetType() == NUMBER, "invalid argument type")
	}
	total := 0.0
	for _, arg := range eArgs[1:] {
		total += arg.Value.(number).Value
	}
	return parserToken{Type: PARSER_SYMBOL, Value: number{Value: total}}, nil
}

func sub(tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	args := tree.Children
	eArgs := make([]parserToken, 0, len(args))

	for _, arg := range args {
		evaluatedArg, err := evaluate(arg)
		if err != nil {
			return parserToken{}, err
		}
		eArgs = append(eArgs, evaluatedArg)
	}
	assert(len(eArgs) > 0 && eArgs[0].Value.GetType() == SYMBOL, "invalid argument type")
	for _, arg := range eArgs[1:] {
		assert(arg.Value.GetType() == NUMBER, "invalid argument type")
	}
	total := eArgs[1].Value.(number).Value
	for _, arg := range eArgs[2:] {
		total -= arg.Value.(number).Value
	}
	return parserToken{Type: PARSER_SYMBOL, Value: number{Value: total}}, nil
}

func mul(tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	args := tree.Children
	eArgs := make([]parserToken, 0, len(args))

	for _, arg := range args {
		evaluatedArg, err := evaluate(arg)
		if err != nil {
			return parserToken{}, err
		}
		eArgs = append(eArgs, evaluatedArg)
	}
	assert(len(eArgs) > 0 && eArgs[0].Value.GetType() == SYMBOL, "invalid argument type")
	for _, arg := range eArgs[1:] {
		assert(arg.Value.GetType() == NUMBER, "invalid argument type")
	}
	total := 1.0
	for _, arg := range eArgs[1:] {
		total *= arg.Value.(number).Value
	}
	return parserToken{Type: PARSER_SYMBOL, Value: number{Value: total}}, nil
}

func div(tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	args := tree.Children
	eArgs := make([]parserToken, 0, len(args))

	for _, arg := range args {
		evaluatedArg, err := evaluate(arg)
		if err != nil {
			return parserToken{}, err
		}
		eArgs = append(eArgs, evaluatedArg)
	}
	assert(len(eArgs) > 0 && eArgs[0].Value.GetType() == SYMBOL, "invalid argument type")
	for _, arg := range eArgs[1:] {
		assert(arg.Value.GetType() == NUMBER, "invalid argument type")
	}
	total := eArgs[1].Value.(number).Value
	for _, arg := range eArgs[2:] {
		total /= arg.Value.(number).Value
	}
	return parserToken{Type: PARSER_SYMBOL, Value: number{Value: total}}, nil
}

func evaluate(tree parserToken) (parserToken, error) {
	if tree.Type == PARSER_SYMBOL {
		return tree, nil
	} else {
		tokens := tree.Children

		if len(tokens) == 0 {
			fmt.Println("Invalid token seq")
			fmt.Println(tokens)
			return tree, nil
		}

		operation := tokens[0].Value.(symbol).Value
		fun, ok := Env[operation]
		if !ok {
			return parserToken{}, fmt.Errorf("unknown operation: %s", operation)
		}
		return fun(tree)

	}
}
