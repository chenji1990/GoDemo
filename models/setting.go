package models

import "github.com/qiniu/qmgo/field"

type Setting struct {
	field.DefaultField `bson:",inline"`
	Name               string  `bson:"name" json:"name"`
	Value              float64 `bson:"value" json:"value"`
}
