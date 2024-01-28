package main

import (
	// "context"
	"fmt"
	// "io/ioutil"
	"log"
	// "net/http"
	// "strings"

	// "github.com/zmb3/spotify/v2"
	"google.golang.org/api/youtube/v3"
	// "github.com/percit/Yt2Spotify/helpers"
	"github.com/percit/Yt2Spotify/yt"
	// "github.com/percit/Yt2Spotify/spotifyAuth"
)
const (
	ytPlaylistID = "PLxKqTrK2bWoe6mmjllCaFMj09F1s3FCRy"
	spotifyPlaylist = "1ppfDotGdYV1FKRGkD9tZM"
)


func main() {


	
	ytClient := yt.GetClient(youtube.Youtube)
	_, err := youtube.New(ytClient)
	
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}








	// //SPOTIFY AUTH
	// http.HandleFunc("/callback", spotifyAuth.CompleteAuth)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Got request for:", r.URL.String())
	// })
	// client, err := spotifyAuth.AuthenticateSpotify()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//getting youtube playlist
	// songs, err := yt.GetYoutubePlaylistItems(ytPlaylistID, config, token)
	// if err != nil {
	// 	log.Fatalf("Unable to get playlist items: %v", err)
	// }

	// var unwantedSongs []string
	// var songsAddedToSpotify []string
	// fmt.Println("\nChoose from 0 to 9 to add a song to the playlist, or enter 11 to skip the song:\n")
	// for _, song := range songs {
	// 	fmt.Println("Youtube Song:" + song.Title)
	// 	results, err := client.Search(context.Background(), song.Title, spotify.SearchTypeTrack)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	if results.Tracks != nil {
	// 		fmt.Println("Songs:")
	// 		for i, track := range results.Tracks.Tracks {
	// 			if i >= 9 {
	// 				break
	// 			}
	// 			fmt.Printf("Track %d: %s - %s\n", i+1, track.Artists[0].Name, track.Name)
	// 		}
	// 	}

	// 	var userReply int
	// 	_, errUserReply := fmt.Scan(&userReply)
	// 	if errUserReply != nil {
	// 		fmt.Println("Invalid input:", errUserReply)
	// 		continue
	// 	}

	// 	if helpers.IsNumberInRange(userReply) {
	// 		fmt.Println("You chose song:", results.Tracks.Tracks[userReply].Name, results.Tracks.Tracks[userReply].Artists[0].Name)
	// 		client.AddTracksToPlaylist(context.Background(), spotify.ID(spotifyPlaylist), results.Tracks.Tracks[userReply].ID)
	// 		songsAddedToSpotify = append(songsAddedToSpotify, song.Title)
	// 		yt.DeleteSongByID(ytPlaylistID, song.PlaylistItemID, config, token)//TODO this should actually delete from yt playlist the ones that are on spotify
	// 	} else if userReply == 11 {
	// 		fmt.Println("You chose to skip this song")
	// 		unwantedSongs = append(unwantedSongs, song.Title)
	// 	} else {
	// 		fmt.Println("Something is wrong")
	// 	}
	// 	fmt.Println("\n")
	// }

	// contentUnwantedSongs := strings.Join(unwantedSongs, "\n")
	// errFile := ioutil.WriteFile("unwantedSongs.txt", []byte(contentUnwantedSongs), 0644)
	// if err != nil {
	// 	log.Fatalf("Unable to write to file: %v", errFile)
	// }
	// contentSongsAddedToSpotify := strings.Join(unwantedSongs, "\n")
	// errFile = ioutil.WriteFile("songsAddedToSpotify.txt", []byte(contentSongsAddedToSpotify), 0644)
	// if err != nil {
	// 	log.Fatalf("Unable to write to file: %v", errFile)
	// }

	fmt.Println("Song list saved to 'unwantedSongs.txt' and 'songsAddedToSpotify.txt'")
}


