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

func (s *server) templateResponse(w http.ResponseWriter, r *http.Request, t string, d interface{}) {

	// Execute the home template.
	logrus.Infof("%s %s\n\t╚ %v", r.Method, r.URL.String(), http.StatusOK)
	if err := s.tpl.ExecuteTemplate(w, "home.html", d); err != nil {
		erro(w, r, err, http.StatusInternalServerError)
		return
	}
}

func (s *server) Home() http.HandlerFunc {

	const tempName = "home.html"

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

			var pageToken string

			for {

				// Setup the playlist items call.
				call := s.ytSrv.PlaylistItems.List([]string{"snippet", "contentDetails"}).
					PlaylistId(r.FormValue("playlistID")).
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

						// Respond with a channel not found error.
						response.Error = "channel not found"
						s.templateResponse(w, r, tempName, response)
						return
					}

					// Internal server error.
					erro(w, r, err, http.StatusInternalServerError)
					return
				}

				// Go through the playlist items and each one to the response.
				for _, v := range res.Items {
					response.Playlist.Items = append(response.Playlist.Items, struct {
						Title        string
						ThumbnailURL string
					}{v.Snippet.Title, v.Snippet.Thumbnails.Medium.Url})
				}

				// If there is no next-page token, break free from the loop.
				if res.NextPageToken == "" {
					break
				}

				// Save the next-page token.
				pageToken = res.NextPageToken
			}
		}

		// Respond.
		s.templateResponse(w, r, tempName, response)
	}
}
