package main

import (
	"bankCLI/pkg"
	"fmt"
)

// Создание клиента, и его счет
// Пополнить счет клиента
// Посмотреть баланс клиента
// Снять деньги с баланса
// Ошибка недостаточно баланса

func main() {
	for {
		var choice int

		fmt.Println("1. Создание клиента, и его счет")
		fmt.Println("2. Пополнить счет клиента")
		fmt.Println("3. Посмотреть баланс клиента")
		fmt.Println("4. Снять деньги с баланса")
		fmt.Println("5. Выйти")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			pkg.RegisterClient()
		case 2:
			pkg.TopUpClientsAccount()
		case 3:
			pkg.ShowClientsAccount()
		case 4:
			pkg.WithdrawClientsAccount()
		case 5:
			return
		}
	}
}
