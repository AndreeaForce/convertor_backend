package repositories

import (
	"fmt"
	"strings"

	"github.com/andreeaforce/test2/datasource/mongodb"
	"github.com/andreeaforce/test2/models"
	"github.com/globalsign/mgo/bson"
)

var (
	ingredienteCollection = "ingrediente"
)

// InsertIngredient inserts an ingredient in the database
func InsertIngredient(ingredient models.Ingredient) {
	ingredient.ID = bson.NewObjectId()
	db := mongodb.New()
	c := db.C(ingredienteCollection)
	err := c.Insert(&ingredient)
	if err != nil {
		fmt.Println(err)
	}
}

// GetAllIngredients returns all the ingredients from the database
func GetAllIngredients(limit int) (ingredients []models.Ingredient, err error) {
	db := mongodb.New()
	c := db.C(ingredienteCollection)
	defer db.Close()
	err = c.Find(nil).Limit(limit).All(&ingredients) /*.Limit(100).Skip(200)*/
	if err != nil {
		return
	}
	return
}

func GetIngredientByName(ingredientName string, page, limit int, sort string) (ingredients []models.Ingredient, count int, err error) {
	var skip = 0
	var ingredientNameArray []string
	query := bson.M{}
	sliceSearch := []bson.M{}
	if ingredientName != "" {
		ingredientNameArray = strings.Fields(ingredientName)

		for _, w := range ingredientNameArray {
			sliceSearch = append(sliceSearch, bson.M{"n": bson.M{"$regex": "^" + w}})
		}
		query["$and"] = sliceSearch
	} else {
		query = nil
	}
	if page > 1 {
		skip = (page-1)*limit + 1
	}
	db := mongodb.New()
	c := db.C(ingredienteCollection)
	defer db.Close()
	count, err = c.Find(query).Count()
	err = c.Find(query).Limit(limit).Skip(skip).Sort(sort).All(&ingredients)
	if err != nil {
		return
	}
	return
}

func CountAllIngredients() (int, error) {
	db := mongodb.New()
	c := db.C(ingredienteCollection)
	defer db.Close()
	count, err := c.Find(nil).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func DeleteIngredientById(ingredientIDHex string) (bool, error) {
	ingredientID := bson.ObjectIdHex(ingredientIDHex)
	db := mongodb.New()
	c := db.C(ingredienteCollection)
	defer db.Close()
	err := c.Remove(bson.M{"_id": ingredientID})
	if err != nil {
		return false, err
	}
	return true, err
}
