package appSettings

import "github.com/gin-gonic/gin"

var _alreadyRunning bool
var _settings *AppSettings

type AppSettings struct {
	port string
	router *gin.Engine
}

func GetSettings(_port string) *AppSettings {
	if _settings != nil {
		return _settings
	}
	settings := AppSettings{}
	settings._setPort(_port)
	settings._setRouter()
	_settings = &settings
	return _settings
}

func (_s *AppSettings) _setPort(_port string) {
	_s.port = _port
}

func (_s *AppSettings) _setRouter() {
	_s.router = gin.New()
}

func (_s *AppSettings) GetPort() string {
	return _s.port
}

func (_s *AppSettings) GetRouter() *gin.Engine {
	return _s.router
}

func App_enter() {
	_alreadyRunning = true
}

func App_exit() {
	_alreadyRunning = false
}

func IsRunning() bool {
	return _alreadyRunning
}