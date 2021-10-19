package models

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	field.DefaultField `bson:",inline"`
	Title              string               `bson:"title" json:"title"`
	Author             primitive.ObjectID   `bson:"author" json:"author"`
	Summary            string               `bson:"summary" json:"summary"`
	Isbn               string               `bson:"isbn" json:"isbn"`
	Genre              []primitive.ObjectID `bson:"genre" json:"genre"`
	Authors            []Author             `bson:"authors" json:"authors"`
	Genres             []Genre              `bson:"genres" json:"genres"`
}
