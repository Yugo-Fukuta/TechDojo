package route

import (
	"github.com/Yugo-Fukuta/TechDojo2/api"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/test", api.Test)
	e.POST("/user/create", api.HandleUserCreate)
	e.GET("/user/get", api.HandleUserGet)
	e.PUT("/user/update", api.HandleUserUpdate)

	return e
}