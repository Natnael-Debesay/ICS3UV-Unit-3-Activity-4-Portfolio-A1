// Author: Natnael Debesay
// Version: 1.0.0
// Date: 2025-11-26
// Fileoverview: This program prompts the user to enter a number between 1 and 10 to indicate how hungry they are and what they should do.

package main

import (
	"fmt"
	"os"
)

func main() {
	var hungerNum int

	// input
	fmt.Print("Enter your hunger on a scale from 1 - 10: ")

	_, err := fmt.Scanf("%d", &hungerNum)

	if err != nil {
		if err.Error() == "expected integer" {
			fmt.Println("Error: Invalid input. Please enter a number.")
		} else {
			fmt.Println("No value entered!")
		}
		os.Exit(1) 
	}

	if hungerNum < 1 || hungerNum > 10 {
		fmt.Println("Please enter a number between 1 and 10.")
		os.Exit(1)
	}

	// output 
	if hungerNum <= 5 {
		fmt.Println("You are not really that hungry.")
	} else if hungerNum > 5 {
		fmt.Println("Please eat!")
	}

	fmt.Println("\nDone.")
}