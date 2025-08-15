package account

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"updateAt"`
}

func NewVault() *Vault {
	return &Vault{
		Accounts: []Account{},
		UpdateAt: time.Now(),
	}
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
