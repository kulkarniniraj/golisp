package main

import (
	"fmt"
	"regexp"
)

type TokenType int

const (
	OPEN_PAREN TokenType = iota
	CLOSE_PAREN
	OPERATOR
	NUMBER
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

type token struct {
	Type  TokenType
	Value string
}

func (t token) String() string {
	return fmt.Sprintf("* %s: %s *", t.Type.String(), t.Value)
}

func matchToToken(match []string) (token, error) {
	for idx, val := range match {
		if idx == 0 || idx == 1 || val == "" {
			continue
		}

		switch idx {
		case 2:
			return token{}, fmt.Errorf("Whitespace")
		case 3:
			return token{OPERATOR, val}, nil
		case 4:
			return token{NUMBER, val}, nil
		case 6:
			return token{OPEN_PAREN, val}, nil
		case 7:
			return token{CLOSE_PAREN, val}, nil
		}
	}
	return token{}, fmt.Errorf("invalid token")
}

func Scan(input string) ([]token, error) {
	var tokens []token
	whiteSpace := `(?P<ws>\s+)`
	operator := `(?P<op>\+|-|\*|/)`
	number := `(?P<num>\d+(\.\d+)?|\.\d+)`
	openParen := `(?P<openParen>\()`
	closeParen := `(?P<closeParen>\))`
	pattern := fmt.Sprintf("(%s|%s|%s|%s|%s)", whiteSpace, operator, number, openParen, closeParen)
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		token, err := matchToToken(match)
		if err != nil {
			continue
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}
