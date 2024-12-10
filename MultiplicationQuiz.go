package main

import (
	"fmt"
	"math/rand"
	"time"
)

type QuizStats struct {
	TotalQuestions int
	CorrectAnswers int
	CurrentStreak  int
	MaxStreak      int
}

func main() {
	// Seed the random number generator once at the start
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the Multiplication Quiz!")

	// Preallocate a buffer for input
	inputBuffer := make([]byte, 20)

	// Get configuration with direct input
	numQuestions := getNumberOfQuestions(inputBuffer)
	minRange, maxRange := getNumberRanges(inputBuffer)

	// Initialize quiz statistics on stack
	var stats QuizStats
	stats.TotalQuestions = numQuestions

	// Start the quiz
	startTime := time.Now()
	runQuiz(numQuestions, minRange, maxRange, &stats)

	// Display final results
	displayResults(&stats, startTime)
}

func runQuiz(numQuestions, minRange, maxRange int, stats *QuizStats) {
	for i := 0; i < numQuestions; i++ {
		// Inline random number generation
		num1 := minRange + rand.Intn(maxRange-minRange+1)
		num2 := minRange + rand.Intn(maxRange-minRange+1)
		answer := num1 * num2

		// Track question performance
		if askQuestion(num1, num2, answer) {
			stats.CurrentStreak++
			if stats.CurrentStreak > stats.MaxStreak {
				stats.MaxStreak = stats.CurrentStreak
			}
			stats.CorrectAnswers++
		} else {
			stats.CurrentStreak = 0
		}
	}
}

func askQuestion(num1, num2, answer int) bool {
	fmt.Printf("What is %d * %d? ", num1, num2)

	// Efficient input reading
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return false
	}

	// Direct comparison
	if input == answer {
		fmt.Println("Correct!")
		return true
	}

	fmt.Printf("Incorrect. The correct answer is %d\n", answer)
	return false
}

func getNumberOfQuestions(buffer []byte) int {
	fmt.Print("Enter the number of questions (default 5, max 50): ")

	var numQuestions int
	_, err := fmt.Scanf("%d", &numQuestions)
	if err != nil || numQuestions <= 0 {
		fmt.Println("Invalid input. Using default of 5 questions.")
		return 5
	}

	if numQuestions > 50 {
		fmt.Println("Maximum of 50 questions allowed. Using 50.")
		return 50
	}

	return numQuestions
}

func getNumberRanges(buffer []byte) (int, int) {
	var minRange, maxRange int

	fmt.Print("Enter the minimum range (default 1): ")
	_, err := fmt.Scanf("%d", &minRange)
	if err != nil || minRange <= 0 {
		minRange = 1
	}

	fmt.Print("Enter the maximum range (default 10): ")
	_, err = fmt.Scanf("%d", &maxRange)
	if err != nil || maxRange < minRange {
		maxRange = 10
	}

	return minRange, maxRange
}

func displayResults(stats *QuizStats, startTime time.Time) {
	elapsedTime := time.Since(startTime)

	fmt.Println("\n--- Quiz Results ---")
	fmt.Printf("Total Questions: %d\n", stats.TotalQuestions)
	fmt.Printf("Correct Answers: %d\n", stats.CorrectAnswers)

	// Precompute float to avoid repeated division
	accuracy := float64(stats.CorrectAnswers) * 100 / float64(stats.TotalQuestions)
	fmt.Printf("Accuracy: %.2f%%\n", accuracy)

	fmt.Printf("Max Streak: %d\n", stats.MaxStreak)
	fmt.Printf("Total Time: %.2f seconds\n", elapsedTime.Seconds())
}
