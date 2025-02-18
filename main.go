package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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
	length := len(hangmanStages)
	word := wordbank[rand.Intn(len(wordbank))]
	var userin string
	for i := 0; i < length-1; i++ {
		fmt.Println(hangmanStages[i])
		fmt.Print("\nGuess: ")
		_, err := fmt.Scan(&userin)
		if err != nil {
			println(err.Error())
			return false
		}
		fmt.Println("User Input: " + userin)
		if userin == word {
			return true
		}
	}
	fmt.Println(hangmanStages[length-1])
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
