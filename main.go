package main

import (
	"demo/app-1/account"
	"demo/app-1/encrypt"
	"demo/app-1/files"
	"demo/app-1/output"
	"fmt"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	color.Cyan(`
		__Менеджер паролей__
	`)

	// С помощью библиотеки мы загружаем из файла env
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось гнайти env файл")
	}

	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypt.NewEnrypter())

Loop:
	for {
		num := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		switch num {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		case "4":
			break Loop
		}
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите email аккаунта"})
	vault.DeleteAccountDyUrl(url)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите email аккаунта"})
	accounts := vault.FindAccountDyUrl(url)

	if len(accounts) == 0 {
		color.Red("Аккаунтов нет!")
	}

	for _, account := range accounts {
		account.Output()
	}
}

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
