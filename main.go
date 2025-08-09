package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func polishCalc(tokens []token) float64 {
	adder := func(nums ...float64) float64 {
		total := 0.0
		for _, num := range nums {
			total += num
		}
		return total
	}
	multiplier := func(nums ...float64) float64 {
		total := 1.0
		for _, num := range nums {
			total *= num
		}
		return total
	}
	subtractor := func(nums ...float64) float64 {
		total := nums[0]
		for _, num := range nums[1:] {
			total -= num
		}
		return total
	}
	divider := func(nums ...float64) float64 {
		total := nums[0]
		for _, num := range nums[1:] {
			total /= num
		}
		return total
	}
	if len(tokens) == 0 || tokens[0].Type != OPEN_PAREN || tokens[len(tokens)-1].Type != CLOSE_PAREN {
		fmt.Println("Invalid token seq")
		fmt.Println(tokens)
		return 0
	}

	operation := tokens[1].Value
	args := []float64{}

	for _, token := range tokens[2 : len(tokens)-1] {
		if token.Type == NUMBER {
			n, _ := strconv.ParseFloat(token.Value, 64)
			args = append(args, n)
		}
	}

	switch operation {
	case "+":
		return adder(args...)
	case "-":
		return subtractor(args...)
	case "*":
		return multiplier(args...)
	case "/":
		return divider(args...)
	default:
		fmt.Println("Invalid operation")
		return 0
	}
}

func repl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("byolisp> ")
		var input string
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		switch input {
		case "quit":
			return
		default:
			tokens, _ := Scan(input)
			// "+ 1 2 23 45.5 (/ 1 2 3)"
			// fmt.Printf("Your input: %s\n", input)
			// fmt.Printf("Tokens: %v\n", polishCalc(tokens))
			fmt.Println(polishCalc(tokens))
		}
	}
}
func main() {
	repl()
}
