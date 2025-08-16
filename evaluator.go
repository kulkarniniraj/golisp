package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

var GlobalEnv = make(map[string]ParserToken, 0)

func initEvaluator(Env map[string]ParserToken) map[string]ParserToken {
	Env["+"] = ParserFunc{Fun: add}
	Env["-"] = ParserFunc{Fun: sub}
	Env["*"] = ParserFunc{Fun: mul}
	Env["/"] = ParserFunc{Fun: div}
	Env["quote"] = ParserFunc{Fun: quote}
	Env["list"] = ParserFunc{Fun: list}
	Env["head"] = ParserFunc{Fun: head}
	Env["tail"] = ParserFunc{Fun: tail}
	Env["join"] = ParserFunc{Fun: join}
	// Env["eval"] = ParserFunc{Fun: evalFun}
	// Env["def"] = ParserFunc{Fun: def}
	return Env
}

func add(Env map[string]ParserToken, tree ParserToken) (ParserToken, error) {
	log.SetLevel(log.InfoLevel)
	list, ok := tree.(ParserList)
	if !ok {
		return ParserList{}, fmt.Errorf("invalid argument type")
	}

	args := list.Children
	eArgs := make([]ParserToken, 0, len(args))

	for _, arg := range args[1:] {
		evaluatedArg, err := evaluate(Env, arg)
		if err != nil {
			return ParserList{}, err
		}
		eArgs = append(eArgs, evaluatedArg)
	}

	assert(len(eArgs) > 1, "invalid argument type")
	for _, arg := range eArgs {
		log.Debug(arg)
		_, ok := arg.(ParserNumber)
		assert(ok, "invalid argument type, expected number")
	}
	total := 0.0
	for _, arg := range eArgs {
		val := arg.(ParserNumber)
		total += val.Value
	}
	return ParserNumber{Value: total}, nil
}

func sub(Env map[string]ParserToken, tree ParserToken) (ParserToken, error) {
	list, ok := tree.(ParserList)
	if !ok {
		return ParserList{}, fmt.Errorf("invalid argument type")
	}
	args := list.Children
	eArgs := make([]ParserToken, 0, len(args))

	for _, arg := range args {
		evaluatedArg, err := evaluate(Env, arg)
		if err != nil {
			return ParserList{}, err
		}
		eArgs = append(eArgs, evaluatedArg)
	}
	assert(len(eArgs) > 1, "invalid argument type")
	for _, arg := range eArgs[1:] {
		_, ok := arg.(ParserNumber)
		assert(ok, "invalid argument type, expected number")
	}
	total := eArgs[1].(ParserNumber).Value
	for _, arg := range eArgs[2:] {
		total -= arg.(ParserNumber).Value
	}
	return ParserNumber{Value: total}, nil
}

func mul(Env map[string]ParserToken, tree ParserToken) (ParserToken, error) {
	list, ok := tree.(ParserList)
	if !ok {
		return ParserList{}, fmt.Errorf("invalid argument type")
	}
	args := list.Children
	eArgs := make([]ParserToken, 0, len(args))

	for _, arg := range args {
		evaluatedArg, err := evaluate(Env, arg)
		if err != nil {
			return ParserList{}, err
		}
		eArgs = append(eArgs, evaluatedArg)
	}
	assert(len(eArgs) > 1, "invalid argument type")
	for _, arg := range eArgs[1:] {
		_, ok := arg.(ParserNumber)
		assert(ok, "invalid argument type, expected number")
	}
	total := 1.0
	for _, arg := range eArgs[1:] {
		total *= arg.(ParserNumber).Value
	}
	return ParserNumber{Value: total}, nil
}

func div(Env map[string]ParserToken, tree ParserToken) (ParserToken, error) {
	list, ok := tree.(ParserList)
	if !ok {
		return ParserList{}, fmt.Errorf("invalid argument type")
	}
	args := list.Children
	eArgs := make([]ParserToken, 0, len(args))

	for _, arg := range args {
		evaluatedArg, err := evaluate(Env, arg)
		if err != nil {
			return ParserList{}, err
		}
		eArgs = append(eArgs, evaluatedArg)
	}
	assert(len(eArgs) > 1, "invalid argument type")
	for _, arg := range eArgs[1:] {
		_, ok := arg.(ParserNumber)
		assert(ok, "invalid argument type, expected number")
	}
	total := eArgs[1].(ParserNumber).Value
	for _, arg := range eArgs[2:] {
		total /= arg.(ParserNumber).Value
	}
	return ParserNumber{Value: total}, nil
}

