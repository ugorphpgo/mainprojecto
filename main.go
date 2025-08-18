package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var menu = map[string]func(db *account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1.Создать аккаунт",
	"2.Найти аккаунт по URL",
	"3.Найти аккаунт по логину",
	"4.удалить аккаунт",
	"5.выйти",
	"Выберите вариант",
}

func main() {
	fmt.Println("Account manager")
	vault := account.NewVault(files.NewJsonDb("data.json")) //Создаем хранилище структур c  зависимостью NewJsonDb
Menu:
	for {

		variant := promptData(menuVariants...)
		if variant == "5" {
			break Menu
		}
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}

}

func createAccount(vault *account.VaultWithDb) { //функция добавления аккаунта,
	login := promptData("Введите логин") //образа
	password := promptData("Введите пароль")
	url := promptData()
	u, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверные формат URL или логина")
		return
	}
	vault.AddAccount(*u)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите url для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		fmt.Println("Аккаунтов не найдено")
	}
	for _, account := range *accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите url для удаления")
	if !vault.DeleteAccountsByUrl(url) {
		color.Green("Аккаунтов с таким url нету")
	} else {
		output.PrintError("Не найдено")
	}

}

func promptData(prompt ...string) string { //динамическое количество аргументов

	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v:", line)
		} else {
			fmt.Println(line)
		}
	}
	var vvod string
	fmt.Scan(&vvod)
	return vvod
}
