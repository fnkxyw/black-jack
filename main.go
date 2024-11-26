package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Name    string
	Balance int
	Bet     int
}

func (peps *Player) MakeBet() {
	fmt.Println("Сделайте ставку на которую хотите играть: ")
	fmt.Scanf("%d", &peps.Bet)
	if peps.Bet > peps.Balance || peps.Bet < 0 {
		fmt.Println("Проверьте набранную сумму или баланс")
		return
	}

}

func CheckWin(sum int) bool {
	if sum == 21 {
		return true
	}
	return false
}

func MapCheck(m map[int]int) bool {
	for _, val := range m {
		if val < 0 {
			return false
		}
	}
	return true
}

func CheckSum(sum int) bool {
	if sum > 21 {
		return false
	}
	return true
}

func (peps *Player) AddBalance() {
	fmt.Println("Введите сумму на которую хотите пополнить баланс: ")
	dep := 0
	fmt.Scanf("%d", &dep)
	peps.Balance += dep

}

func (peps *Player) Game() {
	choose := 0
	win := true
	var cardsSum int = 0
	var num int = 0
	cards := map[int]int{
		2:  4,
		3:  4,
		4:  4,
		5:  4,
		6:  4,
		7:  4,
		8:  4,
		9:  4,
		10: 16,
		11: 4,
	}
	for i := 0; i < 2; i++ {
		num = rand.Intn(10) + 2
		cardsSum += num
		if cardsSum > 22 {
			cardsSum -= 10
		}
		cards[num]--
	}
	fmt.Println("Сумма ваших начальных карт равна: ", cardsSum)
	for {
		fmt.Println("1.Добавить еще одну карту\n" +
			"2.Оставить ")
		fmt.Scanf("%d", &choose)
		switch choose {
		case 1:
			num = rand.Intn(10) + 2
			if !MapCheck(cards) {
				num = rand.Intn(10) + 2
			}
			cards[num]--
			if num == 11 {
				fmt.Println("Выберите 1 или 11")
				fmt.Scanf("%d", &num)

				switch num {
				case 1:
					num = 1
					break
				case 11:
					num = 11
					break
				default:
					fmt.Println("Выбор не верен")
				}

			}
			fmt.Println("Ваша  карта ", num)
			cardsSum += num
			win = CheckSum(cardsSum)
			if !win {
				fmt.Println("Вы проиграли")
				peps.Balance -= peps.Bet
				return
			}
			win = CheckWin(cardsSum)
			if win {
				fmt.Println("Поздравляем у вас 21")
				peps.Balance += peps.Bet
				return
			}
			fmt.Println("Ваша сумма карт на данный момент:", cardsSum)

		case 2:
			sumPc := rand.Intn(7) + 17
			fmt.Println("Сумма карта компьютера:", sumPc)
			fmt.Println("Ваша сумма карт:", cardsSum)
			if sumPc > 21 {
				fmt.Println("У компьютера перебор, вы выагирайли ")
				peps.Balance += peps.Bet
			} else if cardsSum > sumPc {
				fmt.Println("Поздравляем вы выйграли!!!")
				peps.Balance += peps.Bet
			} else if cardsSum == sumPc {
				fmt.Println("Ничья")
			} else if cardsSum == 21 {
				fmt.Println("Поздравляем вас вы выйграли  у вас 21")
				peps.Balance += peps.Bet
			} else if sumPc > cardsSum {
				fmt.Println("Вы проиграли, все ваши деньги наши")
				peps.Balance -= peps.Bet
			}
			return
		default:
			fmt.Println("Выбор неверен")

		}

	}

}

func (peps *Player) ShowBlanace() {
	fmt.Printf("Ваш баланс: %d$\n", peps.Balance)

}

func (peps *Player) NullBalance() {
	peps.Balance = 0
	fmt.Println("Операция прошла успешно")

}

func main() {
	firstPlayer := Player{}
	x := 0
	for {
		fmt.Println("1.Пополнить баланс.\n" +
			"2.Показать баланс.\n" +
			"3.Начать играть.\n" +
			"4.Вывести деньги.\n" +
			"5.Выход.")
		fmt.Scanf("%d", &x)
		switch x {
		case 1:
			firstPlayer.AddBalance()
		case 2:
			firstPlayer.ShowBlanace()
		case 3:
			if firstPlayer.Balance == 0 {
				fmt.Println("У вас нет денег, сделайте депозит")
				break
			}
			firstPlayer.MakeBet()
			firstPlayer.Game()
		case 4:
			firstPlayer.NullBalance()
		case 5:
			return
		default:
			fmt.Println("Неверный ввод. Повторите попытку")
		}

	}

}
