package models

import "github.com/qiniu/qmgo/field"

type Genre struct {
	field.DefaultField `bson:",inline"`
	Name               string `bson:"name" json:"name"`
}
