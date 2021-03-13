package api

import (
	"github.com/Yugo-Fukuta/TechDojo2/model"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func HandleGachaDraw(c echo.Context) error {
	// 認証トークンの読み込み
	//token := c.Request().Header.Get("X-Token")

	// 試行回数の読み込み
	t := new(model.Time)
	if err := c.Bind(t); err != nil {
		return err
	}
	times, _ := strconv.Atoi(t.Times)

	// ガチャ実行(キャラクターテーブルからランダムにget)
	var character *model.Character
	var Results *model.Characters
	for i:=0; i<times; i++ {
		character = model.GetCharacter(&model.Character{ID: 1})
		Results = append(Results, character)
	}
	pp.Println(Results)

	// ユーザーキャラクターテーブルに保存


	// レスポンス
	//jsonMap := map[string] string {
	//	"name" : user.Name ,
	//}
	return c.JSON(http.StatusOK, character)
}