package yt

import (
	"context"
	"fmt"


	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"golang.org/x/oauth2"
)

type PlaylistInfo struct {
	Title          string
	PlaylistItemID string
}

func GetYoutubePlaylistItems(playlistID string, config *oauth2.Config, token *oauth2.Token) ([]PlaylistInfo, error) {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		return nil,err
	}

	var playlistItems []PlaylistInfo
	nextPageToken := ""
	for {
		playlistCall := youtubeService.PlaylistItems.List([]string{"snippet", "id"}).
			PlaylistId(playlistID).
			MaxResults(50).
			PageToken(nextPageToken)

		playlistResponse, err := playlistCall.Do()
		if err != nil {
			return nil,err
		}

		for _, playlistItem := range playlistResponse.Items {
			info := PlaylistInfo {
				Title: playlistItem.Snippet.Title,
				PlaylistItemID: playlistItem.Id,
			}
			playlistItems = append(playlistItems, info)
		}

		nextPageToken = playlistResponse.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	return playlistItems, nil
}

func DeleteSongByID(playlistID string, songID string,config *oauth2.Config, token *oauth2.Token) (error) {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		return err
	}

	if songID != "" {
        err = youtubeService.PlaylistItems.Delete(songID).Do()
        if err != nil {
			fmt.Println("Error deleting playlist item ", err)
			return err
        } else {
            fmt.Println("Video deleted successfully.")
        }
    } else {
        fmt.Println("Video not found in the playlist.")
    }
	return nil
}

// func GetToken(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error) {
// 	// Check if a token already exists in the cache or perform the OAuth 2.0 authorization flow
// 	tokenFile := "token.json"
// 	token, err := tokenFromFile(tokenFile)
// 	if err != nil {
// 		token = getTokenFromWeb(config)
// 		saveToken(tokenFile, token)
// 	}
// 	return token, nil
// }

// func tokenFromFile(file string) (*oauth2.Token, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	tok := &oauth2.Token{}
// 	err = json.NewDecoder(f).Decode(tok)
// 	return tok, err
// }

// func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
// 	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
// 	fmt.Printf("Go to the following link in your browser:\n%v\n", authURL)

// 	var code string
// 	fmt.Scanln(&code)

// 	tok, err := config.Exchange(context.Background(), code)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve token from web: %v", err)
// 	}
// 	return tok
// }

// func saveToken(file string, token *oauth2.Token) {
// 	fmt.Printf("Saving credential file to: %s\n", file)
// 	f, err := os.Create(file)
// 	if err != nil {
// 		log.Fatalf("Unable to cache OAuth token: %v", err)
// 	}
// 	defer f.Close()
// 	json.NewEncoder(f).Encode(token)
// }

// Error deleting playlist item  googleapi: Error 401: API keys are not supported by this API. Expected OAuth2 access token or other authentication credentials that assert a principal. See https://cloud.google.com/docs/authentication
// Details:
// [
//   {
//     "@type": "type.googleapis.com/google.rpc.ErrorInfo",
//     "domain": "googleapis.com",
//     "metadata": {
//       "method": "youtube.api.v3.V3DataPlaylistItemService.Delete",
//       "service": "youtube.googleapis.com"
//     },
//     "reason": "CREDENTIALS_MISSING"
//   }
// ]

// More details:
// Reason: required, Message: Login Required.