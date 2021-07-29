package server

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func respond(w http.ResponseWriter, r *http.Request, err error, code int) {
	if err != nil {
		logrus.Warnf("%s %s\n\t╚ %v) %v", r.Method, r.URL.String(), code, err)
		http.Error(w, err.Error(), code)
		return
	}

	logrus.Infof("%s %s\n\t %v", r.Method, r.URL.String(), code)
}

func (s *server) Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res, err := s.ytSrv.Playlists.List([]string{"status"}).
			ChannelId("GoogleDevelopers").
			//Mine(true).
			Do()
		if err != nil {
			respond(w, r, err, http.StatusInternalServerError)
			return
		}

		// Execute the home template.
		respond(w, r, nil, http.StatusOK)
		if err := s.tpl.ExecuteTemplate(w, "home.html", res.Items[0]); err != nil {
			respond(w, r, err, http.StatusInternalServerError)
			return
		}
	}
}
