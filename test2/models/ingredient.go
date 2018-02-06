package models

import "github.com/globalsign/mgo/bson"

type Ingredient struct {
	ID           bson.ObjectId `bson:"_id" form:"-" json:"id"`
	Nume         string        `bson:"nume" form:"nume" json:"nume"`
	Calorii      float32       `bson:"calorii" form:"calorii" json:"calorii"`
	Proteine     float32       `bson:"proteine" form:"proteine" json:"proteine"`
	Lipide       float32       `bson:"lipide" form:"lipide" json:"lipide"`
	Carbohidrati float32       `bson:"carbohidrati" form:"carbohidrati" json:"carbohidrati"`
	Fibre        float32       `bson:"fibre" form:"fibre" json:"fibre"`
	Aproximari   float32       `bson:"aproximari" form:"aproximari" json:"aproximari"`
}
