package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var target int = rand.Intn(100)
	var guess int
	var totalGuesses int = 0

	fmt.Println("Welcome to Guess The Number!!!")

	for true {
		fmt.Print("Please guess a number: ")
		fmt.Scanln(&guess)

		fmt.Printf("You guessed %v\n", guess)

		if guess == target {
			fmt.Printf("Congratulations, you won the game with %v guesses\n", totalGuesses)
			return
		}

		if guess > target {
			fmt.Println("The target is smaller")
		} else {
			fmt.Println("The target is bigger")
		}

		totalGuesses++

	}
	fmt.Println("Please guess a number: ")
}
