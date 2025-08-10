package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type TokenType int

const (
	OPEN_PAREN TokenType = iota
	CLOSE_PAREN
	OPERATOR
	NUMBER
	EMPTY
)

func (t TokenType) String() string {
	switch t {
	case OPEN_PAREN:
		return "OPEN_PAREN"
	case CLOSE_PAREN:
		return "CLOSE_PAREN"
	case OPERATOR:
		return "OPERATOR"
	case NUMBER:
		return "NUMBER"
	default:
		return "UNKNOWN"
	}
}

type token interface {
	GetType() TokenType
}

type empty struct {
}

func (e empty) GetType() TokenType {
	return EMPTY
}

func (e empty) String() string {
	return fmt.Sprintf(" %s ", e.GetType().String())
}

type symbol struct {
	Type  TokenType // open paren, close paren, operator
	Value string
}

func (s symbol) GetType() TokenType {
	return s.Type
}

func (s symbol) String() string {
	return fmt.Sprintf(" %s: \"%s\" ", s.Type.String(), s.Value)
}

type number struct {
	Value float64
}

func (n number) GetType() TokenType {
	return NUMBER
}

func (n number) String() string {
	return fmt.Sprintf(" %s: %f ", n.GetType().String(), n.Value)
}

func matchToToken(match []string) (token, error) {
	for idx, val := range match {
		if idx == 0 || idx == 1 || val == "" {
			continue
		}

		switch idx {
		case 2:
			return empty{}, nil
		case 3:
			return symbol{OPERATOR, val}, nil
		case 4:
			fVal, _ := strconv.ParseFloat(val, 64)
			return number{fVal}, nil
		case 6:
			return symbol{OPEN_PAREN, val}, nil
		case 7:
			return symbol{CLOSE_PAREN, val}, nil
		default:
			return nil, fmt.Errorf("invalid token")
		}
	}
	return nil, fmt.Errorf("invalid token")
}

func Scan(input string) ([]token, error) {
	var tokens []token
	whiteSpace := `(?P<ws>\s+)`
	operator := `(?P<op>\+|-|\*|/)`
	number := `(?P<num>\d+(\.\d+)?|\.\d+)`
	openParen := `(?P<openParen>\()`
	closeParen := `(?P<closeParen>\))`
	errorPat := `(?P<error>.+)`
	pattern := fmt.Sprintf("(%s|%s|%s|%s|%s|%s)", whiteSpace, operator, number, openParen, closeParen, errorPat)
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		mToken, err := matchToToken(match)
		if err != nil {
			return nil, err
		}
		if mToken == (empty{}) {
			continue
		}
		tokens = append(tokens, mToken)
	}

	return tokens, nil
}
