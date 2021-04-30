package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv(wotoValues.APP_PORT)

	if wotoSecurity.IsEmpty(&port) {
		log.Fatal(wotoValues.PORT_ERROR)
	}
	_settings := appSettings.GetSettings(port)
	runApp(_settings)
}

func runApp(_settings *appSettings.AppSettings) {
	router := _settings.GetRouter()
	router.Use(gin.Logger())
	router.GET(wotoValues.GET_SLASH, answerClient)

	_ = router.Run(wotoValues.HTTP_ADDRESS + _settings.GetPort())
}

func answerClient(c *gin.Context) {
	if !appSettings.IsRunning() {
		appSettings.App_enter()
	} else {
		return
	}
	wotoValues.DebugMode = true
	_token := os.Getenv(wotoValues.TOKEN_KEY)
	if wotoSecurity.IsEmpty(&_token) {
		c.String(http.StatusInternalServerError, wotoValues.TOKEN_ERROR)
		appSettings.App_exit()
	}
	wotoActions.RunBot(_token, c)
	wotoActions.RunNew()
}
