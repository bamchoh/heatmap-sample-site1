package main

func (s *server) routes() {
	s.router.GET("/", s.appIndex)
	s.router.GET("/settings", s.settings)
	s.router.GET("/page1", s.page1)
	s.router.GET("/page2", s.page2)
}
