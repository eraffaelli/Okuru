package routes

import (
	"github.com/eraffaelli/Okuru/controllers"
	"github.com/labstack/echo"
)

func Index(e *echo.Echo) {
	e.GET("/", controllers.Index)
	e.POST("/", controllers.AddIndex)
	e.GET("/:password_key", controllers.ReadIndex)
	e.GET("/remove/:password_key", controllers.DeleteIndex)
}