package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	youtubeAPIKey string
	spotifyAPIKey string

	youtubePlaylistID string
	spotifyPlaylistID string

	ytSrv *youtube.Service
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

	// Get playlist IDs.
	youtubePlaylistID = getENV("YOUTUBE_PLAYLIST_ID")

	var err error

	// Setup youtube service.
	ytSrv, err = youtube.NewService(context.Background(), option.WithAPIKey(youtubeAPIKey))
	if err != nil {
		logrus.Fatalf("cannot create youtube service: %v", err)
	}

	// Get youtube playlist.
	ytPlaylist, err := getYTPlaylist()
	if err != nil {
		logrus.Fatalf("cannot get youtube playlist: %v", err)
	}

	for _, v := range ytPlaylist {
		fmt.Println(v.Snippet.Title)
	}
}
