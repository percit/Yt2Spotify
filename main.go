package main

import (
	// "flag"
	"fmt"
	//github.com/percit/Yt2Spotify
)

// const SPOTIFY_SCOPE = "user-read-private user-read-email playlist-read-private playlist-modify-private"
// const GOOGLE_SCOPE = "https://www.googleapis.com/auth/youtube.readonly"

var GoogleApiToken string
var GoogleClientID string
var SpotifyApiToken string

// func init() {
// 	flag.StringVar(&GoogleApiToken, "g", "", "Google API Token")
// 	flag.StringVar(&GoogleClientID, "c", "", "Google Client ID")
// 	flag.StringVar(&SpotifyApiToken, "s", "", "Spotify API Token")
// 	flag.Parse()
// }

func main() {
	fmt.Println("Type Youtube playlist from which you wish to export songs")
	var ytPlaylist string
	_, err1 := fmt.Scan(&ytPlaylist)
	if err1 != nil {
		fmt.Println("Error reading input:", err1)
		return
	}
	fmt.Println("Input string:", ytPlaylist)

	fmt.Println("Type Spotify playlist to which you wish to import songs")
	var spotifyPlaylist string
	_, err2 := fmt.Scan(&spotifyPlaylist)
	if err2 != nil {
		fmt.Println("Error reading input:", err2)
		return
	}
	fmt.Println("Input string:", spotifyPlaylist)

	fmt.Println("Your list of songs")
	fmt.Println("Choose from 1 to 5 to add song to playlist, use 6 to skip the song")



	var songNumbers int //HERE WE WOULD ADD HOW MANY SONG ARE THERE
	for songNumbers > 0{
		fmt.Println("HERE WOULD BE YT SONG NAME")
		fmt.Println("HERE WOULD BE SPOTIFY OUTPUT WITH 5 FIRST SONGS")
		
		var userReply int
		_, err := fmt.Scan(&userReply)
		if err != nil {
			fmt.Println("Invalid input:", err)
			return
		}
		switch(userReply) {
			case 1 : {
				fmt.Println("you chose song 1")
			}
			case 2: {
				fmt.Println("you chose song 2")
			}
			case 3: {
				fmt.Println("you chose song 3")
			}
			case 4: {
				fmt.Println("you chose song 4")
			}
			case 5: {
				fmt.Println("you chose song 5")
			}
			case 6: {
				fmt.Println("you chose to skip this song")
			}
			default:{
					fmt.Println("Something is wrong")
			}
		}
		songNumbers--
	}

	//HERE WE WOULD BE MAKING .TXT WITH SONGS NOT USEDz
}
