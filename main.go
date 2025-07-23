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

				players[j].sum[i] = calculate(players[j].score)
				// printscore(players[j].score, players[j].name, players[j].flame, players[j].sum)
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
				
				players[j].sum[i] = calculate(players[j].score)
				// printscore(players[j].score, players[j].name, players[j].flame, players[j].sum)	
			}
			multiPrintscore(players, playernum)
		}
	}

	review(players, playernum)
	strike_and_spare_rate(players, playernum)
}