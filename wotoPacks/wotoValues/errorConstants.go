// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoValues

const (
	PORT_ERROR            = `Couldn't find PORT value.`
	TOKEN_ERROR           = `No token provided! Cannot run app!`
	INVALID_API           = `Error! Couldn't resolve API!`
	INVALID_ENGINE        = `Error! Agent has no permission to start!`
	ALREADY_RUNNING       = `Already running the app!`
	WOTO_CLIENT_ERROR     = `Couldn't generate woto client!`
	GET_WCONFIG_ERROR     = `Couldn't get raw in GetWotoConfiguration method!`
	GET_MAIN_SUDO_ERROR   = `Tried to get main sudo id from db, but it's zero!`
	GET_SUDO_LIST_ERROR   = `Tried to get sudo list from db, but it's Empty!`
	PAT_CLIENT_ERROR      = `Couldn't generate pat client!`
	GET_PAT_LIST_ERROR    = `Couldn't get raw in GetPatList method!`
	CLIENT_SETTINGS_NIL   = `Tried to set client settings, but it was nil!!`
	HOST_ADDRESS_NOTSET   = `Couldn't find main host address in the environment!`
	DATABASE1_NOTSET      = `Couldn't find DATABASE1 address in the environment!`
	RUDEUS_URL_NOTSET     = `Couldn't find Rudeus URL value in the environment!`
	SUDOLIST_SET_SETTINGS = `Tried to set sudo list, but couldn't split it!`
	SUDO_ADD_SETTINGS     = `Tried to add a sudo to sudoList db, but couldn't!`
)
