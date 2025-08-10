package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	input := "(+ 1 2 3 (* 5 6 7) 8)"
	expOutput := parserToken{
		Type: PARSER_LIST,
		Children: []parserToken{
			{Type: PARSER_SYMBOL, Value: symbol{OPERATOR, "+"}},
			{Type: PARSER_SYMBOL, Value: number{Value: 1}},
			{Type: PARSER_SYMBOL, Value: number{Value: 2}},
			{Type: PARSER_SYMBOL, Value: number{Value: 3}},
			{Type: PARSER_LIST, Children: []parserToken{
				{Type: PARSER_SYMBOL, Value: symbol{OPERATOR, "*"}},
				{Type: PARSER_SYMBOL, Value: number{Value: 5}},
				{Type: PARSER_SYMBOL, Value: number{Value: 6}},
				{Type: PARSER_SYMBOL, Value: number{Value: 7}},
			}},
			{Type: PARSER_SYMBOL, Value: number{Value: 8}},
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
	if tree.Type != PARSER_LIST {
		t.Errorf("Tree mismatch")
	}
	if len(tree.Children) != len(expOutput.Children) {
		t.Errorf("Tree mismatch")
	}
	for idx, child := range tree.Children {
		if child.Type != expOutput.Children[idx].Type {
			t.Errorf("Tree mismatch")
		}

		if child.Value == nil && expOutput.Children[idx].Value != nil {
			t.Errorf("Tree mismatch")
		}
		if child.Value != nil && expOutput.Children[idx].Value == nil {
			t.Errorf("Tree mismatch")
		}
		if child.Value != nil && expOutput.Children[idx].Value != nil && child.Value.GetType() != expOutput.Children[idx].Value.GetType() {
			t.Errorf("Tree mismatch")
		}
		switch child.Value.(type) {
		case symbol:
			if child.Value.(symbol).Value != expOutput.Children[idx].Value.(symbol).Value {
				t.Errorf("Tree mismatch")
			}
		case number:
			if child.Value.(number).Value != expOutput.Children[idx].Value.(number).Value {
				t.Errorf("Tree mismatch")
			}
		}
	}
}

func TestParseInvalid(t *testing.T) {
	input := "(+ 1 2 3 (* 5 6 7) 8"

	tokens, _ := Scan(input)

	_, err := parse(tokens)
	if err == nil {
		t.Errorf("Parse error: expected error, got nil")
	}
}
