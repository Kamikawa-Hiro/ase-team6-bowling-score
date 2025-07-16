package main

import(
	"fmt"
)

func review(players []player, playernum int){
	if playernum >= 2{
		max = 0
		max_players := 0
		for i:=0;i<playernum;i++{
			if max < players[i].sum{
				max = players[i].sum
				max_players = i
			}
		}
		fmt.Printf("%s が %d ピンで優勝!\n", players[max_players].name, max)
	}
}

func strike_and_spare_rate(players []player, playernum int){
	strike_rate := 0
	spare_rate := 0
	for i:=0; i<playernum; i++{
		strike_chance := 10
		for j:=18;j<20;j++{
			if players[i].score[j] == 10{
				strike_chance += 1
			}
		}

		strike_rate = (players[i].strike / strike_chance) * 100


		spare_chance := 0
		for j:=0;j<10;j++{
			if players[i].score[j*2] != 10{
				spare_chance += 1
			}
		}
		if players[i].score[18] == 10 || players[i].score[19] != 10{
			spare_chance += 1
		}

		spare_rate = (players[i].spare / spare_chance) * 100

		fmt.Printf("%s のストライク率：%d % スペア率：%d %", players[i].name, strike_rate, spare_rate)
	}
}