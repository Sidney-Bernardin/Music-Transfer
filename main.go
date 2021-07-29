package main

import (
	"context"
	"net/http"
	"os"

	"github.com/Sidney-Bernardin/MusicTransfer/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	API_KEY = os.Getenv("API_KEY")
	PORT    = os.Getenv("PORT")
)

func main() {

	// Setup the youtube service.
	ytSrv, err := youtube.NewService(context.Background(), option.WithAPIKey(API_KEY))
	if err != nil {
		logrus.Fatalf("cannot create youtube service: %v", err)
	}

	// Setup the server.
	s, err := server.NewServer(ytSrv)
	if err != nil {
		logrus.Fatalf("cannot create server: %s", err)
	}

	// Start the server.
	logrus.Infof("Listen and serveing on :%s ...\n", PORT)
	if err := http.ListenAndServe(":"+PORT, s); err != nil {
		logrus.Fatalf("listen and serve failed: %v", err)
	}
}
