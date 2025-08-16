package main

import "testing"

func TestAdd(t *testing.T) {
	initEvaluator(GlobalEnv)
	input := "(+ 1 2 3 8)"
	exp_output := float64(1 + 2 + 3 + 8)
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(GlobalEnv, tree)
	if output.(ParserNumber).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.(ParserNumber).Value)
	}
}

func TestEvaluate(t *testing.T) {
	initEvaluator(GlobalEnv)
	input := "(+ 1 2 3 (* 5 6 7) 8)"
	exp_output := float64(1 + 2 + 3 + 5*6*7 + 8)
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(GlobalEnv, tree)
	if output.(ParserNumber).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.(ParserNumber).Value)
	}
}

func TestEvaluateAdd(t *testing.T) {
	initEvaluator(GlobalEnv)
	input := "(+ 1 2)"
	exp_output := float64(1 + 2)
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(GlobalEnv, tree)
	if output.(ParserNumber).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.(ParserNumber).Value)
	}
}

func TestEvaluateArithmetic(t *testing.T) {
	initEvaluator(GlobalEnv)
	input := "(+ 1 2)"
	exp_output := float64(1 + 2)
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(GlobalEnv, tree)
	if output.(ParserNumber).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.(ParserNumber).Value)

	}

	input = "(- 1 2)"
	exp_output = float64(1 - 2)
	tokens, _ = Scan(input)
	tree, _ = parse(tokens)
	output, _ = evaluate(GlobalEnv, tree)
	if output.(ParserNumber).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.(ParserNumber).Value)
	}

	input = "(* 1 2)"
	exp_output = float64(1 * 2)
	tokens, _ = Scan(input)
	tree, _ = parse(tokens)
	output, _ = evaluate(GlobalEnv, tree)
	if output.(ParserNumber).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.(ParserNumber).Value)
	}

	input = "(/ 1 2)"
	exp_output = float64(1.0 / 2.0)
	tokens, _ = Scan(input)
	tree, _ = parse(tokens)
	output, _ = evaluate(GlobalEnv, tree)
	if output.(ParserNumber).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.(ParserNumber).Value)
	}
}

func TestEvaluateListOps(t *testing.T) {
	initEvaluator(GlobalEnv)
	input := "(head (list 1 2 3))"
	exp_output := float64(1)
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(GlobalEnv, tree)
	if output.(ParserNumber).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.(ParserNumber).Value)
	}
}

func TestEvaluateListOps2(t *testing.T) {
	initEvaluator(GlobalEnv)
	input := "(tail (list 1 2 3))"
	exp_output := ParserList{Children: []ParserToken{
		ParserNumber{Value: 2},
		ParserNumber{Value: 3}}}
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(GlobalEnv, tree)
	if len(output.(ParserList).Children) != len(exp_output.Children) {
		t.Errorf("Expected %v, got %v", exp_output, output)
	}
}

func TestEvaluateListOps3(t *testing.T) {
	initEvaluator(GlobalEnv)
	input := "(join (list 1 2 3) (list 4 5 6))"
	exp_output := ParserList{Children: []ParserToken{
		ParserNumber{Value: 1},
		ParserNumber{Value: 2},
		ParserNumber{Value: 3},
		ParserNumber{Value: 4},
		ParserNumber{Value: 5},
		ParserNumber{Value: 6}}}
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(GlobalEnv, tree)
	if len(output.(ParserList).Children) != len(exp_output.Children) {
		t.Errorf("Expected %v, got %v", exp_output, output)
	}
}

// test evalFun function
func TestEvaluateEval(t *testing.T) {
	initEvaluator(GlobalEnv)
	input := "(eval (list + 1 2))"
	exp_output := float64(1 + 2)
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, err := evaluate(GlobalEnv, tree)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if output.(ParserNumber).Value != exp_output {
		t.Errorf("Expected %f, got %f", exp_output, output.(ParserNumber).Value)
	}
}

// test quote
func TestEvaluateQuote(t *testing.T) {
	initEvaluator(GlobalEnv)
	input := "(quote (1 2 3))"
	exp_output := ParserList{Children: []ParserToken{
		ParserNumber{Value: 1},
		ParserNumber{Value: 2},
		ParserNumber{Value: 3}}}
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(GlobalEnv, tree)
	if len(output.(ParserList).Children) != len(exp_output.Children) {
		t.Errorf("Expected %v, got %v", exp_output, output)
	}
}

// test def
func TestEvaluateDef(t *testing.T) {
	initEvaluator(GlobalEnv)
	input := "(def a 5)"
	tokens, _ := Scan(input)
	tree, _ := parse(tokens)
	output, _ := evaluate(GlobalEnv, tree)
	val, ok := GlobalEnv["a"]
	if !ok {
		t.Errorf("Expected %v, got %v", output, val)
	}
	if val.(ParserNumber).Value != float64(5) {
		t.Errorf("Expected %f, got %f", float64(5), val.(ParserNumber).Value)
	}
}
