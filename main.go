package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

type UserController struct{ engine *echo.Group }

func NewUserController(engine *echo.Group) *UserController {
	controller := &UserController{engine}
	engine.GET("/user/:id", controller.getUser)
	engine.POST("/user/:id", controller.getUser)
	engine.PUT("/user/:id", controller.getUser)
	engine.DELETE("/user/:id", controller.getUser)

	return controller
}

func (rc *UserController) getUser(c echo.Context) error {
	return nil
}

func main() {
	e := echo.New()

	userGroup := e.Group("/admin")
	userGroup.GET("/", handler)
	userGroup.POST("/", handler)
	userGroup.PUT("/", handler)
	userGroup.DELETE("/", handler)

	e.Logger.Fatal(e.Start(":8888"))
}
