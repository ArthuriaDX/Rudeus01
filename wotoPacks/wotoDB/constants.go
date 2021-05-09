// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoDB

import "github.com/ALiwoto/rudeus01/wotoPacks/wotoDB/dbTypes"

// databases' name used in the app.
const (
	MainDataBase dbTypes.DATABASE = "bgrsx76y6idxwuk" // wConfig
	PatDataBase  dbTypes.DATABASE = "bqvkw7yvsxypyuj" // pat
)
const (
	// ConfigurationCollection is for configuration of the bot.
	ConfigurationCollection dbTypes.COLLECTION = "Configuration"

	SUDOs_INFO dbTypes.COLLECTION = "SudoInfo"

	PatListCollection  dbTypes.COLLECTION = "Pat"
	HPatListCollection dbTypes.COLLECTION = "HPat"
	// UsersCollection is the main user collection :/
	UsersCollection dbTypes.COLLECTION = "Users"
	// SecuredCollection
	SecuredCollection dbTypes.COLLECTION = "Secured"
	// PlayerInfoCollection
	PlayerInfoCollection dbTypes.COLLECTION = "PlayerInfo"
	// OnlineTokenCollection
	OnlineTokenCollection dbTypes.COLLECTION = "OnlineTokens"
)

// wotoConfiguration constants.
const (
	MAIN_SUDO_KEY = "MainSudo"
	SUDO_LIST_KEY = "SudoList"
)

const (
	CONFIG_CLIENT dbTypes.DB_INDEX = 1
	PAT_CLIENT    dbTypes.DB_INDEX = 2
)
