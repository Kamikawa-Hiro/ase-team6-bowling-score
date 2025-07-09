package main

import(
	"fmt"
)

func getplayernum() int {
	var playernum int
	fmt.Print("人数を入力してください： ")
	fmt.Scanln(&playernum)
	return playernum
}

func getplayername() string {
	var name string
	fmt.Scanln(&name)
	return name
}