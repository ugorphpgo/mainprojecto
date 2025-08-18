package account

import (
	"demo/password/output"
	"encoding/json"
	"strings"
	"time"
)

type Db interface { //Создаем интерфейс для работы с файлами - контракт взаимодействия с файлами
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updateAt"`
}

// пример внедрения зависимости Jsond.db
type VaultWithDb struct { //Создаем структуру которая учитывает файл в котором будет сохраняться vault
	Vault
	db Db //Внедренная зависимость от jsonDb
}

func NewVault(db Db) *VaultWithDb { //Внедрение зависимости files.JsonDb в vault - передаем указатель на бд
	file, err := db.Read()
	if err != nil { //действия не удалось прочитать инфу из файла - его нет//
		return &VaultWithDb{ //возвращаем и vault и его поля и добавляем к этому поле db
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault) //читаем из jsona инфрмацию
	if err != nil {
		output.PrintError("Не удалось получить информацию из json.db")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}

}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Vault.Accounts {
		isMatched := checker(account, str)
		if isMatched == true {
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
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		output.PrintError("Не удалось преобразовать данные  в json")
	}
	vault.db.Write(data)

}
