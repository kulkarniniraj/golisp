package main

import (
	"testing"
)

func compareParserToken(t *testing.T, tree ParserToken, expOutput ParserToken) {
	switch tree.(type) {
	case ParserList:
		if len(tree.(ParserList).Children) != len(expOutput.(ParserList).Children) {
			t.Errorf("Tree mismatch")
		}
		for idx, child := range tree.(ParserList).Children {
			compareParserToken(t, child, expOutput.(ParserList).Children[idx])
		}
	case ParserSymbol:
		if tree.(ParserSymbol).Value != expOutput.(ParserSymbol).Value {
			t.Errorf("Symbol mismatch")
		}
	case ParserNumber:
		if tree.(ParserNumber).Value != expOutput.(ParserNumber).Value {
			t.Errorf("Number mismatch")
		}
	default:
		t.Errorf("Unknown type")
	}
}

func TestParse(t *testing.T) {
	input := "(+ 1 2 3 (* 5 6 7) 8)"
	expOutput := ParserList{
		Children: []ParserToken{
			ParserSymbol{Value: "+"},
			ParserNumber{Value: 1},
			ParserNumber{Value: 2},
			ParserNumber{Value: 3},
			ParserList{
				Children: []ParserToken{
					ParserSymbol{Value: "*"},
					ParserNumber{Value: 5},
					ParserNumber{Value: 6},
					ParserNumber{Value: 7},
				}},
			ParserNumber{Value: 8},
		},
	}
	tokens, err := Scan(input)
	if err != nil {
		t.Errorf("Scan error: %v", err)
	}
	tree, err := parse(tokens)
	if err != nil {
		t.Errorf("Parse error: %v", err)
	}
	compareParserToken(t, tree, expOutput)
}

func TestParseInvalid(t *testing.T) {
	input := "(+ 1 2 3 (* 5 6 7) 8"

	tokens, _ := Scan(input)

	_, err := parse(tokens)
	if err == nil {
		t.Errorf("Parse error: expected error, got nil")
	}
}
