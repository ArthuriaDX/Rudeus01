// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoDB

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
	"time"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoDB/dbTypes"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// the WotoClient struct which will act as a client in the program
type WotoClient struct {
	client          *mongo.Client
	ctx             *context.Context
	settings        interfaces.WSettings
	MainHostAddress string
	DataBase1       dbTypes.DATABASE
	ctxCancel       func()
}

//opt             *options.ClientOptions

// EMPTY is an empty string which is used in many many parts of the program :|
const EMPTY string = ""

// the _client of the woto data base package
var _client *WotoClient

// GenerateClient will generate a new data base client for you,
// if and only if there is no client defined in the program.
// before calling this function, create a new settings in the app.
func GenerateClient() (wa.RESULT, interfaces.WClient) {
	if _client != nil {
		return wa.SUCCESS, _client
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
	return wa.SUCCESS, _client
}

// clientError function will trigger an error and close the
// program asap!
func clientError() {
	log.Fatal("\nUnknown Error in woto data base!")
}

func (_c *WotoClient) PingClientDB(_report bool) {
	s := time.Now()
	_error := _c.client.Ping(*_c.ctx, nil)
	if _error != nil {
		if _report {
			_c.settings.SendSudo("PING ERR" + _error.Error())
		}
		log.Fatal(_error)
	} else {
		pingTime := time.Since(s)
		_c.settings.SendSudo(pingTime.String())
	}
}

// _setClient, will set the client field of the WotoClient struct
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
	if wotoStrings.IsEmpty(&_s) {
		log.Fatal(wotoValues.HOST_ADDRESS_NOTSET)
	}
	_c.MainHostAddress = _s

	_s = os.Getenv(wotoValues.DATEBASE1_KEY)
	if wotoStrings.IsEmpty(&_s) {
		log.Fatal(wotoValues.DATABASE1_NOTSET)
	}
	_c.DataBase1 = dbTypes.DATABASE(_s)
}

// getDataBase returns a database with the specified DATABASE value.
// use DefaultDatabase to get a default database
func (_c *WotoClient) getDataBase(database dbTypes.DATABASE) *mongo.Database {
	switch database {
	case DefaultDatabase:
		_d := _c.client.Database(string(database))
		return _d
	default:
		return nil
	}
}

// getCollection will give you the specified collection.
func (_c *WotoClient) getCollection(database dbTypes.DATABASE,
	collection dbTypes.COLLECTION) *mongo.Collection {
	return _c.getDataBase(database).Collection(string(collection))
}

// DeleteCollection will delete the specified collection.
func (_c *WotoClient) DeleteCollection(database dbTypes.DATABASE,
	collection dbTypes.COLLECTION) wa.RESULT {
	_collection := _c.getCollection(database, collection)
	if _collection == nil {
		return wa.FAILED
	}
	_err := _collection.Drop(*_c.ctx)
	if _err != nil {
		return wa.FAILED
	}
	return wa.SUCCESS
}

// getRAWS will give you the raw data.
func (_c *WotoClient) getRAWS(database dbTypes.DATABASE, collection dbTypes.COLLECTION) (wa.RESULT, []bson.M) {
	_collection := _c.getCollection(database, collection)
	if _collection == nil {
		clientError()
		return wa.FAILED, nil
	}
	_cursor, _cursorError := _collection.Find(*_c.ctx, bson.M{})
	if _cursorError != nil || _cursor == nil {
		return wa.FAILED, nil
	}
	var _raws []bson.M
	_ = _cursor.All(*_c.ctx, &_raws)
	return wa.SUCCESS, _raws
}

func (_c *WotoClient) Destroy() {
	_c.ctxCancel()

}
func (_c *WotoClient) GetWotoConfiguration() wa.RESULT {
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
	return wa.SUCCESS
}
func (_c *WotoClient) ResetWotoConfiguration() wa.RESULT {
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
func (_c *WotoClient) ResetUsersCollection() wa.RESULT {
	_collection := _c.getCollection(MainDataBase, UsersCollection)
	_err := _collection.Drop(*_c.ctx)
	if _err != nil {
		return wa.FAILED
	}
	return wa.SUCCESS
}
func (_c *WotoClient) CreateNewConfiguration() wa.RESULT {
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
	return wa.SUCCESS
}
func (_c *WotoClient) UpdateLastUIDConfiguration(_index uint8) wa.RESULT {
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
	return wa.SUCCESS
}
func (_c *WotoClient) CreateNewAccount() wa.RESULT {
	//_collection := _c.getCollection(MainDataBase, UsersCollection)
	//_R, _ := _collection.InsertOne(*_c.ctx, _a.getForServer())
	//if _R != nil {
	//	return SUCCESS
	//}
	return wa.FAILED
}
func (_c *WotoClient) AccountsLength() int {
	_r, _raws := _c.getRAWS(MainDataBase, UsersCollection)
	if _r != wa.SUCCESS {
		return 0
	}
	return len(_raws)
}
func (_c *WotoClient) FindAccount(_username, _pass *string) wa.RESULT {
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
	return wa.FAILED
}
func (_c *WotoClient) DeleteAccount() wa.RESULT {
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
	return wa.FAILED
}
func (_c *WotoClient) UpdateAccount() wa.RESULT {
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
	return wa.FAILED
}
func (_c *WotoClient) UpdateAccountOnlineToken() wa.RESULT {
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
	return wa.FAILED
}
