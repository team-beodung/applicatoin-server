package user

import "github.com/labstack/echo"

type UserDTO struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserController struct{ 
	echo *echo.Group 

	userRepository *UserRepository
}


func NewUserController(echo *echo.Group, r *UserRepository) *UserController {
	controller := &UserController{echo, r}

	echo.GET("/:id", controller.getUser)
	echo.POST("/:id", controller.getUser)
	echo.PUT("/:id", controller.getUser)
	echo.DELETE("/:id", controller.getUser)
	return controller
}

func (c *UserController) getUser(ctx echo.Context) error {
	userDTO := c.userRepository.GetUser(ctx.Param("id"))

	return c.JSON(200, UserDTO{ID: 1, Name: "John", Email: "hwc9169@gmail.com"})
}

