package yt

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type PlaylistInfo struct {
	Title          string
	PlaylistItemID string
}

func GetYoutubePlaylistItems(playlistID string, apiKey string) ([]PlaylistInfo, error) {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
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

func DeleteSongByID(playlistID string, songID string, apiKey string) (error) {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return err
	}



//tutaj musimy jakos wydobyc to id

	if songID != "" {
        err = youtubeService.PlaylistItems.Delete(songID).Do()
        if err != nil {
			fmt.Println("Error deleting playlist item")
			return err
            // if apiErr, ok := err.(*googleapi.Error); ok && apiErr.Code == http.StatusNotFound {
            //     fmt.Println("Video not found in the playlist.")
            // } else {
            //     log.Fatalf("Error deleting playlist item: %v", err)
            // }
        } else {
            fmt.Println("Video deleted successfully.")
        }
    } else {
        fmt.Println("Video not found in the playlist.")
    }
	return nil
}
