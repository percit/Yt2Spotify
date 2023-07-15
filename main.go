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
	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/percit/Yt2Spotify/helpers"
	"github.com/percit/Yt2Spotify/yt"
)

var (
	GoogleApiToken       string
	SpotifyClientID      string
	SpotifyClientSecret  string
	redirectURI          = "http://localhost:8080/callback"
	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), 
		spotifyauth.WithScopes(
		spotifyauth.ScopePlaylistModifyPublic,
		spotifyauth.ScopePlaylistModifyPrivate,), 
		spotifyauth.WithClientID(SpotifyClientID), 
		spotifyauth.WithClientSecret(SpotifyClientSecret))
	ch    = make(chan *spotify.Client)
	state = "abc123"
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
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	client, err := authenticateSpotify()
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
			unwantedSongs = append(unwantedSongs, song)
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




func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}

	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}

func authenticateSpotify() (*spotify.Client, error) {

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	client := <-ch
	_, err := client.CurrentUser(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to get Spotify user: %v", err)
	}

	return client, nil
}