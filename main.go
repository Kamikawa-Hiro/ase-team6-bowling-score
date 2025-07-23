package main

import(
	"fmt"
)

type player struct{
	name string
	score [21]int
	flame int
	strike int
	spare int
	sum [10]int
}

func main(){
	playernum := getplayernum()
	players := make([]player, playernum)
	for i:=0; i< playernum; i++{
		fmt.Printf("%d 番目のプレイヤー名を入力してください\n", i+1)
		players[i].name = getplayername()
	}

	for i:=0; i<10; i++{
		fmt.Printf("%d フレーム目\n",i+1)
		if i == 9{
			for j:=0; j<playernum; j++{
				players[j].score[2*i], players[j].score[2*i+1], players[j].score[2*i+2] = getscore_10flame()		//ペアプロで
				players[j].flame = i+1
		
				if players[j].score[18] == 10{
					players[j].strike += 1
					if players[j].score[19] == 10{
						players[j].strike += 1
						if players[j].score[20] == 10{
							players[j].strike += 1
						}
					}else if players[j].score[19] + players[j].score[20] == 10{
						players[j].spare += 1
					}
				}else if players[j].score[18] + players[j].score[19] == 10{
					players[j].spare += 1
					if players[j].score[20] == 10{
						players[j].strike += 1
					}
				}

				for k:=0; k < 10; k++{
					players[j].sum[k] = calculateflame(players[j].score, k+1)
				}
					//printscore(players[j].score, players[j].name, players[j].flame, players[j].sum)				


			}
			multiPrintscore(players, playernum)
		}else if i < 9{
			for j:=0; j<playernum; j++{
				players[j].score[2*i], players[j].score[2*i+1] = getscore()
				players[j].flame = i+1

				if players[j].score[2*i] == 10{
					players[j].strike += 1
				}else if players[j].score[2*i] + players[j].score[2*i+1] == 10{
					players[j].spare += 1
				}

				for k:=0; k < 10; k++{
					players[j].sum[k] = calculateflame(players[j].score, k+1)
				}
				//printscore(players[j].score, players[j].name, players[j].flame, players[j].sum)	

			}
			multiPrintscore(players, playernum)
		}
	}

	review(players, playernum)
	strike_and_spare_rate(players, playernum)


	var high  [21]int
	var highsum  [10]int
	high = highflame(players, playernum)
	for k:=0; k < 10; k++{
		highsum[k] = calculateflame(high, k+1)
	}
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Println("|       |   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |  10   |")
	printscore(high, "highflame", 10, highsum)

}