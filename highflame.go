package main


func highflame(players []player, playernum int) [21]int {
	var highscore [21]int
	ittoume := 0
	nitoume := 0
	playersum := 0
	for flame := 0; flame <9; flame++ {
		ittoume = 0
		nitoume = 0
		for i := 0; i < playernum; i++ {
			playersum = players[i].score[flame*2] + players[i].score[flame*2 + 1]

			if players[i].score[flame*2] == 10{
				ittoume = 10
				nitoume = 0
				break;
			}else if playersum == 10{
				if (ittoume + nitoume == 10) && (players[i].score[flame*2] > ittoume){
					ittoume = players[i].score[flame*2]
					nitoume = players[i].score[flame*2 +1]
				}
			}else{
				if playersum > (ittoume + nitoume){
					ittoume = players[i].score[flame*2]
					nitoume = players[i].score[flame*2 +1]
				}else if (playersum == (ittoume + nitoume)) && (players[i].score[flame*2] > ittoume){
					ittoume = players[i].score[flame*2]
					nitoume = players[i].score[flame*2 +1]
				}
			}
		}
		highscore[flame*2] = ittoume
		highscore[flame*2 + 1] = nitoume
	}

	ittoume = 0
	nitoume = 0
	santoume := 0
	
	for i := 0; i < playernum; i++ {
		playersum = players[i].score[18] + players[i].score[19] + players[i].score[20]
		if playersum > (ittoume + nitoume + santoume){
			ittoume = players[i].score[18]
			nitoume = players[i].score[19]
			santoume = players[i].score[20]
		}else if (playersum == (ittoume + nitoume + santoume)){
			if players[i].score[18] > ittoume{
				ittoume = players[i].score[18]
				nitoume = players[i].score[19]
				santoume = players[i].score[20]
			}else if (players[i].score[18] == ittoume) && (players[i].score[19] > nitoume){
				ittoume = players[i].score[18]
				nitoume = players[i].score[19]
				santoume = players[i].score[20]
			}

		}
	}
	highscore[18] = ittoume
	highscore[19] = nitoume	
	highscore[20] = santoume

	return highscore
}