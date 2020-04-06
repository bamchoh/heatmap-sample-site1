package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

type server struct {
	router gin.IRouter
}

func (s *server) appIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"user_id": "user name"})
}

func (s *server) settings(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.tmpl.html", gin.H{"user_id": "user name"})
}

func (s *server) page1(c *gin.Context) {
	c.HTML(http.StatusOK, "page1.tmpl.html", gin.H{"user_id": "user name"})
}

func (s *server) page2(c *gin.Context) {
	c.HTML(http.StatusOK, "page2.tmpl.html", gin.H{"user_id": "user name"})
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	engine := gin.New()
	store := memstore.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("mysession", store))
	engine.Use(gin.Logger())
	engine.LoadHTMLGlob("templates/*.tmpl.html")
	engine.Static("/assets", "./assets")
	engine.Static("/static", "static")

	svr := server{engine}

	svr.routes()

	engine.Run(":" + port)
}
