package helpers

import (
	"fmt"
	"flag"
)

var (
	GoogleClientID       string
	GoogleClientSecret   string
	SpotifyClientID      string
	SpotifyClientSecret  string
)

func init() {
	flag.StringVar(&GoogleClientID, "gi", "", "Google ClientID")
	flag.StringVar(&GoogleClientSecret, "gs", "", "Google ClientSecret")
	flag.StringVar(&SpotifyClientID, "sc", "", "Spotify Client ID")
	flag.StringVar(&SpotifyClientSecret, "ss", "", "Spotify Client Secret")
	flag.Parse()
}

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