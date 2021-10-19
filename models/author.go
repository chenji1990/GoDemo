package models

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	field.DefaultField `bson:",inline"`
	FirstName          string             `bson:"first_name" json:"first_name"`
	FamilyName         string             `bson:"family_name" json:"family_name"`
	DateOfBirth        primitive.DateTime `bson:"date_of_birth" json:"date_of_birth"`
	DateOfDeath        primitive.DateTime `bson:"date_of_death" json:"date_of_death"`
}
