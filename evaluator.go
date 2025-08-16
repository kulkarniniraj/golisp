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
	Fun func(map[string]EnvVal, parserToken) (parserToken, error)
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
var GlobalEnv = make(map[string]EnvVal, 0)

func initEvaluator(Env map[string]EnvVal) map[string]EnvVal {
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
	return Env
}

func add(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
	log.SetLevel(log.InfoLevel)
	log.Debug(tree.Type, tree.Type == PARSER_LIST)
	log.Debug(tree.Value)
	log.Debug(tree.Children)
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	args := tree.Children
	eArgs := make([]parserToken, 0, len(args))

	for _, arg := range args {
		evaluatedArg, err := evaluate(Env, arg)
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

func sub(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	args := tree.Children
	eArgs := make([]parserToken, 0, len(args))

	for _, arg := range args {
		evaluatedArg, err := evaluate(Env, arg)
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

func mul(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	args := tree.Children
	eArgs := make([]parserToken, 0, len(args))

	for _, arg := range args {
		evaluatedArg, err := evaluate(Env, arg)
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

func div(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	args := tree.Children
	eArgs := make([]parserToken, 0, len(args))

	for _, arg := range args {
		evaluatedArg, err := evaluate(Env, arg)
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

func quote(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children

	return parserToken{Type: PARSER_LIST, Children: nodes[1:]}, nil
}

func list(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	eNodes := make([]parserToken, 0, len(nodes))
	for _, node := range nodes[1:] {
		evaluatedNode, err := evaluate(Env, node)
		if err != nil {
			return parserToken{}, err
		}
		eNodes = append(eNodes, evaluatedNode)
	}

	return parserToken{Type: PARSER_LIST, Children: eNodes}, nil
}

func head(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	child1, err := evaluate(Env, nodes[1])
	if err != nil {
		return parserToken{}, err
	}
	return child1.Children[0], nil
}

func tail(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	child1, err := evaluate(Env, nodes[1])
	if err != nil {
		return parserToken{}, err
	}
	return parserToken{Type: PARSER_LIST, Children: child1.Children[1:]}, nil
}

func join(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	enodes := make([]parserToken, 0, len(nodes))
	joinCnt := 0
	for _, node := range nodes[1:] {
		evaluatedNode, err := evaluate(Env, node)
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

func evalFun(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	child1, err := evaluate(Env, nodes[1])
	if err != nil {
		return parserToken{}, err
	}
	return evaluate(Env, child1)
}

func def(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
	// (def a 5)
	assert(tree.Type == PARSER_LIST, "invalid argument type")
	nodes := tree.Children
	variable := nodes[1]
	assert(variable.Type == PARSER_SYMBOL, "invalid argument type")
	assert(variable.Value.GetType() == SYMBOL, "invalid argument type")
	value, err := evaluate(Env, nodes[2])
	if err != nil {
		return parserToken{}, err
	}

	Env[variable.Value.(symbol).Value] = EnvSymbol{Val: value}
	return parserToken{}, nil
}
func evaluate(Env map[string]EnvVal, tree parserToken) (parserToken, error) {
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

		operationPT, err := evaluate(Env, tokens[0])
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
		return fun.(EnvFunc).Fun(Env, tree)

	}
}
