package wotoActions

/*
* sudo rm /var/lib/mongodb/mongod.lock
* sudo mongod --repair
* sudo service mongod start
* sudo service mongod status
**/
import (
	"context"
	"log"
	"os"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// the DATABASE type
type DATABASE string
type COLLECTION string

// the WotoClient struct which will act as a client in the program
type WotoClient struct {
	client          *mongo.Client
	ctx             *context.Context
	opt             *options.ClientOptions
	settings        *appSettings.AppSettings
	MainHostAddress string
	DataBase1       DATABASE
	ctxCancel       func()
}

const (
	// the main data base of the program
	MainDataBase DATABASE = "bgrsx76y6idxwuk"
	// the DefaultDatabase of the program
	DefaultDatabase = MainDataBase
)
const (
	// ConfigurationCollection is for configuration of the status of the players
	ConfigurationCollection COLLECTION = "Configuration"
	// UsersCollection is the main user collection :/
	UsersCollection COLLECTION = "Users"
	// SecuredCollection
	SecuredCollection COLLECTION = "Secured"
	// PlayerInfoCollection
	PlayerInfoCollection COLLECTION = "PlayerInfo"
	// OnlineTokenCollection
	OnlineTokenCollection COLLECTION = "OnlineTokens"
)

// EMPTY is an empty string which is used in many many parts of the program :|
const EMPTY string = ""

// the _client of the woto data base package
var _client *WotoClient

// GenerateClient will generate a new data base client for you,
// if and only if there is no client defined in the program.
// NOTICE: in the new version, I added the defining of the
// _wotoConfig in this function, so you don't need to use
// WotoClient.GetWotoConfiguration() in the main() function
func GenerateClient() (RESULT, *WotoClient) {
	if _client != nil {
		return SUCCESS, _client
	}
	_c := WotoClient{client: nil}
	_c._setHosts()
	_c._setSettings()
	_c._setClient()
	_client = &_c
	// uncomment this if and only if you need to create a new
	// wotoConfiguration.
	//_client.CreateNewConfiguration()
	//	_result, _wotoConfig = _client.GetWotoConfiguration()
	//	if _result == FAILED {
	//		_result, _wotoConfig = _client.ResetWotoConfiguration()
	//	}
	//	return _result
	//}
	return SUCCESS, _client
}

// clientError function will trigger an error and close the
// program asap!
func clientError() {
	log.Fatal("\nUnknown Error in woto data base!")
}

func (_c *WotoClient) PingClientDB(_report bool) {
	_error := _c.client.Ping(*_c.ctx, nil)
	_s := appSettings.GetExisting()
	if _error != nil {
		if _report {
			SendSudo("PING ERR"+_error.Error(), _s)
		}
		log.Fatal(_error)
	} else {
		SendSudo("PING DONE - ", _s)
	}
}

// setClient, will set the client field of the WotoClient struct
func (_c *WotoClient) _setClient() {
	//_c.opt = options.Client().ApplyURI(hostAddress)
	//client, _err := mongo.Connect(context.TODO(), _c.opt)
	//if _err != nil {
	//	log.Fatal(_err)
	//	return
	//}
	//_c.client = client

	//ctx, _cancelFunc := context.WithTimeout(context.Background(), 10 * time.Second)
	ctx := context.Background()
	_c.ctx = &ctx
	//_c.ctxCancel = _cancelFunc
	//clientOptions := options.Client().ApplyURI(hostAddress)
	clientOptions := options.Client().ApplyURI(_c.MainHostAddress)
	client, _err := mongo.NewClient(clientOptions)
	if _err != nil {
		log.Fatal(_err)
		return
	}
	if client == nil {
		clientError()
		return
	}
	_c.client = client
	_err = _c.client.Connect(ctx)
	if _err != nil {
		log.Fatal(_err)
	}
}

func (_c *WotoClient) _setSettings() {
	_s := appSettings.GetExisting()
	if _s == nil {
		log.Fatal(wotoValues.CLIENT_SETTINGS_NIL)
	}
	_c.settings = _s
}

func (_c *WotoClient) _setHosts() {
	_s := os.Getenv(wotoValues.HOST_ADDRESS1_KEY)
	if wotoSecurity.IsEmpty(&_s) {
		log.Fatal(wotoValues.HOST_ADDRESS_NOTSET)
	}
	_c.MainHostAddress = _s

	_s = os.Getenv(wotoValues.DATEBASE1_KEY)
	if wotoSecurity.IsEmpty(&_s) {
		log.Fatal(wotoValues.DATABASE1_NOTSET)
	}
	_c.DataBase1 = DATABASE(_s)
}

// getDataBase returns a database with the specified DATABASE value.
// use DefaultDatabase to get a default database
func (_c *WotoClient) getDataBase(database DATABASE) *mongo.Database {
	switch database {
	case DefaultDatabase:
		_d := _c.client.Database(string(database))
		return _d
	default:
		return nil
	}
}

// getCollection will give you the specified collection.
func (_c *WotoClient) getCollection(database DATABASE, collection COLLECTION) *mongo.Collection {
	return _c.getDataBase(database).Collection(string(collection))
}

// DeleteCollection will delete the specified collection.
func (_c *WotoClient) DeleteCollection(database DATABASE, collection COLLECTION) RESULT {
	_collection := _c.getCollection(database, collection)
	if _collection == nil {
		return FAILED
	}
	_err := _collection.Drop(*_c.ctx)
	if _err != nil {
		return FAILED
	}
	return SUCCESS
}

// getRAWS will give you the raw data.
func (_c *WotoClient) getRAWS(database DATABASE, collection COLLECTION) (RESULT, []bson.M) {
	_collection := _c.getCollection(database, collection)
	if _collection == nil {
		clientError()
		return FAILED, nil
	}
	_cursor, _cursorError := _collection.Find(*_c.ctx, bson.M{})
	if _cursorError != nil || _cursor == nil {
		return FAILED, nil
	}
	var _raws []bson.M
	_ = _cursor.All(*_c.ctx, &_raws)
	return SUCCESS, _raws
}

func (_c *WotoClient) destroy() {
	_c.ctxCancel()

}
func (_c *WotoClient) GetWotoConfiguration() RESULT {
	//_w := wotoConfiguration{
	//	UIDKeyName: [MaxUIDIndex]string{},
	//	LastUID:    [MaxUIDIndex]UID{},
	//}
	//_re, _raws := _c.getRAWS(MainDataBase, ConfigurationCollection)
	//if _re != SUCCESS {
	//	return FAILED, nil
	//}
	//_w.setUIDKeys()
	//for _i, _current := range _raws {
	//	_value, _ok := _current[UIDKeyName].(string)
	//	_w.LastUID[_i] = UID(_value)
	//	if !_ok {
	//		clientError()
	//		return FAILED, nil
	//	}
	//}
	return SUCCESS
}
func (_c *WotoClient) ResetWotoConfiguration() RESULT {
	//_collection := _c.getCollection(MainDataBase, ConfigurationCollection)
	//_err := _collection.Drop(*_c.ctx)
	//if _err != nil {
	//	panic(_err)
	//	return FAILED, nil
	//}
	//_re := _c.CreateNewConfiguration()
	//if _re != SUCCESS {
	//	clientError()
	//	return FAILED, nil
	//}
	return _c.GetWotoConfiguration()
}
func (_c *WotoClient) ResetUsersCollection() RESULT {
	_collection := _c.getCollection(MainDataBase, UsersCollection)
	_err := _collection.Drop(*_c.ctx)
	if _err != nil {
		return FAILED
	}
	return SUCCESS
}
func (_c *WotoClient) CreateNewConfiguration() RESULT {
	//_w := getNewWotoConfiguration()
	//_collection := _c.getCollection(MainDataBase, ConfigurationCollection)

	//var _b bson.D
	//for i := BaseUIDIndex - UIDIndexOffSet; i <= MaxUIDIndex-UIDIndexOffSet; i++ {
	//	_b = bson.D{
	//		{UIDIndexValue, _w.UIDKeyName[i]},
	//		{UIDKeyName, _w.LastUID[i]},
	//	}
	//	_R, _err := _collection.InsertOne(*_c.ctx, _b)
	//	if _R == nil || _err != nil {
	//		return FAILED
	//	}
	//}
	return SUCCESS
}
func (_c *WotoClient) UpdateLastUIDConfiguration(_index uint8) RESULT {
	//_collection := _c.getCollection(MainDataBase, ConfigurationCollection)
	//if _collection == nil {
	//	return FAILED
	//}
	//_cursor, _cursorError := _collection.Find(*_c.ctx, bson.M{})
	//if _cursorError != nil || _cursor == nil {
	//	return FAILED
	//}
	//_i := _index - UIDIndexOffSet
	// _b := bson.D{
	// {SetValue, bson.D{
	// {UIDIndexValue, _w.UIDKeyName[_i]},
	// {UIDKeyName, _w.LastUID[_i]},
	// }},
	// }
	// var _raws []bson.M
	// _ = _cursor.All(*_c.ctx, &_raws)
	//_current := _raws[_i]
	//_re, _err := _collection.UpdateOne(*_c.ctx, _current, _b)
	//if _re == nil || _err != nil {
	//	panic(_err)
	//	return FAILED
	//}
	return SUCCESS
}
func (_c *WotoClient) CreateNewAccount() RESULT {
	//_collection := _c.getCollection(MainDataBase, UsersCollection)
	//_R, _ := _collection.InsertOne(*_c.ctx, _a.getForServer())
	//if _R != nil {
	//	return SUCCESS
	//}
	return FAILED
}
func (_c *WotoClient) AccountsLength() int {
	_r, _raws := _c.getRAWS(MainDataBase, UsersCollection)
	if _r != SUCCESS {
		return 0
	}
	return len(_raws)
}
func (_c *WotoClient) FindAccount(_username, _pass *string) RESULT {
	//if isEmpty(_username) {
	//	return FAILED, nil
	//}
	//_r, _raws := _c.getRAWS(MainDataBase, UsersCollection)
	//if _r != SUCCESS {
	//	return FAILED, nil
	//}
	//for _, _current := range _raws {
	//	_uStrong := getStrong(_current[UserNameValue].(string))
	//	_userStrong := generateStrongString(*_username)
	//	_pStrong := getStrong(_current[PasswordValue].(string))
	//	_passStrong := generateStrongString(*_pass)
	//	if _uStrong.isEqual(&_userStrong) && _pStrong.isEqual(&_passStrong) {
	//		acc := Account{
	//			Username: &_uStrong,
	//			Password: &_pStrong,
	//			IsOwner:  _current[IsOwnerValue].(bool),
	//			IsAdmin:  _current[IsAdminValue].(bool),
	//		}
	//		return SUCCESS, &acc
	//	}
	//}
	return FAILED
}
func (_c *WotoClient) DeleteAccount() RESULT {
	//_collection := _c.getCollection(MainDataBase, UsersCollection)
	//if _collection == nil {
	//	return FAILED
	//}
	//_cursor, _cursorError := _collection.Find(*_c.ctx, bson.M{})
	//if _cursorError != nil || _cursor == nil {
	//	return FAILED
	//}
	//var _raws []bson.M
	//_ = _cursor.All(*_c.ctx, &_raws)
	//for _, _current := range _raws {
	//	_uStrong := getStrong(_current[UserNameValue].(string))
	//	_pStrong := getStrong(_current[PasswordValue].(string))
	//	if _uStrong.isEqual(_a.Username) && _pStrong.isEqual(_a.Password) {
	//		_re, _err := _collection.DeleteOne(*_c.ctx, _current)
	//		if _re == nil || _err != nil {
	//			return FAILED
	//		}
	//		return SUCCESS
	//	}
	//}
	return FAILED
}
func (_c *WotoClient) UpdateAccount() RESULT {
	//_collection := _c.getCollection(MainDataBase, UsersCollection)
	//if _collection == nil {
	//	return FAILED
	//}
	//_cursor, _cursorError := _collection.Find(*_c.ctx, bson.M{})
	//if _cursorError != nil || _cursor == nil {
	//	return FAILED
	//}
	//var _raws []bson.M
	//_ = _cursor.All(*_c.ctx, &_raws)
	//for _, _current := range _raws {
	//	_uStrong := getStrong(_current[UserNameValue].(string))
	//	_pStrong := getStrong(_current[PasswordValue].(string))
	//	if _uStrong.isEqual(_a.Username) && _pStrong.isEqual(_a.Password) {
	//		_re, _err := _collection.UpdateOne(*_c.ctx, _current, _a.getForServerUpdate())
	//		if _re == nil || _err != nil {
	//			panic(_err)
	//			return FAILED
	//		}
	//		return SUCCESS
	//	}
	//}
	return FAILED
}
func (_c *WotoClient) UpdateAccountOnlineToken() RESULT {
	//_collection := _c.getCollection(MainDataBase, OnlineTokenCollection)
	//if _collection == nil {
	//	return FAILED
	//}
	//_cursor, _cursorError := _collection.Find(*_c.ctx, bson.M{})
	//if _cursorError != nil || _cursor == nil {
	//	panic(_cursorError)
	//	return FAILED
	//}
	//var _raws []bson.M
	//_ = _cursor.All(*_c.ctx, &_raws)
	//for _, _current := range _raws {
	//	_uStrong := getStrong(_current[UserNameValue].(string))
	//	_pStrong := getStrong(_current[PasswordValue].(string))
	//	if _uStrong.isEqual(_a.Username) && _pStrong.isEqual(_a.Password) {
	//		_re, _err := _collection.UpdateOne(*_c.ctx, _current, _a.getForServerTokenUpdate())
	//		if _re == nil || _err != nil {
	//			panic(_err)
	//			return FAILED
	//		}
	//		return SUCCESS
	//	}
	//}
	return FAILED
}
func (_c *WotoClient) CheckOnlineToken() (RESULT, bool) {
	//_re, _raws := _c.getRAWS(MainDataBase, OnlineTokenCollection)
	//if _re != SUCCESS {
	//	return FAILED, false
	//}
	//for _, _current := range _raws {
	//	_value, _ok := _current[PlayerUIDValue].(string)
	//	if !_ok {
	//		return FAILED, false
	//	}
	//	_uid := UID(_value)
	//	if _a.PlayerUID.isEqual(&_uid) {
	//		_value, _ok = _current[OnlineTokenValue].(string)
	//		if !_ok {
	//			return FAILED, false
	//		}
	//		_token := wotoToken(_value)
	//		if _a.LastToken.isEqual(&_token) {
	//			return SUCCESS, true
	//		} else {
	//			return SUCCESS, false
	//		}
	//	}
	//}
	return FAILED, false
}
