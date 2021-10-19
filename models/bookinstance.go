package models

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookInstance struct {
	field.DefaultField `bson:",inline"`
	Book               primitive.ObjectID `bson:"book" json:"book"`
	Imprint            string             `bson:"imprint" json:"imprint"`
	Status             string             `bson:"status" json:"status"`
	Due_back           primitive.DateTime `bson:"due_back" json:"due_back"`
	Books              []Book             `bson:"books" json:"books"`
}
