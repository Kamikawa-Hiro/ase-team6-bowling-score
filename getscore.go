package main

import(
	"fmt"
)

func getscore() (int, int) {
	var scoreFirst, scoreSecond int
	fmt.Println("1投目のスコアを記入してください")
	fmt.Scan(&scoreFirst)
	fmt.Println("2投目のスコアを記入してください")
	fmt.Scan(&scoreSecond)

	return scoreFirst, scoreSecond
}

func getscore_10flame() (int, int, int){
	var scoreFirst, scoreSecond, scoreThird int
	fmt.Println("1投目のスコアを記入してください")
	fmt.Scan(&scoreFirst)
	fmt.Println("2投目のスコアを記入してください")
	fmt.Scan(&scoreSecond)
	if scoreFirst + scoreSecond >= 10 {
		fmt.Println("3投目のスコアを記入してください")
		fmt.Scan(&scoreThird)
		return scoreFirst, scoreSecond, scoreThird
	}
	return scoreFirst, scoreSecond, 0
}