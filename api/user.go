package api

import (
	"github.com/Yugo-Fukuta/TechDojo2/model"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"net/http"
)

func HandleUserGet(c echo.Context) error {
	// 認証トークンの読み込み
	token := c.Request().Header.Get("X-Token")

	// ユーザー情報の取得
	user := model.GetUser(&model.User{Token: token})
	pp.Println(user)

	// レスポンス
	jsonMap := map[string] string {
		"name" : user.Name ,
	}
	return c.JSON(http.StatusOK, jsonMap)
}

func HandleUserUpdate(c echo.Context) error {
	// 認証トークンの読み込み
	token := c.Request().Header.Get("X-Token")

	// ユーザー情報の取得
	user := model.GetUser(&model.User{Token: token})
	pp.Println(user)

	// リクエストボディの読み込み
	nuser := new(model.User)
	if err := c.Bind(nuser); err != nil {
		return err
	}

	// ユーザー情報の更新
	user.Name = nuser.Name

	// 更新情報の保存
	model.UpdateUSer(user)
	pp.Println(user)

	// レスポンス
	jsonMap := map[string] string {
		"name" : user.Name ,
	}
	return c.JSON(http.StatusOK, jsonMap)
}