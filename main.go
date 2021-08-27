package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	youtubeAPIKey string
	spotifyToken  string

	youtubePlaylistID string
	spotifyPlaylistID string
)

func getENV(k string) string {
	env, ok := os.LookupEnv(k)
	if !ok {
		logrus.Fatalf("%s is required", k)
	}
	return env
}

func main() {

	// Get api keys.
	youtubeAPIKey = getENV("YOUTUBE_API_KEY")
	spotifyToken = getENV("SPOTIFY_TOKEN")

	// Get playlist IDs.
	youtubePlaylistID = getENV("YOUTUBE_PLAYLIST_ID")
	spotifyPlaylistID = getENV("SPOTIFY_PLAYLIST_ID")

	// Setup youtube service.
	ytSrv, err := youtube.NewService(context.Background(), option.WithAPIKey(youtubeAPIKey))
	if err != nil {
		logrus.Fatalf("cannot create youtube service: %v", err)
	}

	// Setup spotify service.
	t := &oauth2.Token{AccessToken: spotifyToken}
	c := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(t))
	spCli := spotify.NewClient(c)

	// Empty the playlist.
	if err := emptySpotifyPlaylist(spCli); err != nil {
		logrus.Fatalf("cannot empty spotify playlist: %v", err)
	}

	// Get the youtube playlist.
	songs, err := getYoutubePlaylistSongs(ytSrv)
	if err != nil {
		logrus.Fatalf("cannot get youtube playlist: %v", err)
	}

	// Loop over the youtube playlist songs.
	var songsToAdd []spotify.ID
	for _, v := range songs {

		// Remove leading ' - Topic' from the channel name.

		// Search for the song in spotify.
		searchResult, err := spCli.Search(v.Snippet.Title, spotify.SearchTypeTrack)
		if err != nil {
			log.Fatalf("cannot search spotify for %s: %v", v.Snippet.Title, err)
		}

		// If the spotify artist and youtube channel name match up, add the song.

		// Else, let the user decide what song to add.
	}

	// Add the songs to the spotify playlist.
	_, err = spCli.AddTracksToPlaylist(spotify.ID(spotifyPlaylistID), songsToAdd...)
	if err != nil {
		logrus.Fatalf("cannot add spotify songs: %v", err)
	}

	// Done!
	fmt.Println("Done!")
}
