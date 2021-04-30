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

	_ = router.Run(":" + port)
}

func AnswerClient(c *gin.Context) {
	//c.HTML(http.StatusOK, "index.tmpl.html", nil)
	if !_alreadyRunning {
		_alreadyRunning = true
	} else {
		return
	}
	_token := c.Request.Header.Get("TOKEN_BOT")
	_count++
	c.String(http.StatusOK, "%v %v; counter: %v", "yes! running with token:", _token, _count)
	wotoActions.RunBot(_token)
	c.String(http.StatusOK, "\n%v %v; counter: %v", "END! running with token:", _token, _count)
}