package main

import(
	"fmt"
)

func getscore() (int, int) {
	var scoreFirst, scoreSecond int
	for {
		fmt.Println("1投目のスコアを記入してください")
		_, err := fmt.Scan(&scoreFirst)
		if err != nil {
			fmt.Println("エラー：有効な値を入力してください")
			continue
		}
		if 0 <= scoreFirst && scoreFirst <= 10 {
			break
		}
		fmt.Println("エラー：有効な値を入力してください")
	}

	if scoreFirst == 10 {
		return scoreFirst, 0
	}

	for {
		fmt.Println("2投目のスコアを記入してください")
		_, err := fmt.Scan(&scoreSecond)
		if err != nil {
			fmt.Println("エラー：有効な値を入力してください")
			continue
		}
		if 0 <= scoreSecond && scoreSecond <= 10 - scoreFirst {
			break
		}
		fmt.Println("エラー：有効な値を入力してください")
	}

	return scoreFirst, scoreSecond
}

func getscore_10flame() (int, int, int){
	var scoreFirst, scoreSecond, scoreThird int
	for {
		fmt.Println("1投目のスコアを記入してください")
		_, err := fmt.Scan(&scoreFirst)
		if err != nil {
			fmt.Println("エラー：有効な値を入力してください")
			continue
		}
		if 0 <= scoreFirst && scoreFirst <= 10 {
			break
		}
		fmt.Println("エラー：有効な値を入力してください")
	}
	if scoreFirst == 10 {
		for {
			fmt.Println("2投目のスコアを記入してください")
			_, err := fmt.Scan(&scoreSecond)
			if err != nil {
				fmt.Println("エラー：有効な値を入力してください")
				continue
			}
			if 0 <= scoreSecond && scoreSecond <= 10 {
				break
			}
			fmt.Println("エラー：有効な値を入力してください")
		}
		if scoreSecond == 10 {
			for {
				fmt.Println("3投目のスコアを記入してください")
				_, err := fmt.Scan(&scoreThird)
				if err != nil {
					fmt.Println("エラー：有効な値を入力してください")
					continue
				}
				if 0 <= scoreThird && scoreThird <= 10 {
					break
				}
				fmt.Println("エラー：有効な値を入力してください")
			}
			return scoreFirst, scoreSecond, scoreThird
		}else {
			for {
				fmt.Println("3投目のスコアを記入してください")
				_, err := fmt.Scan(&scoreThird)
				if err != nil {
					fmt.Println("エラー：有効な値を入力してください")
					continue
				}
				if 0 <= scoreThird && scoreThird <= 10 - scoreSecond {
					break
				}
				fmt.Println("エラー：有効な値を入力してください")
			}
			return scoreFirst, scoreSecond, scoreThird
		}
	}else {
		for {
			fmt.Println("2投目のスコアを記入してください")
			_, err := fmt.Scan(&scoreSecond)
			if err != nil {
				fmt.Println("エラー：有効な値を入力してください")
				continue
			}
			if 0 <= scoreSecond && scoreSecond <= 10 - scoreFirst {
				break
			}
			fmt.Println("エラー：有効な値を入力してください")
		}

		if scoreFirst + scoreSecond == 10 {
			for {
				fmt.Println("3投目のスコアを記入してください")
				_, err := fmt.Scan(&scoreThird)
				if err != nil {
					fmt.Println("エラー：有効な値を入力してください")
					continue
				}
				if 0 <= scoreThird && scoreThird <= 10 {
					break
				}
				fmt.Println("エラー：有効な値を入力してください")
			}
			return scoreFirst, scoreSecond, scoreThird
		}
	}
	
	return scoreFirst, scoreSecond, 0
}