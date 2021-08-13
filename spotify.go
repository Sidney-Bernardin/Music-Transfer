package main

import (
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

func spotifyEmptyPlaylist(client spotify.Client) error {

	const operation = "spotifyEmptyPlaylist"

	// Get the playlist tracks.
	tracks, err := client.GetPlaylistTracks(spotify.ID(spotifyPlaylistID))
	if err != nil {
		return errors.Wrap(err, operation)
	}

	// Loop through the pages and add the IDs to a slice.
	var ids []spotify.ID
	for page := 1; ; page++ {

		// Add the track ids of this page to the ids slice.
		for _, v := range tracks.Tracks {
			ids = append(ids, v.Track.ID)
		}

		// Go to the next page.
		if err = client.NextPage(tracks); err == spotify.ErrNoMorePages {
			if err == spotify.ErrNoMorePages {
				break
			}

			return errors.Wrap(err, operation)
		}
	}

	_, err = client.RemoveTracksFromPlaylist(spotify.ID(spotifyPlaylistID), ids[0])
	return errors.Wrap(err, operation)
}
