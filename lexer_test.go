package main

import "testing"

func TestScan(t *testing.T) {
	input := "(+ 1 2 3 (* 5 6 7) 8)"
	exp_output := []LexToken{
		LexOpenParen{},
		LexSymbol{"+"},
		LexNumber{Value: 1},
		LexNumber{Value: 2},
		LexNumber{Value: 3},
		LexOpenParen{},
		LexSymbol{"*"},
		LexNumber{Value: 5},
		LexNumber{Value: 6},
		LexNumber{Value: 7},
		LexCloseParen{},
		LexNumber{Value: 8},
		LexCloseParen{},
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
	input := "(+ 1 2 3 (* 5 6 7) 8 asd #qwe"
	out, err := Scan(input)
	t.Log("scan result ", out)
	if err == nil {
		t.Errorf("Scan error: expected error, got nil")

	}
}
