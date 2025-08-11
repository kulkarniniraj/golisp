package main

import "testing"

func TestEvaluate(t *testing.T) {
	initEvaluator()
	input := "(+ 1 2 3 (* 5 6 7) 8)"
	exp_output := float64(1 + 2 + 3 + 5*6*7 + 8)
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(tree)
	if output.Value.(number).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.Value.(number).Value)
	}
}

func TestEvaluateAdd(t *testing.T) {
	initEvaluator()
	input := "(+ 1 2)"
	exp_output := float64(1 + 2)
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(tree)
	if output.Value.(number).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.Value.(number).Value)
	}
}

func TestEvaluateListOps(t *testing.T) {
	initEvaluator()
	input := "(head (list 1 2 3))"
	exp_output := float64(1)
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(tree)
	if output.Value.(number).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.Value.(number).Value)
	}
}

func TestEvaluateListOps2(t *testing.T) {
	initEvaluator()
	input := "(tail (list 1 2 3))"
	exp_output := parserToken{Type: PARSER_LIST, Children: []parserToken{{Type: PARSER_SYMBOL, Value: number{Value: 2}}, {Type: PARSER_SYMBOL, Value: number{Value: 3}}}}
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(tree)
	if len(output.Children) != len(exp_output.Children) {
		t.Errorf("Expected %f, got %f", exp_output, output.Value.(number).Value)
	}
}

func TestEvaluateListOps3(t *testing.T) {
	initEvaluator()
	input := "(join (list 1 2 3) (list 4 5 6))"
	exp_output := parserToken{Type: PARSER_LIST, Children: []parserToken{{Type: PARSER_SYMBOL, Value: number{Value: 1}}, {Type: PARSER_SYMBOL, Value: number{Value: 2}}, {Type: PARSER_SYMBOL, Value: number{Value: 3}}, {Type: PARSER_SYMBOL, Value: number{Value: 4}}, {Type: PARSER_SYMBOL, Value: number{Value: 5}}, {Type: PARSER_SYMBOL, Value: number{Value: 6}}}}
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(tree)
	if len(output.Children) != len(exp_output.Children) {
		t.Errorf("Expected %f, got %f", exp_output, output.Value.(number).Value)
	}
}
