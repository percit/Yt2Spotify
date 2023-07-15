package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/zmb3/spotify/v2"

	"github.com/percit/Yt2Spotify/helpers"
	"github.com/percit/Yt2Spotify/yt"
	"github.com/percit/Yt2Spotify/spotifyAuth"
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

func main() {
	ytPlaylistID, err := getUserInput("Type YouTube playlist from which you wish to export songs:")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	spotifyPlaylist, err := getUserInput("Type Spotify playlist to which you wish to import songs:")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	//YOUTUBE STUFF
	songs, err := yt.GetYoutubePlaylistItems(ytPlaylistID, GoogleApiToken)
	if err != nil {
		log.Fatalf("Unable to get playlist items: %v", err)
	}

	//SPOTIFY STUFF
	http.HandleFunc("/callback", spotifyAuth.CompleteAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	client, err := spotifyAuth.AuthenticateSpotify()
	if err != nil {
		log.Fatal(err)
	}

	var unwantedSongs []string
	fmt.Println("\nChoose from 0 to 9 to add a song to the playlist, or enter 11 to skip the song:\n")
	for _, song := range songs {
		fmt.Println("Youtube Song:" + song)
		results, err := client.Search(context.Background(), song, spotify.SearchTypeTrack)
		if err != nil {
			log.Fatal(err)
		}

		if results.Tracks != nil {
			fmt.Println("Songs:")
			for i, track := range results.Tracks.Tracks {
				if i >= 9 {
					break
				}
				fmt.Printf("Track %d: %s - %s\n", i+1, track.Artists[0].Name, track.Name)
			}
		}

		var userReply int
		_, errUserReply := fmt.Scan(&userReply)
		if errUserReply != nil {
			fmt.Println("Invalid input:", errUserReply)
			continue
		}

		if helpers.IsNumberInRange(userReply) {
			fmt.Println("You chose song:", results.Tracks.Tracks[userReply].Name, results.Tracks.Tracks[userReply].Artists[0].Name)
			client.AddTracksToPlaylist(context.Background(), spotify.ID(spotifyPlaylist), results.Tracks.Tracks[userReply].ID)
		} else if userReply == 11 {
			fmt.Println("You chose to skip this song")
			unwantedSongs = append(unwantedSongs, song) //TODO this should actually delete from yt playlist the ones that are on spotify
		} else {
			fmt.Println("Something is wrong")
		}
		fmt.Println("\n")
	}

	content := strings.Join(unwantedSongs, "\n")
	errFile := ioutil.WriteFile("song_list.txt", []byte(content), 0644)
	if err != nil {
		log.Fatalf("Unable to write to file: %v", errFile)
	}

	fmt.Println("Song list saved to 'song_list.txt'")
}
