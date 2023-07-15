package spotifyAuth
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

var (
	redirectURI          = "http://localhost:8080/callback"
	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), 
		spotifyauth.WithScopes(
		spotifyauth.ScopePlaylistModifyPublic,
		spotifyauth.ScopePlaylistModifyPrivate,), 
		spotifyauth.WithClientID(""), 
		spotifyauth.WithClientSecret(""))
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func CompleteAuth(w http.ResponseWriter, r *http.Request) {
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

func AuthenticateSpotify() (*spotify.Client, error) {

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