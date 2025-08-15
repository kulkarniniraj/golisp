package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type EnvValType int

const (
	ENVVAL_SYMBOL EnvValType = iota
	ENVVAL_FUNCTION
)

type EnvVal interface {
	GetType() EnvValType
}
type EnvFunc struct {
	Fun func(parserToken) (parserToken, error)
}

func (e EnvFunc) GetType() EnvValType {
	return ENVVAL_FUNCTION
}

type EnvSymbol struct {
	Val parserToken
}

func (e EnvSymbol) GetType() EnvValType {
	return ENVVAL_SYMBOL
}

// var Env = make(map[string]EnvFunc)
var Env = make(map[string]EnvVal, 4)

func initEvaluator() {
	Env["+"] = EnvFunc{Fun: add}
	Env["-"] = EnvFunc{Fun: sub}
	Env["*"] = EnvFunc{Fun: mul}
	Env["/"] = EnvFunc{Fun: div}
	Env["quote"] = EnvFunc{Fun: quote}
	Env["list"] = EnvFunc{Fun: list}
	Env["head"] = EnvFunc{Fun: head}
	Env["tail"] = EnvFunc{Fun: tail}
	Env["join"] = EnvFunc{Fun: join}
	Env["eval"] = EnvFunc{Fun: evalFun}
	Env["def"] = EnvFunc{Fun: def}
}

func add(tree parserToken) (parserToken, error) {
	log.SetLevel(log.InfoLevel)
	log.Debug(tree.Type, tree.Type == PARSER_LIST)
	log.Debug(tree.Value)
	log.Debug(tree.Children)
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

func quote(tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children

	return parserToken{Type: PARSER_LIST, Children: nodes[1:]}, nil
}

func list(tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	eNodes := make([]parserToken, 0, len(nodes))
	for _, node := range nodes[1:] {
		evaluatedNode, err := evaluate(node)
		if err != nil {
			return parserToken{}, err
		}
		eNodes = append(eNodes, evaluatedNode)
	}

	return parserToken{Type: PARSER_LIST, Children: eNodes}, nil
}

func head(tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	child1, err := evaluate(nodes[1])
	if err != nil {
		return parserToken{}, err
	}
	return child1.Children[0], nil
}

func tail(tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	child1, err := evaluate(nodes[1])
	if err != nil {
		return parserToken{}, err
	}
	return parserToken{Type: PARSER_LIST, Children: child1.Children[1:]}, nil
}

func join(tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	enodes := make([]parserToken, 0, len(nodes))
	joinCnt := 0
	for _, node := range nodes[1:] {
		evaluatedNode, err := evaluate(node)
		if err != nil {
			return parserToken{}, err
		}
		joinCnt += len(evaluatedNode.Children)
		enodes = append(enodes, evaluatedNode)
	}

	child1 := make([]parserToken, 0, joinCnt)
	for _, node := range enodes {
		child1 = append(child1, node.Children...)
	}
	return parserToken{Type: PARSER_LIST, Children: child1}, nil
}

func evalFun(tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	child1, err := evaluate(nodes[1])
	if err != nil {
		return parserToken{}, err
	}
	return evaluate(child1)
}

func def(tree parserToken) (parserToken, error) {
	// (def a 5)
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	variable := nodes[1]
	assert(variable.Type == PARSER_SYMBOL, "invalid argument type")
	assert(variable.Value.GetType() == SYMBOL, "invalid argument type")
	value, err := evaluate(nodes[2])
	if err != nil {
		return parserToken{}, err
	}

	Env[variable.Value.(symbol).Value] = EnvSymbol{Val: value}
	return parserToken{}, nil
}
func evaluate(tree parserToken) (parserToken, error) {
	if tree.Type == PARSER_SYMBOL {
		switch tree.Value.(type) {
		case symbol:
			val := Env[tree.Value.(symbol).Value]
			switch val := val.(type) {
			case EnvSymbol:
				return val.Val, nil
			case EnvFunc:
				return tree, nil
			default:
				return parserToken{}, fmt.Errorf("invalid symbol type: %s", tree.Value)
			}
		case number:
			return tree, nil
		default:
			return parserToken{}, fmt.Errorf("invalid symbol type: %s", tree.Value)
		}
	} else {
		tokens := tree.Children

		if len(tokens) == 0 {
			fmt.Println("Invalid token seq")
			fmt.Println(tokens)
			return tree, nil
		}

		operationPT, err := evaluate(tokens[0])
		if err != nil {
			return parserToken{}, err
		}
		operation := operationPT.Value.(symbol)
		operationVal := operation.Value
		fun, ok := Env[operationVal]
		if !ok {
			return parserToken{}, fmt.Errorf("unknown operation: %s", operationVal)
		}
		assert(fun.GetType() == ENVVAL_FUNCTION, "invalid operation type")
		return fun.(EnvFunc).Fun(tree)

	}
}
