package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("Account manager")
	vault := account.NewVault()
Menu:
	for {
		variant := getMenu()
		if variant == 4 {
			break Menu
		}
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
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

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин:")
	password := promptData("Введите пароль:")
	url := promptData("Введите url:")
	u, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверные заполненые данные")
		return
	}
	vault.AddAccount(*u)
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Не получилось преобразовать в JSON")
		return
	}
	files.WriteFile(data, "data.json")

}

func findAccount(vault *account.Vault) {
	url := promptData("Введите url для поиска")
	accounts := vault.FindAccountsByUrl(url)
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите url для удаления")
	if !vault.DeleteAccountsByUrl(url) {
		color.Green("Аккаунтов с таким url нету")
	} else {
		color.Red("Аккаунты с таким url удалены")
	}

}

func promptData(x string) string {
	fmt.Println(x + " ")
	var vvod string
	fmt.Scanln(&vvod)
	return vvod
}
