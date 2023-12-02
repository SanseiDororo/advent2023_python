package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	total := 0

	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read each line
	for scanner.Scan() {
		line := scanner.Text()

		// Extract digits from each line
		var digits []rune
		for _, char := range line {
			if char >= '0' && char <= '9' {
				digits = append(digits, char)
			}
		}

		// Check if digits exist in the line
		if len(digits) > 0 {
			// Convert first and last digits to integers
			firstDigit, _ := strconv.Atoi(string(digits[0]))
			lastDigit, _ := strconv.Atoi(string(digits[len(digits)-1]))

			// Create the calibration number from the first and last digits
			calNum, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))

			// Add the calibration number to the total
			total += calNum
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(total)
}
