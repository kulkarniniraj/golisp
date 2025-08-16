package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type LexToken interface {
	lexTokenMarker()
}

type LexEmpty struct {
}
type LexOpenParen struct {
}
type LexCloseParen struct {
}
type LexSymbol struct {
	Value string
}
type LexNumber struct {
	Value float64
}

func (e LexEmpty) lexTokenMarker()      {}
func (e LexOpenParen) lexTokenMarker()  {}
func (e LexCloseParen) lexTokenMarker() {}
func (e LexSymbol) lexTokenMarker()     {}
func (e LexNumber) lexTokenMarker()     {}

func matchToToken(match []string) (LexToken, error) {
	for idx, val := range match {
		if idx == 0 || idx == 1 || val == "" {
			continue
		}

		switch idx {
		case 2:
			return LexEmpty{}, nil
		case 3:
			fVal, _ := strconv.ParseFloat(val, 64)
			return LexNumber{fVal}, nil
		case 5:
			return LexOpenParen{}, nil
		case 6:
			return LexCloseParen{}, nil
		case 7:
			return LexSymbol{val}, nil
		default:
			return nil, fmt.Errorf("invalid token")
		}
	}
	return nil, fmt.Errorf("invalid token")
}

func Scan(input string) ([]LexToken, error) {
	var tokens []LexToken
	whiteSpace := `(?P<ws>\s+)`
	symbolChar := `[a-zA-Z0-9+\-*/_=<>@$%&?!~^]`
	operator := `(?P<op>` + symbolChar + `+)`
	number := `(?P<num>\d+(\.\d+)?|\.\d+)`
	openParen := `(?P<openParen>\()`
	closeParen := `(?P<closeParen>\))`
	errorPat := `(?P<error>.+)`
	pattern := fmt.Sprintf("(%s|%s|%s|%s|%s|%s)", whiteSpace, number, openParen, closeParen, operator, errorPat)
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		mToken, err := matchToToken(match)
		if err != nil {
			return nil, err
		}
		if mToken == (LexEmpty{}) {
			continue
		}
		tokens = append(tokens, mToken)
	}

	return tokens, nil
}
