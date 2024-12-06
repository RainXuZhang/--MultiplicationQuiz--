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
        fmt.Println("Welcome to the Multiplication Quiz!")

        // Get user input for number of questions
        numQuestions := getNumberOfQuestions()

        // Get user input for number ranges
        minRange, maxRange := getNumberRanges()

        // Create a new random source with a seed based on current time
        source := rand.NewSource(time.Now().UnixNano())
        // Create a new random number generator
        r := rand.New(source)

        scanner := bufio.NewScanner(os.Stdin)

        var correctAnswers int
        startTime := time.Now()

        for i := 0; i < numQuestions; i++ {
                questionStart := time.Now()
                num1, num2 := generateNumbers(minRange, maxRange, r)
                answer := num1 * num2

                fmt.Printf("What is %d * %d? \n", num1, num2)

                scanner.Scan()
                if err := scanner.Err(); err != nil {
                        fmt.Println("Error reading input:", err)
                        return
                }
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

                fmt.Printf("Time taken for this question: %.2fs\n", time.Since(questionStart).Seconds())
        }

        endTime := time.Now()
        elapsedTime := endTime.Sub(startTime)

        fmt.Printf("You got %d out of %d correct (%.2f%% accuracy).\n", correctAnswers, numQuestions, float64(correctAnswers)/float64(numQuestions)*100)
        fmt.Printf("Total time taken: %.2fs\n", elapsedTime.Seconds())
}

func getNumberOfQuestions() int {
        fmt.Print("Enter the number of questions: ")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        if err := scanner.Err(); err != nil {
                fmt.Println("Error reading input:", err)
                return 5 // Default to 5 questions on error
        }
        numQuestions, err := strconv.Atoi(scanner.Text())
        if err != nil || numQuestions <= 0 {
                fmt.Println("Invalid input. Setting number of questions to 5 by default.")
                return 5
        }
        return numQuestions
}

func getNumberRanges() (int, int) {
        fmt.Print("Enter the minimum range: ")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        if err := scanner.Err(); err != nil {
                fmt.Println("Error reading input:", err)
                return 1, 10 // Default to range 1-10 on error
        }
        minRange, err := strconv.Atoi(scanner.Text())
        if err != nil || minRange <= 0 {
                fmt.Println("Invalid minimum range. Setting minimum range to 1 by default.")
                minRange = 1
        }

        fmt.Print("Enter the maximum range: ")
        scanner.Scan()
        if err := scanner.Err(); err != nil {
                fmt.Println("Error reading input:", err)
                return 1, 10 // Default to range 1-10 on error
        }
        maxRange, err := strconv.Atoi(scanner.Text())
        if err != nil || maxRange <= minRange {
                fmt.Println("Invalid maximum range. Setting maximum range to 10 by default.")
                maxRange = 10
        }

        return minRange, maxRange
}

func generateNumbers(minRange, maxRange int, r *rand.Rand) (int, int) {
        return r.Intn(maxRange-minRange+1) + minRange, r.Intn(maxRange-minRange+1) + minRange
}
