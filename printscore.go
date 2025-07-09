package main

import(
	"fmt"
)
func printscore(score [21]int, name string) {
	sum := 0
	for i:=0; i<21; i++{
		sum = sum + score[i]
	}

	fmt.Printf("%s の現在のスコア：%d\n", name, sum)
	fmt.Println("スコア：", score)

// dif={
// 	{+,|,-,-,-,-,-,-,-,-,-,-,-,-,-,+},
// 	{|,1,"",2,"",3,"",,},
// 	{+,-,-,-,-,-,-,-,-,-,-,-,-,-,+},
// 	{|,name,|},
}