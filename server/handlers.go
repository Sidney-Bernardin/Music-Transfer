package server

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/googleapi"
)

func erro(w http.ResponseWriter, r *http.Request, err error, code int) {
	if err != nil {
		logrus.Warnf("%s %s\n\t╚ %v) %v", r.Method, r.URL.String(), code, err)
		http.Error(w, err.Error(), code)
		return
	}
}

func (s *server) Home() http.HandlerFunc {

	var response struct {
		PlaylistID string
		Playlist   struct {
			ThumbnailURL string
			Items        []struct {
				Title        string
				ThumbnailURL string
			}
		}
		Error string
	}

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {

			// Get the playlists of the channel.
			res, err := s.ytSrv.PlaylistItems.List([]string{"snippet", "contentDetails"}).
				PlaylistId(r.FormValue("PlaylistID")).
				MaxResults(200).
				Do()

			if err != nil {

				// Check for a channel not found error.
				if err.(*googleapi.Error).Errors[0].Reason == "channelNotFound" {
					response.Error = "channel not found"
				} else {
					erro(w, r, err, http.StatusInternalServerError)
					return
				}
			} else {
				for _, v := range res.Items {
					response.Playlist.Items = append(response.Playlist.Items, struct {
						Title        string
						ThumbnailURL string
					}{v.Snippet.Title, v.Snippet.Thumbnails.Medium.Url})
				}
			}
		}

		// Execute the home template.
		logrus.Infof("%s %s\n\t╚ %v", r.Method, r.URL.String(), http.StatusOK)
		if err := s.tpl.ExecuteTemplate(w, "home.html", response); err != nil {
			erro(w, r, err, http.StatusInternalServerError)
			return
		}
	}
}
