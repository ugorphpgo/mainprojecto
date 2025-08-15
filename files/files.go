package account

import (
	"demo/password/files"
	"encoding/json"
	"github.com/fatih/color"
	"strings"
	"time"
)

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"updateAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts: []Account{},
			UpdateAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &vault

}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}

	}
	return accounts
}

func (vault *Vault) DeleteAccountsByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			isDeleted = true
		}

	}
	vault.Accounts = accounts
	vault.save()

	return isDeleted
}
func (Vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(Vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) save() {
	vault.UpdateAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать данные  в json")
	}
	files.WriteFile(data, "data.json")
}
