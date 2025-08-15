package account

import (
	"demo/password/files"
	"encoding/json"
	"github.com/fatih/color"
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
	vault.UpdateAt = time.Now()
}

func (Vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(Vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
