package main

import (
	"training-api/controller"
	"training-api/database"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	ctrl := controller.Controller{database.Create()}

	g := e.Group("/api/v1")
	g.GET("/", ctrl.GetHome)
	g.GET("/posts", ctrl.GetPostsAll)
	g.GET("/posts/:id", ctrl.GetPostById)
	g.POST("/posts", ctrl.PostPost)
	g.PUT("/posts/:id", ctrl.PutPost)
	g.PATCH("/posts/:id", ctrl.PatchPost)
	g.DELETE("/posts/:id", ctrl.DeletePost)

	e.Logger.Fatal(e.Start(":1323"))
}
