package helpers

import (
	"fmt"
	"flag"
)

var (
	GoogleApiToken       string
	SpotifyClientID      string
	SpotifyClientSecret  string
)

func init() {
	flag.StringVar(&GoogleApiToken, "g", "", "Google API Token")
	flag.StringVar(&SpotifyClientID, "c", "", "Spotify Client ID")
	flag.StringVar(&SpotifyClientSecret, "s", "", "Spotify Client Secret")
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