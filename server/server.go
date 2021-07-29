package server

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"google.golang.org/api/youtube/v3"
)

type server struct {
	router *mux.Router
	ytSrv  *youtube.Service
	tpl    *template.Template
}

func NewServer(ytSrv *youtube.Service) (*server, error) {

	const operation = "server.NewServer"

	// Parse the templates.
	tpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		return nil, errors.Wrap(err, operation)
	}

	// Setup the server.
	s := &server{
		router: mux.NewRouter(),
		ytSrv:  ytSrv,
		tpl:    tpl,
	}
	s.loadRoutes()

	return s, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
