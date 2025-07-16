package main

import(
	"fmt"
)

func calculate(score [21]int) int{
	sum := 0
	for i:=0; i<21; i++{
		sum = sum + score[i]
	}

	return sum 
}
func printscore(score [21]int, name string, flame int) {
	sum := calculate(score)
	fmt.Printf("%s の現在のスコア：%d\n", name, sum)
	fmt.Println("スコア：", score)

	maxscore := score
	for i:=flame; i< 10; i++{
		maxscore[(i)*2] = 10
	}
	if flame != 10{
		maxscore[18] = 10
		maxscore[19] = 10
		maxscore[20] = 10
	}

	maxsum := calculate(maxscore)
	fmt.Printf("%s の最大スコア：%d\n", name, maxsum)
	fmt.Println("最大スコア：", maxscore)
	

// dif={
// 	{+,|,-,-,-,-,-,-,-,-,-,-,-,-,-,+},
// 	{|,1,"",2,"",3,"",,},
// 	{+,-,-,-,-,-,-,-,-,-,-,-,-,-,+},
// 	{|,name,|},
}