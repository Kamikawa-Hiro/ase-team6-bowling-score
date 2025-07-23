package main

import(
	"fmt"
)

func getplayernum() int {
	var playernum int
	for {
		fmt.Print("人数を入力してください： ")
		_, err := fmt.Scanln(&playernum)
		if err != nil {
			fmt.Println("エラー：有効な値を入力してください")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		break
	}
	return playernum
}

func getplayername() string {
	var name string
	for {
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Println("エラー：有効な名前を入力してください")
			continue
		}
		break
	}
	return name
}