package main

import (
	"demo/app-1/account"
	"demo/app-1/encrypt"
	"demo/app-1/files"
	"demo/app-1/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

// var menu = map[string]func(string)(int){}
type menuFunc func(*account.VaultWithDb, string)

// Обновляем карту меню, чтобы функции были приведены к одному типу menuFunc
var menu = map[string]menuFunc{
	"1": func(vault *account.VaultWithDb, _ string) {
		createAccount(vault)
	},
	"2": func(vault *account.VaultWithDb, variant string) {
		findAccount(vault, variant)
	},
	"3": func(vault *account.VaultWithDb, variant string) {
		findAccount(vault, variant)
	},
	"4": func(vault *account.VaultWithDb, _ string) {
		deleteAccount(vault)
	},
}

func main() {
	color.Cyan(`
		__Менеджер паролей__
	`)

	// С помощью библиотеки мы загружаем из файла env
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось гнайти env файл")
	}

	vault := account.NewVault(files.NewJsonDb("data.json"), *encrypt.NewEnrypter())

Loop:
	for {
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт по email",
			"3. Найти аккаунт по login",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите вариант",
		})
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Loop
		}

		menuFunc(vault, variant)

		// switch num {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// case "4":
		// 	break Loop
		// }
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите email аккаунта"})
	vault.DeleteAccountDyUrl(url)
}

type SearchVariant struct {
	Text       string
	SearchFunc func(acc account.Account, str string) bool
}

func findAccount(vault *account.VaultWithDb, variant string) {
	variantsText := map[string]SearchVariant{
		"2": {
			Text: "Введите email аккаунта",
			SearchFunc: func(acc account.Account, str string) bool {
				return strings.Contains(acc.Url, str)
			},
		},
		"3": {
			Text: "Введите login аккаунта",
			SearchFunc: func(acc account.Account, str string) bool {
				return strings.Contains(acc.Login, str)
			},
		},
	}
	str := promptData([]string{variantsText[variant].Text})
	// accounts := vault.FindAccount(url, checkUrl)
	// анонимная функция
	accounts := vault.FindAccount(str, variantsText[variant].SearchFunc)

	// accounts := vault.FindAccount(str, func(acc account.Account, str string) bool {
	// 	return strings.Contains(acc.Url, str)
	// })

	getInfo(&accounts)
}

func getInfo(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунтов нет!")
	}

	for _, account := range *accounts {
		account.Output()
	}
}

// func checkUrl(acc account.Account, str string) bool {
// 	return strings.Contains(acc.Url, str)
// }

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL: 'https://example.com'"})

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Println("Неверный формат URL или Login")
		return
	}

	vault.AddAccount(*myAccount)
}

// принимает слайс люого типа
// Выводит строкой каждый элемент
func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	// fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
