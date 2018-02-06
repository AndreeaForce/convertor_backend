package main

import (
	"github.com/andreeaforce/test2/routes"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

type lala string

func main() {
	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})
	app.Use(crs)
	app.RegisterView(iris.HTML("./views", ".html").Layout("shared/default.html").Reload(true))

	users := app.Party("/ingredients")
	{
		// http://localhost:8080/users/42/profile
		users.HandleMany("GET", "/list /list/{limit int}", hero.Handler(routes.GetListIngredients))
		users.Get("/add", hero.Handler(routes.GetAddIngredients))
		users.Post("/add", hero.Handler(routes.PostAddIngredients)).Name = "addIngredients"
		users.HandleMany("GET POST OPTIONS", "/get /get/{ingredientName string}/{page int} /get/{page int}", hero.Handler(routes.GetIngredientByName))
		users.Post("/count", hero.Handler(routes.GetIngredientsCount))
		users.Get("/delete/{ingredientID string}", hero.Handler(routes.DeleteIngredientByID))
	}

	app.Run(iris.Addr(":8081"))
}
