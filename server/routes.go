package server

func (s *server) loadRoutes() {
	s.router.Handle("/", s.Home())
}
