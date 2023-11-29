package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type CoinAmountSelection int

const (
	Ten CoinAmountSelection = iota
	Thirty
	Fifty
	QuitGame
)

// ポイントやコインの初期値
const (
	NORMAL_COIN  = 50
	NORMAL_POINT = 50
	SEVEN_COIN   = 100
	SEVEN_POINT  = 100
	SEVEN_NUMBER = "7"
)

var Coins = []int{10, 30, 50}

func selectedCoinAmount(coins, point int) int {
	var input string
	for {
		fmt.Println("---------")
		fmt.Println("何コイン入れますか？")
		fmt.Printf("残りコイン: %s枚\nポイント: %spt\n", strconv.Itoa(coins), strconv.Itoa(point))
		fmt.Println("1(10コイン) 2(30コイン) 3(50コイン) 4(やめとく)")
		fmt.Println("---------")

		fmt.Scan(&input)
		amount, err := strconv.Atoi(input)

		if err == nil {
			amount--
			if amount >= int(Ten) && amount <= int(QuitGame) {
				return amount
			}
		}
		fmt.Println("正しい値を入力してください。")
	}
}

func getEnterKey() {
	for {
		fmt.Print("エンターのみを押してください >")

		var input string
		fmt.Scanln(&input)

		if input == "" {
			return
		}
	}
}

// TO DO
func calcCoinAndPoint(numbers [][]string) (int, int) {
	point, coin := 0, 0

	// 各行の判定
	for i := 0; i < 3; i++ {
		if numbers[i][0] == numbers[i][1] && numbers[i][1] == numbers[i][2] && numbers[i][0] != "" {
			fmt.Printf("ヨコに%sが揃いました!\n", numbers[i][0])
			if numbers[i][0] == SEVEN_NUMBER {
				coin += SEVEN_COIN
				point += SEVEN_POINT
			} else {
				coin += NORMAL_COIN
				point += NORMAL_POINT
			}
		}
	}

	// 各列の判定
	for i := 0; i < 3; i++ {
		if numbers[0][i] == numbers[1][i] && numbers[1][i] == numbers[2][i] && numbers[0][i] != "" {
			fmt.Printf("タテに%sが揃いました!\n", numbers[0][i])
			if numbers[0][i] == SEVEN_NUMBER {
				coin += SEVEN_COIN
				point += SEVEN_POINT
			} else {
				coin += NORMAL_COIN
				point += NORMAL_POINT
			}
		}
	}

	// 対角線の判定
	if (numbers[0][0] == numbers[1][1] && numbers[1][1] == numbers[2][2] && numbers[0][0] != "") ||
		(numbers[0][2] == numbers[1][1] && numbers[1][1] == numbers[2][0] && numbers[0][2] != "") {
		fmt.Printf("ナナメに%sが揃いました!\n", numbers[1][1])
		if numbers[1][1] == SEVEN_NUMBER {
			coin += SEVEN_COIN
			point += SEVEN_POINT
		} else {
			coin += NORMAL_COIN
			point += NORMAL_POINT
		}
	}
	return coin, point
}

func initSlotPanel(n int) [][]string {
	slice := make([][]string, n)
	for i := range slice {
		slice[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			slice[i][j] = "*"
		}
	}
	return slice
}

func main() {
	totalcoins := 100
	totalpoint := 0

	for {
		selected := selectedCoinAmount(totalcoins, totalpoint)

		if selected == int(QuitGame) {
			return
		} else if totalcoins < Coins[selected] {
			fmt.Println("コインの枚数が不足しています。")
			continue
		}
		totalcoins -= Coins[selected]
		numbers := initSlotPanel(3)

		for i := 0; i < 3; i++ {
			getEnterKey()
			for j := 0; j < 3; j++ {
				numbers[i][j] = strconv.Itoa(rand.Intn(8) + 2)
			}
			fmt.Println("---------")
			fmt.Printf("|%s|%s|%s|\n", numbers[0][0], numbers[0][1], numbers[0][2])
			fmt.Printf("|%s|%s|%s|\n", numbers[1][0], numbers[1][1], numbers[1][2])
			fmt.Printf("|%s|%s|%s|\n", numbers[2][0], numbers[2][1], numbers[2][2])
			fmt.Println("---------")
		}

		coin, point := calcCoinAndPoint(numbers)
		if selected == int(Thirty) {
			totalcoins += coin * 3
			totalpoint += point * 3
		} else if selected == int(Fifty) {
			totalcoins += coin * 5
			totalpoint += point * 5
		} else {
			totalcoins += coin
			totalpoint += point
		}

		if totalcoins < Coins[int(Ten)] {
			fmt.Println("コインがなくなりました")
			fmt.Println("//GAME OVER//")
			return
		}
	}
}
