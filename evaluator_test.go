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
