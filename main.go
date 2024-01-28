package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	// "os"
	// "encoding/json"

	"github.com/zmb3/spotify/v2"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"

	"github.com/percit/Yt2Spotify/helpers"
	"github.com/percit/Yt2Spotify/yt"
	"github.com/percit/Yt2Spotify/spotifyAuth"
)
const (
	redirectURL  = "http://localhost:3000/login"
	ytPlaylistID = "PLxKqTrK2bWoe6mmjllCaFMj09F1s3FCRy"
	spotifyPlaylist = "1ppfDotGdYV1FKRGkD9tZM"
)

var (
	token *oauth2.Token
	googleOAuthConfig *oauth2.Config
	oauthStateString  = "random" // a random string for CSRF protection
)
func main() {
	config := &oauth2.Config{
		ClientID:     helpers.GoogleClientID,
		ClientSecret: helpers.GoogleClientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{youtube.YoutubeForceSslScope},
		Endpoint:     google.Endpoint,
	}

	// ctx := context.Background()

	// Get OAuth 2.0 token
	// token, err := yt.GetToken(ctx, config)
	// if err != nil {
	// 	log.Fatalf("Unable to get OAuth 2.0 token: %v", err)
	// }
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)

	fmt.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

	//SPOTIFY AUTH
	http.HandleFunc("/callback", spotifyAuth.CompleteAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	client, err := spotifyAuth.AuthenticateSpotify()
	if err != nil {
		log.Fatal(err)
	}

	//getting youtube playlist
	songs, err := yt.GetYoutubePlaylistItems(ytPlaylistID, config, token)
	if err != nil {
		log.Fatalf("Unable to get playlist items: %v", err)
	}

	var unwantedSongs []string
	var songsAddedToSpotify []string
	fmt.Println("\nChoose from 0 to 9 to add a song to the playlist, or enter 11 to skip the song:\n")
	for _, song := range songs {
		fmt.Println("Youtube Song:" + song.Title)
		results, err := client.Search(context.Background(), song.Title, spotify.SearchTypeTrack)
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
			songsAddedToSpotify = append(songsAddedToSpotify, song.Title)
			yt.DeleteSongByID(ytPlaylistID, song.PlaylistItemID, config, token)//TODO this should actually delete from yt playlist the ones that are on spotify
		} else if userReply == 11 {
			fmt.Println("You chose to skip this song")
			unwantedSongs = append(unwantedSongs, song.Title)
		} else {
			fmt.Println("Something is wrong")
		}
		fmt.Println("\n")
	}

	contentUnwantedSongs := strings.Join(unwantedSongs, "\n")
	errFile := ioutil.WriteFile("unwantedSongs.txt", []byte(contentUnwantedSongs), 0644)
	if err != nil {
		log.Fatalf("Unable to write to file: %v", errFile)
	}
	contentSongsAddedToSpotify := strings.Join(unwantedSongs, "\n")
	errFile = ioutil.WriteFile("songsAddedToSpotify.txt", []byte(contentSongsAddedToSpotify), 0644)
	if err != nil {
		log.Fatalf("Unable to write to file: %v", errFile)
	}

	fmt.Println("Song list saved to 'unwantedSongs.txt' and 'songsAddedToSpotify.txt'")
}


func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Println("Invalid oauth state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	tok, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("Code exchange failed: %v\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	token = tok

	fmt.Println("Access Token:", token.AccessToken)

	// You can use the 'token' here or save it for later use.

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

func redirectionHandler(w http.ResponseWriter, req *http.Request) {
	code := req.FormValue("code")
	if code != "" {
		w.Write([]byte("Copy the following auth code and paste it to the terminal:\n\n" + code))
	} else {
		w.Write([]byte("Error: " + req.FormValue("error")))
	}
}

// func main() {
// 	srv := &http.Server{Addr: ":8090"}
// 	http.HandleFunc("/", redirectionHandler)
// 	go func() {
// 		err := srv.ListenAndServe()
// 		fmt.Println(err)
// 	}()

// 	// Only needed in first try
// 	fmt.Println("Creating OAuth 2 token file...")
// 	err := CreateOauthToken()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	} else {
// 		fmt.Println("Create the oauth 2 token file at ~/.credentials successfully")
// 	}
// 	srv.Close()
// }

func CreateOauthToken() error {
	b, err := os.ReadFile("client_secret.json")
	if err != nil {
		return err
	}

	// If modifying the scope, delete your previously saved credentials
	// at ./client_secret.json
	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope, youtube.YoutubeUploadScope, youtube.YoutubeScope)
	if err != nil {
		return err
	}

	// Use a redirect URI like this for a web app. The redirect URI must be a
	// valid one for your OAuth2 credentials.
	config.RedirectURL = "http://localhost:8090"

	tok := getTokenFromWeb(config)

	_, err = saveToken("yt-token.json", tok)
	return err
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) (string, error) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(token)

	return fmt.Sprintf("Saving credential file to: %s", file), err
}