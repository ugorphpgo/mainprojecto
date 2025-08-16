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

// пример внедрения зависимости Jsond.db
type VaultWithDb struct { //Создаем структуру которая учитывает файл в котором будет сохраняться vault
	Vault
	db files.Jsondb //Внедренная зависимость от jsonDb
}

func NewVault(db *files.Jsondb) *VaultWithDb { //Внедрение зависимости files.JsonDb в vault - передаем указатель на бд
	file, err := db.Read()
	if err != nil { //действия не удалось прочитать инфу из файла - его нет//
		return &VaultWithDb{ //возвращаем и vault и его поля и добавляем к этому поле db
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db: *db,
		}
	}
	var vault VaultWithDb
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &VaultWithDb{
		Vault: Vault{
			Accounts: []Account{},
			UpdateAt: time.Now(),
		},
		db: *db,
	}

}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}

	}
	return accounts
}

func (vault *VaultWithDb) DeleteAccountsByUrl(url string) bool {
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
func (Vault *Vault) ToBytes() ([]byte, error) { // Запись значение из хранилища vault в json файл
	file, err := json.Marshal(Vault) //Важно чтобы этот метод работал только с начальным хранилищем Vault
	if err != nil {                  //без значения jdondb - потому что идёт преобразование только зависимости без самого файла
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDb) save() { //сохранение файла
	vault.UpdateAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать данные  в json")
	}
	vault.db.Write(data)

}
