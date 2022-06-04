package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// create a secret number
	secret := getRandomNumber()
	//fmt.Println(secret)

	for matching := false; !matching; {
		// Get user input
		guess := getUserInput()
		fmt.Println(secret, guess)

		// make comparison
		matching = isMatching(secret, guess)
		fmt.Println(matching)

	}
}

func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Your guess is too big")
		return false
	} else if guess < secret {
		fmt.Println("Your guess is too small")
		return false
	} else {
		fmt.Println("Your got it!!!")
		return true
	}
}

func getUserInput() int {

	fmt.Print("Please input your guess: ")
	var input int

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Failed to parse your input")
	} else {
		fmt.Println("You guessed: ", input)
	}

	return input

}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11 // gets random num from 0 to 10
}
