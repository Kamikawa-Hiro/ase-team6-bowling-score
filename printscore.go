package main

import(
	"fmt"
)

func calculateflame(score [21]int, k int) int{
	totalScore := 0
	rollIndex := 0 // 現在処理しているscore配列のインデックス

	// 10フレーム分のスコアを計算
	for frame := 0; frame < k; frame++ {
		
		// ストライクの判定
		if frame < 8{
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
		}else if frame == 8{
			if score[16] == 10 { // 1投目が10ピンならストライク
				totalScore += 10 + score[18] + score[19]

			} else if score[16]+score[17] == 10 { // 1投目と2投目の合計が10ならスペア
				totalScore += 10 + score[18]

			} else { // オープンフレーム
				totalScore += score[16] + score[17]
			}
		}
	}
	if k == 10{
		totalScore += score[18] + score[19] + score[20]
	}


	return totalScore
}

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

func printscore(score [21]int, name string, flame int, sum [10]int) {


	// fmt.Printf("%s の現在のスコア：%d\n", name,sum[flame-1])



	// fmt.Println("スコア：", score)

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
	// fmt.Printf("%s の最大スコア：%d\n", name, maxsum)
	// fmt.Println("最大スコア：", maxscore)
	

	// fmt.Println("---------------------------------------------------------------------------------")
	// fmt.Println("|   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |  10   |")
	fmt.Println("----------------------------------------------------------------------------------------------------")
	fmt.Print("|")
	fmt.Printf("%-10s|", name)
	for i := 0; i < flame; i++ {
		if i < 9 {
			if score[2*i] == 10 {
				fmt.Print(" X |   |")
			}else if score[2*i] + score[2*i+1] == 10 {
				fmt.Printf(" %d |", score[2*i])
				fmt.Printf(" / |")
			}else {
				if score[2*i] == 0{
					fmt.Printf(" G |")
				}else{
					fmt.Printf(" %d |", score[2*i])
				}
				if score[2*i+1] == 0{
					fmt.Printf(" - |")
				}else{
					fmt.Printf(" %d |", score[2*i+1])
				}				
			}
		}else {
			if score[2*i] + score[2*i+1] < 10 {
				if score[2*i] == 0{
					fmt.Printf(" G |")
				}else{
					fmt.Printf(" %d |", score[2*i])
				}
				if score[2*i+1] == 0{
					fmt.Printf(" - |")
				}else{
					fmt.Printf(" %d |", score[2*i+1])
				}
			}else if score[2*i] != 10{
				if score[2*i] == 0{
					fmt.Printf(" G |")
				}else{
					fmt.Printf(" %d |", score[2*i])
				}
				fmt.Printf("/|")
				if score[2*i+2] == 10 {
					fmt.Print("X |")
				}else {
					fmt.Printf("%d |", score[2*i+2])
				}
			}else {
				fmt.Print(" X|")
				if score[2*i+1] == 10 {
					fmt.Print("X|")
					if score[2*i+2] == 10 {
						fmt.Print("X |")
					}else {
						if score[2*i] == 0{
							fmt.Printf("G |")
						}else{
							fmt.Printf("%d |", score[2*i+2])
						}						
					}
				}else {
					if score[2*i] == 0{
						fmt.Printf("G |")
					}else{
						fmt.Printf("%d |", score[2*i+1])
					}
					if score[2*i+1] + score[2*i+2] == 10 {
						fmt.Print("/ |")
					}else {
						if score[2*i+1] == 0{
							fmt.Printf("- |")
						}else{
							fmt.Printf("%d |", score[2*i+2])
						}
					}
				}
			}
		}
	}
	for i:=0; i<10-flame; i++ {
		fmt.Print("       |")
	}
	fmt.Printf("  %3d  |", sum[flame-1])
	fmt.Println("")
	fmt.Print("|")
	fmt.Print("          |")
	for i := 0; i < flame; i++ {
		fmt.Printf("  %3d  |", sum[i])
	}
	for i:=0; i<10-flame; i++ {
		fmt.Print("       |")
	}
	fmt.Printf("  %3d  |", maxsum)
	fmt.Println("")
// 	{|,name,|},
}

func multiPrintscore(players []player, playernum int) {
	fmt.Println("----------------------------------------------------------------------------------------------------")
	fmt.Println("|name      |   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |  10   | score |")
	for i:=0; i<playernum; i++ {
		printscore(players[i].score, players[i].name, players[i].flame, players[i].sum)
	}
	fmt.Println("----------------------------------------------------------------------------------------------------")
}