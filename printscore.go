package main

import(
	"fmt"
)

func calculate(score [21]int) int{
	totalScore := 0
	rollIndex := 0 // 現在処理しているscore配列のインデックス

	// 10フレーム分のスコアを計算
	for frame := 0; frame < 9; frame++ {
		
		// ストライクの判定
		if score[rollIndex] == 10 { // 1投目が10ピンならストライク
			if score[rollIndex+2] == 10{
				totalScore += 20 + score[rollIndex+4]
			}else{
				totalScore += 10 + score[rollIndex+2] + score[rollIndex+3]
			} 
			rollIndex += 2 

		} else if score[rollIndex]+score[rollIndex+1] == 10 { // 1投目と2投目の合計が10ならスペア
			totalScore += 10 + score[rollIndex+2]
			rollIndex += 2 

		} else { // オープンフレーム
			totalScore += score[rollIndex] + score[rollIndex+1]
			rollIndex += 2 
		}
	}

	totalScore += score[18] + score[19] + score[20]

	return totalScore
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