package database

import "github.com/qiniu/qmgo"

var DB *qmgo.Database

var AuthorCollection *qmgo.Collection
var BookCollection *qmgo.Collection
var BookInstanceCollection *qmgo.Collection
var GenreCollection *qmgo.Collection
var SettingCollection *qmgo.Collection
var UserCollection *qmgo.Collection

func InitDatabase(client *qmgo.Client) {
	DB = client.Database("my_database")
	AuthorCollection = DB.Collection("authors")
	BookCollection = DB.Collection("books")
	BookInstanceCollection = DB.Collection("bookinstances")
	GenreCollection = DB.Collection("genres")
	SettingCollection = DB.Collection("settings")
	UserCollection = DB.Collection("users")
}
