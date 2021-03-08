package api

import (
	"github.com/Yugo-Fukuta/TechDojo2/model"
	"github.com/google/uuid"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"net/http"
)

func HandleUserCreate(c echo.Context) error {
	// インスタンスの作成
	user := new(model.User)

	// リクエストの読み込み
	if err := c.Bind(user); err != nil {
		return err
	}

	// IDの生成
	userID, _ := uuid.NewRandom()
	user.ID = userID.ClockSequence()

	// トークンの生成
	authToken, _ := uuid.NewRandom()
	user.Token = authToken.String()

	pp.Println(user)

	// データベースへ保存
	model.CreateUser(user)

	// レスポンス
	jsonMap := map[string] string {
		"token" : user.Token,
	}
	return c.JSON(http.StatusOK, jsonMap)
}



