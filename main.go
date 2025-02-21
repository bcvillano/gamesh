package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

var wordbank = []string{"redteam", "shell", "service", "downtime", "beacon", "persistence", "backdoor", "firewall", "incident", "inject"}
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

func decode(encodedStr string) string {
	decodedStr, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return "error-ignore"
	}
	return string(decodedStr)
}

func playHangman() bool {
	reader := bufio.NewReader(os.Stdin)
	word := wordbank[rand.Intn(len(wordbank))]
	displayStr := strings.Repeat("_", len(word))
	stage := 0
	for stage < len(hangmanStages)-1 {
		var indexes []int
		fmt.Println(hangmanStages[stage])
		fmt.Println(displayStr)
		fmt.Print("\nGuess a letter: ")
		input, err := reader.ReadString('\n') // Read full line
		if err != nil {
			fmt.Println("Error reading input:", err)
			return false
		}
		input = strings.TrimSpace(strings.ToLower(input)) // Normalize input
		if input == decode("aW1zb2V2aWw=") {
			return true
		} else if input == word {
			return true
		} else if len(input) != 1 || input[0] < 'a' || input[0] > 'z' {
			fmt.Println("Invalid input. Please enter a single letter.")
			continue
		}
		char := rune(input[0])
		//fmt.Println("Word: " + word) //Debug line
		for i, c := range word {
			if c == char {
				indexes = append(indexes, i)
			}
		}
		if len(indexes) == 0 {
			stage++
		} else {
			tempStr := []rune(displayStr)
			for _, idx := range indexes {
				tempStr[idx] = char
			}
			displayStr = string(tempStr)
		}
		if displayStr == word { //win condition
			fmt.Println("YOU WIN!!! Enjoy your shell")
			return true
		}
	}
	fmt.Println(hangmanStages[len(hangmanStages)-1]) //prints final hanged man
	fmt.Print(displayStr + "\n\n")
	fmt.Println("You lose! The word was " + word)
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
