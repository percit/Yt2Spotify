package yt

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func GetYoutubePlaylistItems(playlistID string, apiKey string) ([]string, error) {
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