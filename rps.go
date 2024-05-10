package main

import (
	"fmt"
	"log"
	"math/rand"
)

const (
	rock = iota
	paper
	scissor
)

func main() {
	fmt.Println("Let's play, chose Rock (0), Paper (1), scissor(2)" +
		"input number:")
	playRockPaperScissor()
}

func playRockPaperScissor() {
	answers := []int{rock, paper, scissor}
	userInput := userInput()
	//userInput := rand.Intn(len(answers))

	fmt.Printf("read user_input: %d-\n", userInput)

	compInput := rand.Intn(len(answers))
	fmt.Println("Computer choose", compInput)

	evaluateChoices(userInput, compInput)
}

func evaluateChoices(userInput int, compInput int) {
	if userInput == compInput {
		fmt.Printf("Doh, You %d vs Computer %d Play again!\n\n", userInput, compInput)
		main()
	} else if userInput == 0 && compInput == 2 {
		fmt.Println("You  won!")
	} else if userInput == 1 && compInput == 0 {
		fmt.Println("You  won!")
	} else if userInput == 2 && compInput == 1 {
		fmt.Println("You  won!")
	} else {
		fmt.Println("Computer won!")
	}
}

func userInput() int {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal(err)
	}
	if number < 0 || number > 2 {
		fmt.Println("Try again!")
		number = userInput()
	}
	return number
}
