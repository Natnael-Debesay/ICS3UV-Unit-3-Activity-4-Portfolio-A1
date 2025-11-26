/**
 * @author Natnael Debesay
 * @version 1.0.0
 * @date 2025-11-26
 * @fileoverview This program prompts the user to enter a number between 1 and 10 to indicate how hungry they are and what they should do.
 */

// variables
let hunger;
let hungerNum;

// input
hunger = prompt("Enter your hunger on a scale from 1 - 10. ") || ("No value entered!");
hungerNum = parseInt(hunger);

// output
if (hungerNum <= 5){
  console.log("You are not really that hungry.");
} else if (hungerNum > 5) {
  console.log("Please eat!");
}

console.log("\nDone.")