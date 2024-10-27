package account

import (
	"demo/app-1/encrypt"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteWriter
	ByteReader
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypt.Encrypt
}

// dependency injection
// db передается с типом files.JsonDb
func NewVault(db Db, enc encrypt.Encrypt) *VaultWithDb {
	// db := files.NewJsonDb("data.json")
	file, err := db.Read()
	if err != nil || len(file) == 0 {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	// data, _ := enc.Decrypt(file)

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red("Не удалось разобрать файл data.vault")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

// / код дублируется
// func (vault *VaultWithDb) FindAccountByUrl(url string) []Account {
// 	var accounts []Account

// 	for _, account := range vault.Accounts {
// 		isMatched := strings.Contains(account.Url, url)
// 		if isMatched {
// 			accounts = append(accounts, account)
// 		}
// 	}

// 	return accounts
// }

// func (vault *VaultWithDb) FindAccountByLogin(login string) []Account {
// 	var accounts []Account

// 	for _, account := range vault.Accounts {
// 		isMatched := strings.Contains(account.Login, login)
// 		if isMatched {
// 			accounts = append(accounts, account)
// 		}
// 	}

// 	return accounts
// }

// /
func (vault *VaultWithDb) FindAccount(str string, checker func(Account, string) bool) []Account {
	var accounts []Account

	for _, account := range vault.Accounts {
		// isMatched := strings.Contains(account.Url, url)
		isMatched := checker(account, str)

		if isMatched {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

func (vault *VaultWithDb) DeleteAccountDyUrl(url string) {
	var accounts []Account

	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
		}
	}

	vault.Accounts = accounts
	vault.Save()
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.Save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (vault *VaultWithDb) Save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()

	encData := vault.enc.Encrypt(data)
	if err != nil {
		color.Red("Не удалось перобразовать!")
	}
	vault.db.Write(encData)
}
