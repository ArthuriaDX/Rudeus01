package appSettings

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _alreadyRunning bool
var _settings *AppSettings

type AppSettings struct {
	port   string
	router *gin.Engine
	botAPI *tgbotapi.BotAPI
	RunNew func()
}

func GetSettings(_port string, _runNew func()) *AppSettings {
	if _settings != nil {
		return GetExisting()
	}
	if wotoSecurity.IsEmpty(&_port) {
		return nil
	}
	settings := AppSettings{}
	settings._setPort(_port)
	settings._setRouter()
	_settings = &settings
	return _settings
}

func GetExisting() *AppSettings {
	return _settings
}

func (_s *AppSettings) SetAPI(_api *tgbotapi.BotAPI) {
	_s.botAPI = _api
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

func (_s *AppSettings) GetAPI() *tgbotapi.BotAPI {
	return _s.botAPI
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
