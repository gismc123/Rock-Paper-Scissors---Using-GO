package main // Declare the main package

import (
	"bufio"   // For buffered input from the user
	"fmt"     // For formatted I/O like Println
	"math/rand" // For generating random numbers (bot's choice)
	"os"      // For OS-related functions, like reading input
	"strings" // For string manipulation (trimming, lowercasing)
	"time"    // For seeding the random number generator
)

func main() {
	fmt.Println("Starting the game of rock paper scissors..") // Print welcome message
	fmt.Println("First to 5 wins the game!")                  // Explain the winning condition

	rand.Seed(time.Now().UnixNano()) // Seed random number generator with current time

	// Array holding possible bot choices
	botAnswers := [3]string{"rock", "paper", "scissors"}

	yourScore := 0 // Initialize player's score
	botScore := 0  // Initialize bot's score
	drawScore := 0 // Initialize Draw Score

	// Print initial scores
	fmt.Printf("Bot score: %d \n", botScore)
	fmt.Printf("Your score: %d \n", yourScore)
	fmt.Printf("Draw score: %d \n", drawScore)

	reader := bufio.NewReader(os.Stdin) // Create a buffered reader to read user input

	for {
		fmt.Print("Choose your pick (rock, paper, scissors): ") // Prompt user for input

		answer, _ := reader.ReadString('\n')           // Read input until newline
		text := strings.TrimSpace(strings.ToLower(answer)) // Trim whitespace and convert to lowercase

		if text == "exit" { // If user types "exit", quit the game
			fmt.Println("Exiting the game...")
			break
		} else if text == "" { // If input is empty, warn user and continue loop
			fmt.Println("Empty input detected")
			continue
		} else if text != "rock" && text != "paper" && text != "scissors" { // Validate input
			fmt.Println("Please choose either rock, paper or scissors!")
			continue
		}

		// Generate a random index to pick bot's choice
		randomIndex := rand.Intn(len(botAnswers))
		botPick := botAnswers[randomIndex] // Bot's choice

		// Show choices
		fmt.Println("You chose " + text)
		fmt.Println("Bot chose " + botPick)

		// Determine the winner and update scores accordingly
		if botPick == "paper" && text == "rock" {
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "paper" && text == "scissors" {
			fmt.Println("You win!")
			yourScore++
		} else if botPick == "rock" && text == "scissors" {
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "rock" && text == "paper" {
			fmt.Println("You win!")
			yourScore++
		} else if botPick == "scissors" && text == "rock" {
			fmt.Println("You win!")
			yourScore++
		} else if botPick == "scissors" && text == "paper" {
			fmt.Println("You lose!")
			botScore++
		} else {
			// If none of the above, it's a draw
			fmt.Println("Draw!")
			drawScore++
		}

		// Print updated scores
		fmt.Printf("Bot score: %d \n", botScore)
		fmt.Printf("Your score: %d \n", yourScore)
		fmt.Printf("Draw score: %d \n", drawScore)

		// Check if either player reached 5 points to end the game
		if yourScore >= 5 {
			fmt.Println("You won the game! Congratulations!")
			break
		} else if botScore >= 5 {
			fmt.Println("You lose the game! Better luck next time!")
			break
		}
	}
}
