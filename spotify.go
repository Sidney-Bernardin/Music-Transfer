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

	// Go through the songs one page at a time.
	var songsToRemove []spotify.ID
	for {

		// Add the song IDs of theis page.
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

	// Setup the proper increment value.
	incr := 100
	if len(songsToRemove) < 100 {
		incr = len(songsToRemove)
	}

	// Delete the songs.
	for i := 0; i < len(songsToRemove); i += incr {
		_, err = client.RemoveTracksFromPlaylist(spotify.ID(spotifyPlaylistID), songsToRemove[i:i+incr]...)
		if err != nil {
			return errors.Wrap(err, operation)
		}
	}

	return nil
}
