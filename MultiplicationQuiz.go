//Created by Rain Zhang

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Difficulty level selection
	difficulty := selectDifficulty()

	// Create a new random source with a seed based on current time
	source := rand.NewSource(time.Now().UnixNano())
	// Create a new random number generator
	r := rand.New(source)

	scanner := bufio.NewScanner(os.Stdin)

	var totalQuestions, correctAnswers int
	startTime := time.Now()

	for i := 0; i < 5; i++ {
		questionStart := time.Now()
		num1, num2 := generateNumbers(difficulty, r)
		answer := num1 * num2

		fmt.Printf("What is %d * %d? (You have used approximately %d seconds so far) \n", num1, num2, int(time.Since(startTime))/1e9)

		scanner.Scan()
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			i-- // Repeat the question if input is invalid
			continue
		}

		if input == answer {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Println("Incorrect. The answer is", answer)
		}
		totalQuestions++
		fmt.Printf("Time taken for this question: %.2fs\n", time.Since(questionStart).Seconds())
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("You got %d out of %d correct (%.2f%% accuracy).\n", correctAnswers, totalQuestions, float64(correctAnswers)/float64(totalQuestions)*100)
	fmt.Printf("Total time taken: %.2fs\n", elapsedTime.Seconds())
}

func selectDifficulty() int {
	fmt.Println("Select difficulty level (enter the corresponding number):")
	fmt.Println("1. Easy (1-5)")
	fmt.Println("2. Medium (1-10)")
	fmt.Println("3. Hard (1-12)")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	difficulty, err := strconv.Atoi(scanner.Text())
	if err != nil || difficulty < 1 || difficulty > 3 {
		fmt.Println("Invalid selection. Choosing Easy level by default.")
		return 1
	}
	return difficulty
}

func generateNumbers(difficulty int, r *rand.Rand) (int, int) {
	switch difficulty {
	case 1:
		return r.Intn(5) + 1, r.Intn(5) + 1
	case 2:
		return r.Intn(10) + 1, r.Intn(10) + 1
	case 3:
		return r.Intn(12) + 1, r.Intn(12) + 1
	default:
		return 1, 1 // Should never reach here, but returning default values in case
	}
}
