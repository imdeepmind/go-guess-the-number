package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

func main() {
	// Clear the screen
	screen.Clear()
	screen.MoveTopLeft()

	// Making a guess
	var target int = rand.Intn(100)
	var guess int
	var totalGuesses int = 0

	// Opening messages
	color.Red("Welcome to Guess The Number!!!")
	color.Cyan("Please guess the number between 0 and 100")

	for true {
		// Asking for input from the screen
		fmt.Print("Your guess: ")
		fmt.Scanln(&guess)

		// Printing the typed number
		fmt.Printf("You guessed %v\n", guess)

		// Checking the guess with target
		if guess == target {
			color.Green("Congratulations, you won the game with %v guesses\n", totalGuesses)
			return
		}

		// Proving direction to the user based on guessed number and target
		if guess > target {
			fmt.Println("The target is smaller")
		} else {
			fmt.Println("The target is bigger")
		}

		// Maintaining a counter for counting guesses
		totalGuesses++
	}
	fmt.Println("Please guess a number: ")
}
