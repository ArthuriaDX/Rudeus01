// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoDB

import "github.com/ALiwoto/rudeus01/wotoPacks/wotoDB/dbTypes"

const (
	// the main data base of the program
	MainDataBase dbTypes.DATABASE = "bgrsx76y6idxwuk"
	// the DefaultDatabase of the program
	DefaultDatabase = MainDataBase
)
const (
	// ConfigurationCollection is for configuration of the bot.
	ConfigurationCollection dbTypes.COLLECTION = "Configuration"

	SUDOs_INFO dbTypes.COLLECTION = "SudoInfo"

	// UsersCollection is the main user collection :/
	UsersCollection dbTypes.COLLECTION = "Users"
	// SecuredCollection
	SecuredCollection dbTypes.COLLECTION = "Secured"
	// PlayerInfoCollection
	PlayerInfoCollection dbTypes.COLLECTION = "PlayerInfo"
	// OnlineTokenCollection
	OnlineTokenCollection dbTypes.COLLECTION = "OnlineTokens"
)

const (
	MAIN_SUDO_KEY = "MainSudo"
	SUDO_LIST_KEY = "SudoList"
)
