package route

import (
	"github.com/Yugo-Fukuta/TechDojo2/api"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/user/create", api.HandleUserCreate)
	e.GET("/user/get", api.HandleUserGet)
	e.PUT("/user/update", api.HandleUserUpdate)

	e.POST("/gacha/draw", api.HandleGachaDraw)

	return e
}