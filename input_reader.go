package main

import (
	"fmt"

	"github.com/elk-language/go-prompt"
)

// Global variable to store session history
var sessionHistory []string

func inputReader() (string, error) {
	exit := false
	t := prompt.Input(
		prompt.WithPrefix("byolisp> "),
		prompt.WithKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn: func(_ *prompt.Prompt) bool {
				exit = true
				return false
			},
		}),
		prompt.WithHistory(sessionHistory), // Use the persistent history
	)
	if exit {
		return "", fmt.Errorf("EOF")
	}
	
	// Add the input to history (only if it's not empty and not a duplicate of the last entry)
	if t != "" && (len(sessionHistory) == 0 || t != sessionHistory[len(sessionHistory)-1]) {
		sessionHistory = append(sessionHistory, t)
	}
	
	return t, nil
}
