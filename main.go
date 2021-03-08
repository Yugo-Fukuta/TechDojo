package main

import "github.com/Yugo-Fukuta/TechDojo2/route"

func main() {
	router := route.Init()
	// サーバー起動
	router.Start(":80")
}



