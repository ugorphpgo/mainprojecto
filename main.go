package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	fmt.Println("Account manager")
Menu:
	for {
		variant := getMenu()
		if variant == 4 {
			break Menu
		}
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}

}
func getMenu() int {
	var choice int
	fmt.Println("Меню - выберите что хотите сделать")
	fmt.Println("1.Create account")
	fmt.Println("2.Find accout")
	fmt.Println("3.Delete account")
	fmt.Println("4.Exit from app")
	fmt.Scanln(&choice)
	return choice

}

func createAccount() {
	fmt.Println("Введите логин:")
	login := promptData()
	fmt.Println("Введите пароль:")
	password := promptData()
	fmt.Println("Введите url:")
	url := promptData()
	u, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверные заполненые данные")
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*u)
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Не получилось преобразовать в JSON")
		return
	}
	files.WriteFile(data, "data.json")
}

func findAccount() (f []byte, err error) {
	return f, err
}

func deleteAccount() {

}

func promptData() string {
	fmt.Print(" ")
	var vvod string
	fmt.Scanln(&vvod)
	return vvod
}
