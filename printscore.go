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
func printscore(score [21]int, name string, flame int, sum [10]int) {
	fmt.Printf("%s の現在のスコア：%d\n", name)
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
	

	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Println("|   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |  10   |")
	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Print("|")
	for i := 0; i < flame; i++ {
		if i < 9 {
			if score[2*i] == 10 {
				fmt.Print(" X |   |")
			}else {
				fmt.Printf(" %d |", score[2*i])
				fmt.Printf(" %d |", score[2*i+1])
			}
		}else {
			if score[2*i] + score[2*i+1] < 10 {
				fmt.Printf(" %d |", score[2*i])
				fmt.Printf(" %d |", score[2*i+1])
			}else {
				fmt.Printf(" %d|", score[2*i])
				fmt.Printf("%d|", score[2*i+1])
				fmt.Printf("%d |", score[2*i+2])
			}
		}
	}
	fmt.Println("")
	fmt.Print("|")
	for i := 0; i < flame; i++ {
		fmt.Printf("  %3d  |", sum[i])
	}
	fmt.Println("")
// 	{|,name,|},
}