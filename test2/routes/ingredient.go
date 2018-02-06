package routes

import (
	"fmt"

	"github.com/andreeaforce/test2/models"
	"github.com/andreeaforce/test2/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

var (
	addIngredientsPath = "/ingredients/add"
)

type ListIngredientsResponse struct {
	Data  []models.Ingredient `json:"data"`
	Error string              `json:"error"`
}

var (
	successListIngredients = ListIngredientsResponse{Error: ""}
	errorListIngredients   = ListIngredientsResponse{Error: "There was an error retriving the ingredients!"}
)

func GetListIngredients(ctx iris.Context) ListIngredientsResponse {
	var (
		limit       int
		err         error
		ingredients []models.Ingredient
	)
	limit, err = ctx.Params().GetInt("limit")
	if err != nil {
		fmt.Println(err)
		return errorListIngredients
	}
	ingredients, err = repositories.GetAllIngredients(limit)
	if err != nil {
		return errorListIngredients
	}
	successListIngredients.Data = ingredients
	return successListIngredients
}

func GetAddIngredients(ctx iris.Context) {
	ctx.View(addIngredientsPath + ".html")
}

func PostAddIngredients(ctx iris.Context) hero.Response {
	ingredient := models.Ingredient{}
	err := ctx.ReadForm(&ingredient)
	if err != nil {
		fmt.Println(err)
	}
	repositories.InsertIngredient(ingredient)
	return hero.Response{
		Err:  err,
		Path: addIngredientsPath,
	}
}

type IngredientByNameResponse struct {
	Data  []models.Ingredient `json:"data"`
	Count int                 `json:"count"`
}

func GetIngredientByName(ctx iris.Context) IngredientByNameResponse {
	var (
		ingredients    []models.Ingredient
		limit          int
		page           int
		err            error
		ingredientName string
		count          int
	)
	ingredientName = ctx.Params().Get("ingredientName")
	if ingredientName == "" {
		ingredientName = ctx.PostValue("ingredientName")
	}
	sort := ctx.PostValue("sort")
	if sort == "" {
		sort = "calorii"
	}
	page, err = ctx.Params().GetInt("page")
	if err != nil {
		return IngredientByNameResponse{[]models.Ingredient{}, 0}
	}
	limit, err = ctx.PostValueInt("limit")
	if err != nil {
		return IngredientByNameResponse{[]models.Ingredient{}, 0}
	}
	ingredients, count, err = repositories.GetIngredientByName(ingredientName, page, limit, sort)
	if err != nil {
		return IngredientByNameResponse{[]models.Ingredient{}, 0}
		fmt.Println(err)
	}
	return IngredientByNameResponse{ingredients, count}
}

func GetIngredientsCount(ctx iris.Context) iris.Map {
	count, err := repositories.CountAllIngredients()
	if err != nil {
		fmt.Println(err)
	}
	return iris.Map{"count": count}
}

type deleteIngredientResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

var (
	successDeleteIngredient = deleteIngredientResponse{true, "Deleted successfully"}
	errorDeleteIngredient   = deleteIngredientResponse{false, "There was an error deleteing"}
)

// deleteIngredientByID deletes an ingredient from the database by ID
func DeleteIngredientByID(ctx iris.Context) deleteIngredientResponse {
	ingredientID := ctx.Params().Get("ingredientID")
	ok, err := repositories.DeleteIngredientById(ingredientID)
	if !ok {
		fmt.Println(err)
		return errorDeleteIngredient
	}
	return successDeleteIngredient
}
