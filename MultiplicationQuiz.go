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
	// Create a new random source with a seed based on current time
	source := rand.NewSource(time.Now().UnixNano())
	// Create a new random number generator
	r := rand.New(source)

	scanner := bufio.NewScanner(os.Stdin)

	var correct int
	for i := 0; i < 5; i++ {
		num1 := r.Intn(12) + 1
		num2 := r.Intn(12) + 1
		fmt.Printf("What is %d * %d? ", num1, num2)

		scanner.Scan()
		input, _ := strconv.Atoi(scanner.Text())

		if input == num1*num2 {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Incorrect. The answer is", num1*num2)
		}
	}

	fmt.Printf("You got %d out of 5 correct.\n", correct)
}
