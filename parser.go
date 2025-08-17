package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type ParserToken interface {
	parserTokenMarker()
}

type ParserNumber struct {
	Value float64
}

type ParserSymbol struct {
	Value string
}

type ParserList struct {
	Children []ParserToken
}

type ParserFunc struct {
	Fun func(map[string]ParserToken, ParserToken) (ParserToken, error)
}

type ParserBool struct {
	Value bool
}

func (p ParserNumber) parserTokenMarker() {}
func (p ParserSymbol) parserTokenMarker() {}
func (p ParserList) parserTokenMarker()   {}
func (p ParserFunc) parserTokenMarker()   {}
func (p ParserBool) parserTokenMarker()   {}

func isOpenParen(token ParserToken) bool {
	symbol, ok := token.(ParserSymbol)
	if !ok {
		return false
	}
	return symbol.Value == "("
}

func parse(tokens []LexToken) (ParserToken, error) {
	log.Debug("Parsing tokens:", len(tokens))
	// stack of parserToken, size of stack is len(tokens)
	stack := make([]ParserToken, 0, len(tokens))

	if len(tokens) == 0 {
		return ParserList{}, nil
	}
	idx := 0
	for idx < len(tokens) {
		switch tokens[idx].(type) {
		case LexOpenParen:
			stack = append(stack, ParserSymbol{Value: "("})
			idx++
		case LexCloseParen:
			temp := make([]ParserToken, 0, len(tokens))
			for !isOpenParen(stack[len(stack)-1]) {
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
			stack = append(stack, ParserList{Children: temp})
			log.Debug("stack after folding: ", stack)
			idx++
		case LexNumber:
			stack = append(stack, ParserNumber{Value: tokens[idx].(LexNumber).Value})
			idx++
		case LexSymbol:
			symbol := tokens[idx].(LexSymbol)
			switch symbol.Value {
			case "true":
				stack = append(stack, ParserBool{Value: true})
			case "false":
				stack = append(stack, ParserBool{Value: false})
			default:
				stack = append(stack, ParserSymbol{Value: symbol.Value})
			}
			idx++
		default:
			return ParserList{}, fmt.Errorf("invalid token type")
		}
	}
	if len(stack) != 1 {
		log.Debug("Stack size is not 1, ", len(stack))
		log.Debug("Stack: ", stack)
		return ParserList{}, fmt.Errorf("Did you forget to close a list?")
	}
	return stack[0], nil
}
