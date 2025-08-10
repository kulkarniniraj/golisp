package main

import (
	"fmt"
	"strconv"
)

func evaluate(tree parserToken) (parserToken, error) {

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

	if tree.Type == PARSER_SYMBOL {
		return tree, nil
	} else {
		tokens := tree.Children

		if len(tokens) == 0 {
			fmt.Println("Invalid token seq")
			fmt.Println(tokens)
			return tree, nil
		}

		operation := tokens[0].Value.Value
		args := []float64{}

		for _, token := range tokens[1:] {
			eToken, err := evaluate(token)
			if err != nil {
				return parserToken{}, err
			}
			if eToken.Value.Type == NUMBER {
				n, _ := strconv.ParseFloat(eToken.Value.Value, 64)
				args = append(args, n)
			}
		}

		retVal := 0.0

		switch operation {
		case "+":
			retVal = adder(args...)
		case "-":
			retVal = subtractor(args...)
		case "*":
			retVal = multiplier(args...)
		case "/":
			retVal = divider(args...)
		default:
			fmt.Println("Invalid operation: ", operation)
			retVal = 0
		}

		return parserToken{
			Type: PARSER_SYMBOL,
			Value: token{
				Type:  NUMBER,
				Value: fmt.Sprintf("%f", retVal),
			},
		}, nil
	}
}
