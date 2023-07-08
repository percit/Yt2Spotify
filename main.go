package main

import (
	"flag"
	"fmt"
	"context"
	"io/ioutil"
	"log"
	"strings"

	// "google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	//github.com/percit/Yt2Spotify
)

var GoogleApiToken string
var GoogleClientID string
var SpotifyApiToken string

func init() {
	flag.StringVar(&GoogleApiToken, "g", "", "Google API Token")
// 	flag.StringVar(&GoogleClientID, "c", "", "Google Client ID")
// 	flag.StringVar(&SpotifyApiToken, "s", "", "Spotify API Token")
	flag.Parse()
}

func main() {
	fmt.Println("Type Youtube playlist from which you wish to export songs")
	var ytPlaylistID string
	_, err1 := fmt.Scan(&ytPlaylistID)
	if err1 != nil {
		fmt.Println("Error reading input:", err1)
		return
	}
	fmt.Println("Input string:", ytPlaylistID)

	fmt.Println("Type Spotify playlist to which you wish to import songs")
	var spotifyPlaylist string
	_, err2 := fmt.Scan(&spotifyPlaylist)
	if err2 != nil {
		fmt.Println("Error reading input:", err2)
		return
	}
	fmt.Println("Input string:", spotifyPlaylist)
	


	// ytPlaylistID := "PLxKqTrK2bWod_CaZ0JZ7twpdW6xJPLnzc"
	songs, err := getYoutubePlaylistItems(ytPlaylistID, GoogleApiToken)
	if err != nil {
		log.Fatalf("Unable to get playlist items: %v", err)
	}
	var unwantedSongs []string


	fmt.Println("Choose from 1 to 5 to add song to playlist, use 6 to skip the song")
	for _, song := range songs {
		fmt.Println(song)
		
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
				unwantedSongs = append(unwantedSongs, song)
			}
			default:{
					fmt.Println("Something is wrong")
			}
		}
	}

	content := strings.Join(unwantedSongs, "\n")
	errFile := ioutil.WriteFile("song_list.txt", []byte(content), 0644)
	if err != nil {
		log.Fatalf("Unable to write to file: %v", errFile)
	}

	fmt.Println("Song list saved to 'song_list.txt'")
}

func getYoutubePlaylistItems(playlistID string, apiKey string) ([]string, error) {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	var playlistItems []string

	nextPageToken := ""
	for {
		playlistCall := youtubeService.PlaylistItems.List([]string{"snippet"}).
		PlaylistId(playlistID).
		MaxResults(50).
		PageToken(nextPageToken)

		playlistResponse, err := playlistCall.Do()
		if err != nil {
			return nil, err
		}

		for _, playlistItem := range playlistResponse.Items {
			playlistItems = append(playlistItems, playlistItem.Snippet.Title)
		}

		nextPageToken = playlistResponse.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	return playlistItems, nil
}

//PLxKqTrK2bWod_CaZ0JZ7twpdW6xJPLnzc