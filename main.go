package main

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

)

var _count int
var _alreadyRunning bool

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Println("we are here!" + port)
	router := gin.New()
	router.Use(gin.Logger())
	//router.LoadHTMLGlob("templates/*.tmpl.html")
	//router.Static("/static", "static")

	router.GET("/", AnswerClient)

	router.Run(":" + port)
}

func AnswerClient(c *gin.Context) {
	//c.HTML(http.StatusOK, "index.tmpl.html", nil)
	if !_alreadyRunning {
		_alreadyRunning = true
	} else {
		return
	}
	_s := c.Request.Header.Get("test")
	_count++
	c.String(http.StatusOK, "%v %v; countr: %v", "yes!", _s, _count)
	wotoActions.RunBot()
}