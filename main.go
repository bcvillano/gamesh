package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

var wordbank = []string{"redteam", "shell", "service", "downtime"}
var hangmanStages = []string{
	`
         -----
         |   |
             |
             |
             |
             |
        =======
        `,
	`
         -----
         |   |
         O   |
             |
             |
             |
        =======
        `,
	`
         -----
         |   |
         O   |
         |   |
             |
             |
        =======
        `,
	`
         -----
         |   |
         O   |
        /|   |
             |
             |
        =======
        `,
	`
         -----
         |   |
         O   |
        /|\  |
             |
             |
        =======
        `,
	`
         -----
         |   |
         O   |
        /|\  |
        /    |
             |
        =======
        `,
	`
         -----
         |   |
         O   |
        /|\  |
        / \  |
        =======
        `,
}

func playHangman() bool {
	word := wordbank[rand.Intn(len(wordbank))]
	displayStr := strings.Repeat("_", len(word))
	var userin string
	for i := 0; i < len(hangmanStages)-1; i++ {
		fmt.Println(hangmanStages[i])
		fmt.Println(displayStr)
		fmt.Print("\nGuess: ")
		_, err := fmt.Scan(&userin)
		if err != nil {
			println(err.Error())
			return false
		}
		fmt.Println("User Input: " + userin)
		fmt.Println("Word: " + word)
		if userin == word {
			return true
		} else {
			for j := 0; j < len(word); j++ {
				if word[j] == userin[j] {
					//Strings in Go are immutable so we need to make a new one to replace old displayStr
					runes := []rune(displayStr)
					runes[j] = unicode.ToLower(rune(word[j]))
					displayStr = string(runes)
				}
			}
		}
	}
	fmt.Println(hangmanStages[len(hangmanStages)-1])
	fmt.Println("You lose! Try again...\n\n")
	return false
}

func main() {
	for {
		gameWon := playHangman()
		if gameWon {
			break
		}
	}
	fmt.Println("Launching bash...")
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
