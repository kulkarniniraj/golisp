package main

import "testing"

func TestScan(t *testing.T) {
	input := "(+ 1 2 3 (* 5 6 7) 8)"
	exp_output := []token{
		{OPEN_PAREN, "("},
		{OPERATOR, "+"},
		{NUMBER, "1"},
		{NUMBER, "2"},
		{NUMBER, "3"},
		{OPEN_PAREN, "("},
		{OPERATOR, "*"},
		{NUMBER, "5"},
		{NUMBER, "6"},
		{NUMBER, "7"},
		{CLOSE_PAREN, ")"},
		{NUMBER, "8"},
		{CLOSE_PAREN, ")"},
	}
	tokens, err := Scan(input)
	if err != nil {
		t.Errorf("Scan error: %v", err)
	}
	for idx, token := range tokens {
		if token != exp_output[idx] {
			t.Errorf("Token mismatch at index %d: expected %v, got %v", idx, exp_output[idx], token)
		}
	}
}

func TestScanInvalid(t *testing.T) {
	input := "(+ 1 2 3 (* 5 6 7) 8 asd"
	out, err := Scan(input)
	t.Log(out)
	if err == nil {
		t.Errorf("Scan error: expected error, got nil")

	}
}
