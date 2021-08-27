package main

import (
	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

func emptySpotifyPlaylist(client spotify.Client) error {

	const operation = "spotifyEmptyPlaylist"

	// Get the playlist songs.
	songs, err := client.GetPlaylistTracks(spotify.ID(spotifyPlaylistID))
	if err != nil {
		return errors.Wrap(err, operation)
	}

	// If there are no songs to remove, return.
	if len(songs.Tracks) == 0 {
		return nil
	}

	// Add the songs to the removal list.
	var songsToRemove []spotify.ID
	for page := 1; ; page++ {

		// Add the song IDs of this page.
		for _, v := range songs.Tracks {
			songsToRemove = append(songsToRemove, v.Track.ID)
		}

		// Go to the next page.
		if err = client.NextPage(songs); err == spotify.ErrNoMorePages {

			// If there are no more pages, break free.
			if err == spotify.ErrNoMorePages {
				break
			}

			return errors.Wrap(err, operation)
		}
	}

	_, err = client.RemoveTracksFromPlaylist(spotify.ID(spotifyPlaylistID), songsToRemove...)
	return errors.Wrap(err, operation)
}
