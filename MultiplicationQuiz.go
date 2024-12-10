package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

type QuizStats struct {
	TotalQuestions int
	CorrectAnswers int
	CurrentStreak  int
	MaxStreak      int
}

func main() {
	fmt.Println("Created by Rain Zhang\nVersion 1.0.2")
	fmt.Println("Welcome to the Multiplication Quiz!")

	// Create a buffered scanner for more efficient input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024) // Increase buffer size

	// Get user input for number of questions
	numQuestions := getNumberOfQuestions(scanner)

	// Get user input for number ranges
	minRange, maxRange := getNumberRanges(scanner)

	// Initialize quiz statistics
	stats := &QuizStats{}
	stats.TotalQuestions = numQuestions

	// Start the quiz
	startTime := time.Now()
	runQuiz(scanner, numQuestions, minRange, maxRange, stats)

	// Display final results
	displayResults(stats, startTime)
}

func runQuiz(scanner *bufio.Scanner, numQuestions, minRange, maxRange int, stats *QuizStats) {
	for i := 0; i < numQuestions; i++ {
		// Generate cryptographically secure random numbers
		num1 := generateSecureRandomNumber(minRange, maxRange)
		num2 := generateSecureRandomNumber(minRange, maxRange)
		answer := num1 * num2

		// Track question performance
		questionCorrect := askQuestion(scanner, num1, num2, answer, stats)

		// Update streak
		if questionCorrect {
			stats.CurrentStreak++
			stats.MaxStreak = max(stats.MaxStreak, stats.CurrentStreak)
			stats.CorrectAnswers++
		} else {
			stats.CurrentStreak = 0
		}
	}
}

func askQuestion(scanner *bufio.Scanner, num1, num2, answer int, stats *QuizStats) bool {
	maxAttempts := 3
	attempts := 0

	for attempts < maxAttempts {
		fmt.Printf("What is %d * %d? ", num1, num2)

		// Provide hints based on attempt number
		if attempts > 0 {
			provideHint(num1, num2, answer, attempts)
		}

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// Allow user to skip with hint
		if input == "hint" {
			attempts++
			continue
		}

		// Validate input
		userAnswer, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			attempts++
			continue
		}

		// Check answer
		if userAnswer == answer {
			fmt.Println("Correct!")
			return true
		}

		fmt.Println("Incorrect.")
		attempts++
	}

	fmt.Printf("The correct answer was %d * %d = %d\n", num1, num2, answer)
	return false
}

func provideHint(num1, num2, answer, attempt int) {
	switch attempt {
	case 1:
		fmt.Println("Hint 1: Range Check")
		fmt.Printf("The answer is between %d and %d\n",
			min(num1*num2, num2*num1),
			max(num1*num2, num2*num1))
	case 2:
		fmt.Println("Hint 2: Multiplication Breakdown")
		fmt.Printf("Hint: %d is one of the factors\n",
			[]int{num1, num2}[attempt%2])
	}
}

func generateSecureRandomNumber(min, max int) int {
	// Use cryptographically secure random number generation
	delta := max - min + 1
	bigDelta := big.NewInt(int64(delta))

	// Generate a cryptographically secure random number
	n, err := rand.Int(rand.Reader, bigDelta)
	if err != nil {
		// Fallback to time-based random if secure generation fails
		return min + time.Now().Nanosecond()%delta
	}

	return min + int(n.Int64())
}

func getNumberOfQuestions(scanner *bufio.Scanner) int {
	fmt.Print("Enter the number of questions (default 5, max 50): ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	if input == "" {
		return 5
	}

	numQuestions, err := strconv.Atoi(input)
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

func getNumberRanges(scanner *bufio.Scanner) (int, int) {
	fmt.Print("Enter the minimum range (default 1): ")
	scanner.Scan()
	minInput := strings.TrimSpace(scanner.Text())
	minRange := 1
	if minInput != "" {
		min, err := strconv.Atoi(minInput)
		if err == nil && min > 0 {
			minRange = min
		}
	}

	fmt.Print("Enter the maximum range (default 10): ")
	scanner.Scan()
	maxInput := strings.TrimSpace(scanner.Text())
	maxRange := 10
	if maxInput != "" {
		max, err := strconv.Atoi(maxInput)
		if err == nil && max >= minRange {
			maxRange = max
		}
	}

	return minRange, maxRange
}

func displayResults(stats *QuizStats, startTime time.Time) {
	elapsedTime := time.Since(startTime)

	fmt.Println("\n--- Quiz Results ---")
	fmt.Printf("Total Questions: %d\n", stats.TotalQuestions)
	fmt.Printf("Correct Answers: %d\n", stats.CorrectAnswers)
	fmt.Printf("Accuracy: %.2f%%\n",
		float64(stats.CorrectAnswers)/float64(stats.TotalQuestions)*100)
	fmt.Printf("Max Streak: %d\n", stats.MaxStreak)
	fmt.Printf("Total Time: %.2f seconds\n", elapsedTime.Seconds())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
