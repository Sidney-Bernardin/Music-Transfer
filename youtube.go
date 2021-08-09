package main

import (
	"errors"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/youtube/v3"
)

var (
	errChannelNotFound = errors.New("channel not found")
)

func getYTPlaylist() ([]*youtube.PlaylistItem, error) {

	var ret []*youtube.PlaylistItem
	var pageToken string

	// Get the Youtube playlist.
	for {

		// Setup the playlist items call.
		call := ytSrv.PlaylistItems.List([]string{"snippet", "contentDetails"}).
			PlaylistId(youtubePlaylistID).
			MaxResults(50)

		// Add the page token if its set.
		if pageToken != "" {
			call.PageToken(pageToken)
		}

		// Do the call.
		res, err := call.Do()
		if err != nil {

			// Check for a channel not found error.
			if err.(*googleapi.Error).Errors[0].Reason == "channelNotFound" {
				return nil, errChannelNotFound
			}

			return nil, err
		}

		// Go through the playlist items and add each one to tplData.
		for _, v := range res.Items {
			ret = append(ret, v)
		}

		// If there is no next-page token, break free from the loop.
		if res.NextPageToken == "" {
			break
		}

		// Save the next-page token.
		pageToken = res.NextPageToken
	}

	return ret, nil
}