func quote(Env map[string]ParserToken, tree ParserToken) (ParserToken, error) {
	list, ok := tree.(ParserList)
	if !ok {
		return ParserList{}, fmt.Errorf("invalid argument type")
	}
	nodes := list.Children

	return nodes[1], nil
}

func list(Env map[string]ParserToken, tree ParserToken) (ParserToken, error) {
	list, ok := tree.(ParserList)
	if !ok {
		return ParserList{}, fmt.Errorf("invalid argument type")
	}
	nodes := list.Children
	eNodes := make([]ParserToken, 0, len(nodes))
	for _, node := range nodes[1:] {
		evaluatedNode, err := evaluate(Env, node)
		if err != nil {
			return ParserList{}, err
		}
		eNodes = append(eNodes, evaluatedNode)
	}

	return ParserList{Children: eNodes}, nil
}

func head(Env map[string]ParserToken, tree ParserToken) (ParserToken, error) {
	list, ok := tree.(ParserList)
	if !ok {
		return ParserList{}, fmt.Errorf("invalid argument type")
	}
	nodes := list.Children
	child1, err := evaluate(Env, nodes[1])
	if err != nil {
		return ParserList{}, err
	}
	return child1.(ParserList).Children[0], nil
}

func tail(Env map[string]ParserToken, tree ParserToken) (ParserToken, error) {
	list, ok := tree.(ParserList)
	if !ok {
		return ParserList{}, fmt.Errorf("invalid argument type")
	}
	nodes := list.Children
	child1, err := evaluate(Env, nodes[1])
	if err != nil {
		return ParserList{}, err
	}
	return ParserList{Children: child1.(ParserList).Children[1:]}, nil
}

func join(Env map[string]ParserToken, tree ParserToken) (ParserToken, error) {
	list, ok := tree.(ParserList)
	if !ok {
		return ParserList{}, fmt.Errorf("invalid argument type")
	}
	nodes := list.Children
	enodes := make([]ParserToken, 0, len(nodes))
	joinCnt := 0
	for _, node := range nodes[1:] {
		evaluatedNode, err := evaluate(Env, node)
		if err != nil {
			return ParserList{}, err
		}
		joinCnt += len(evaluatedNode.(ParserList).Children)
		enodes = append(enodes, evaluatedNode)
	}

	child1 := make([]ParserToken, 0, joinCnt)
	for _, node := range enodes {
		child1 = append(child1, node.(ParserList).Children...)
	}
	return ParserList{Children: child1}, nil
}

// func evalFun(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
// 	assert(tree.Type == PARSER_LIST, "invalid argument type")
// 	nodes := tree.Children
// 	child1, err := evaluate(Env, nodes[1])
// 	if err != nil {
// 		return parserToken{}, err
// 	}
// 	return evaluate(Env, child1)
// }

// func def(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
// 	// (def a 5)
// 	assert(tree.Type == PARSER_LIST, "invalid argument type")
// 	nodes := tree.Children
// 	variable := nodes[1]
// 	assert(variable.Type == PARSER_SYMBOL, "invalid argument type")
// 	assert(variable.Value.GetType() == SYMBOL, "invalid argument type")
// 	value, err := evaluate(Env, nodes[2])
// 	if err != nil {
// 		return parserToken{}, err
// 	}

// 	Env[variable.Value.(symbol).Value] = EnvSymbol{Val: value}
// 	return parserToken{}, nil
// }

// func lambda(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
// 	assert(tree.Type == PARSER_LIST, "invalid argument type")
// 	assert(len(tree.Children) == 3, "Need arglist and body for lambda")

// 	arglist := tree.Children[1]
// 	body := tree.Children[2]

// }
func evaluate(Env map[string]ParserToken, tree ParserToken) (ParserToken, error) {
	switch tree.(type) {
	case ParserSymbol:
		val := Env[tree.(ParserSymbol).Value]
		return val, nil
	case ParserNumber:
		return tree, nil
	case ParserFunc:
		return tree, nil
	case ParserList:
		tokens := tree.(ParserList).Children

		if len(tokens) == 0 {
			// fmt.Println("Invalid token seq")
			// fmt.Println(tokens)
			return tree, nil
		}

		operationPT, err := evaluate(Env, tokens[0])
		if err != nil {
			return ParserList{}, err
		}
		operation, ok := operationPT.(ParserFunc)
		if !ok {
			return ParserList{}, fmt.Errorf("unknown operation: %s", operationPT)
		}
		return operation.Fun(Env, tree)

	}
	return tree, nil
}
