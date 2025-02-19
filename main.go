package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
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
	reader := bufio.NewReader(os.Stdin)
	word := wordbank[rand.Intn(len(wordbank))]
	displayStr := strings.Repeat("_", len(word))
	stage := 0
	for stage < len(hangmanStages) {
		var indexes []int
		fmt.Println(hangmanStages[stage])
		fmt.Println(displayStr)
		fmt.Print("\nGuess a letter: ")
		char, _, err := reader.ReadRune()
		if err != nil {
			println(err.Error())
			return false
		}
		fmt.Println("Word: " + word)
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
