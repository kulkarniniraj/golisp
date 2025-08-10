package main

import "testing"

func TestScan(t *testing.T) {
	input := "(+ 1 2 3 (* 5 6 7) 8)"
	exp_output := []token{
		symbol{OPEN_PAREN, "("},
		symbol{OPERATOR, "+"},
		number{Value: 1},
		number{Value: 2},
		number{Value: 3},
		symbol{OPEN_PAREN, "("},
		symbol{OPERATOR, "*"},
		number{Value: 5},
		number{Value: 6},
		number{Value: 7},
		symbol{CLOSE_PAREN, ")"},
		number{Value: 8},
		symbol{CLOSE_PAREN, ")"},
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
