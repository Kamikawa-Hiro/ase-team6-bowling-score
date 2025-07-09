package main

import(
	"fmt"
)

func getscore() (int, int) {
	var scoreFirst, scoreSecond int
	fmt.Println("1投目のスコアを記入してください")
	fmt.Scan(&scoreFirst)
	fmt.Println("2投目のスコアを記入してください")
	fmt.Scan(&scoreFirst)

	return scoreFirst, scoreSecond
}

func getscore_10flame(score [21]int){
	var scoreFirst, scoreSecond, scoreThird int
	fmt.Println("1投目のスコアを記入してください")
	fmt.Scan(&scoreFirst)
	fmt.Println("2投目のスコアを記入してください")
	fmt.Scan(&scoreSecond)
	score[18] = scoreFirst
	score[19] = scoreSecond
	if scoreFirst + scoreSecond >= 10 {
		fmt.Println("3投目のスコアを記入してください")
		fmt.Scan(&scoreThird)
		score[20] = scoreThird
	}
}