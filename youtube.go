package main

import (
	"github.com/pkg/errors"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/youtube/v3"
)

var (
	errChannelNotFound = errors.New("channel not found")
)

func getYoutubePlaylistSongs(srv *youtube.Service) ([]*youtube.PlaylistItem, error) {

	var (
		operation = "ytGetPlaylist"
		ret       []*youtube.PlaylistItem
	)

	// Get the playlist songs using a forever loop to keep track of page tokens.
	var nextPageToken string
	for {

		// Setup the call to get the playlist songs.
		call := srv.PlaylistItems.List([]string{"snippet", "contentDetails"}).
			PlaylistId(youtubePlaylistID).
			MaxResults(50)

		// Set the page token if its not empty.
		if nextPageToken != "" {
			call.PageToken(nextPageToken)
		}

		// Do the call.
		res, err := call.Do()
		if err != nil {

			// Check for a channel not found error.
			if err.(*googleapi.Error).Errors[0].Reason == "channelNotFound" {
				return nil, errChannelNotFound
			}

			return nil, errors.Wrap(err, operation)
		}

		// Add the songs to the response.
		for _, v := range res.Items {
			ret = append(ret, v)
		}

		// If there is no next page token, break free from the loop.
		if res.NextPageToken == "" {
			break
		}

		// Save the next page token.
		nextPageToken = res.NextPageToken
	}

	return ret, nil
}
