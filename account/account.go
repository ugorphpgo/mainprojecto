package account

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFHIJKLMNOPQRSTUVWXYZ123456789")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) generatePassword(lenght int) {
	gSlice := make([]rune, lenght)
	for i := range gSlice {
		gSlice[i] = letterRunes[rand.Intn(len(letterRunes))]

	}
	acc.Password = string(gSlice)
}
func (acc *Account) OutputPassword() {
	fmt.Println(acc.Login, acc.Password, acc.Url, acc.CreatedAt, acc.UpdatedAt)
	color.Red(acc.Login)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if len(login) == 0 {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		Url:       urlString,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
