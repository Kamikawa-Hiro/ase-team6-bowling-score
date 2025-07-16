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
				players[j].flame = i
				printscore(players[j].score, players[j].name)				
			}
		}else if i < 9{
			for j:=0; j<playernum; j++{
				players[j].score[2*i], players[j].score[2*i+1] = getscore()
				players[j].flame = i
				printscore(players[j].score, players[j].name)
			}
		}
	}
}