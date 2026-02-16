package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			words := cleanInput(userInput)

			if len(words) == 0 {
				continue
			}

			cmd := words[0]
			fmt.Println("Your command was:", cmd)

		}
	}
}

func cleanInput(test string) []string {
	words := strings.Fields(test)

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	return words
}
