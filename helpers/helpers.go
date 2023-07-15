package helpers

import (
	"fmt"
)

func IsNumberInRange(num int) bool {
	return num >= 0 && num <= 9
}

func GetUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}
	return input, nil
}