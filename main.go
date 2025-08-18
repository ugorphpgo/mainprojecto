package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("Account manager")
	vault := account.NewVault(files.NewJsonDb("data.json")) //Создаем хранилище структур c  зависимостью NewJsonDb

Menu:
	for {
		variant := promptData([]string{
			"1.Create account",
			"2.Find accout",
			"3.Delete account",
			"4.Exit from app",
			"Меню - выберите что хотите сделать",
		})
		if variant == "4" {
			break Menu
		}
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}

}

func createAccount(vault *account.VaultWithDb) { //функция добавления аккаунта,
	login := promptData([]string{"Введите логин"}) //образа
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите url"})
	u, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверные формат URL или логина")
		return
	}
	vault.AddAccount(*u)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите url для поиска"})
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		fmt.Println("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите url для удаления"})
	if !vault.DeleteAccountsByUrl(url) {
		color.Green("Аккаунтов с таким url нету")
	} else {
		output.PrintError("Не найдено")
	}

}

func promptData[T any](prompt []T) string { //generic используищийся в запросе значения
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v :", line)
		} else {
			fmt.Println(line)
		}
	}
	var vvod string
	fmt.Scan(&vvod)
	return vvod
}
