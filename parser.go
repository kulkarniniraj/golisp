package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// log "github.com/sirupsen/logrus"

type parserTokenType int

const (
	PARSER_SYMBOL parserTokenType = iota
	PARSER_LIST
)

type parserToken struct {
	Type     parserTokenType
	Value    token
	Children []parserToken
}

func (p parserToken) String() string {
	switch p.Type {
	case PARSER_SYMBOL:
		return p.Value.String()
	case PARSER_LIST:
		var s string
		for _, child := range p.Children {
			s += " " + child.String()
		}
		return "(" + s + ")"
	default:
		return ""
	}
}

func parse(tokens []token) (parserToken, error) {
	log.SetLevel(log.InfoLevel)
	log.Info("Parsing tokens:", len(tokens))
	// stack of parserToken, size of stack is len(tokens)
	stack := make([]parserToken, 0, len(tokens))

	if len(tokens) == 0 {
		return parserToken{Type: PARSER_LIST}, nil
	} else if tokens[0].Type != OPEN_PAREN {
		if len(tokens) != 1 {
			return parserToken{}, fmt.Errorf("multiple symbols without list")
		}
		return parserToken{Type: PARSER_SYMBOL, Value: tokens[0]}, nil
	} else {
		idx := 0
		for idx < len(tokens) {
			log.Debug("Processing index: ", idx)
			if tokens[idx].Type != CLOSE_PAREN {
				log.Debug("Adding token: ", tokens[idx])
				stack = append(stack, parserToken{Type: PARSER_SYMBOL, Value: tokens[idx]})
			} else {
				log.Debug("Found close paren")
				temp := make([]parserToken, 0, len(tokens))
				for !(stack[len(stack)-1].Value.Type == OPEN_PAREN && stack[len(stack)-1].Type == PARSER_SYMBOL) {
					temp = append(temp, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = stack[:len(stack)-1]
				log.Debug("folding list: ", temp)
				log.Debug("stack: ", stack)
				// reverse temp and append to stack
				for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
					temp[i], temp[j] = temp[j], temp[i]
				}
				stack = append(stack, parserToken{Type: PARSER_LIST, Children: temp})
				log.Debug("stack after folding: ", stack)
			}
			idx++
		}
		if len(stack) != 1 {
			log.Error("Stack size is not 1, ", len(stack))
			log.Error("Stack: ", stack)
			return parserToken{}, fmt.Errorf("cannot parse to single list")
		}
		return stack[0], nil

	}
}
