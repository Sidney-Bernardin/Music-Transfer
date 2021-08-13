package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
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

	/*
		// Setup youtube service.
		ytSrv, err := youtube.NewService(context.Background(), option.WithAPIKey(youtubeAPIKey))
		if err != nil {
			logrus.Fatalf("cannot create youtube service: %v", err)
		}
	*/

	// Setup spotify service.
	t := &oauth2.Token{AccessToken: spotifyToken}
	c := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(t))
	spCli := spotify.NewClient(c)

	// Get spotify playlist.
	if err := spotifyEmptyPlaylist(spCli); err != nil {
		logrus.Fatalf("cannot empty playlist: %v", err)
	}

	/*
		// Get youtube playlist.
		ytPlaylist, err := ytGetPlaylist(ytSrv)
		if err != nil {
			logrus.Fatalf("cannot get youtube playlist: %v", err)
		}
	*/
}
