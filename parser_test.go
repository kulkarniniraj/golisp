package main

import "testing"

func TestParse(t *testing.T) {
	input := "(+ 1 2 3 (* 5 6 7) 8)"
	expOutput := parserToken{
		Type: PARSER_LIST,
		Children: []parserToken{
			{Type: PARSER_SYMBOL, Value: token{OPERATOR, "+"}},
			{Type: PARSER_SYMBOL, Value: token{NUMBER, "1"}},
			{Type: PARSER_SYMBOL, Value: token{NUMBER, "2"}},
			{Type: PARSER_SYMBOL, Value: token{NUMBER, "3"}},
			{Type: PARSER_LIST, Children: []parserToken{
				{Type: PARSER_SYMBOL, Value: token{OPERATOR, "*"}},
				{Type: PARSER_SYMBOL, Value: token{NUMBER, "5"}},
				{Type: PARSER_SYMBOL, Value: token{NUMBER, "6"}},
				{Type: PARSER_SYMBOL, Value: token{NUMBER, "7"}},
			}},
			{Type: PARSER_SYMBOL, Value: token{NUMBER, "8"}},
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
		if child.Value.Type != expOutput.Children[idx].Value.Type {
			t.Errorf("Tree mismatch")
		}
		if child.Value.Value != expOutput.Children[idx].Value.Value {
			t.Errorf("Tree mismatch")
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
