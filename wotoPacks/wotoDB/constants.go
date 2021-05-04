package wotoDB

import "github.com/ALiwoto/rudeus01/wotoPacks/wotoDB/dbTypes"

const (
	// the main data base of the program
	MainDataBase dbTypes.DATABASE = "bgrsx76y6idxwuk"
	// the DefaultDatabase of the program
	DefaultDatabase = MainDataBase
)
const (
	// ConfigurationCollection is for configuration of the status of the players
	ConfigurationCollection dbTypes.COLLECTION = "Configuration"
	// UsersCollection is the main user collection :/
	UsersCollection dbTypes.COLLECTION = "Users"
	// SecuredCollection
	SecuredCollection dbTypes.COLLECTION = "Secured"
	// PlayerInfoCollection
	PlayerInfoCollection dbTypes.COLLECTION = "PlayerInfo"
	// OnlineTokenCollection
	OnlineTokenCollection dbTypes.COLLECTION = "OnlineTokens"
)
