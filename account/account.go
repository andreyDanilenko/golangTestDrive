package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("1234567890zaqxswcdevfrbgtnhymjukliop")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// type AccountWithTimeStamp struct {
// 	createdAt time.Time
// 	updatedAt time.Time
// 	Account
// }

// метод структуры
func (acc *Account) OutputPassword() {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) Output() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for index := range res {
		res[index] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

// укзатель позволяет не копировать дополнительно данные
// 1. Если нет логина бросить ошибку
// 2. Eсли нет пароля
func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
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
		Url:       urlString,
		Password:  password,
		// Использование основного конструктора для создания аккаунта
		// Account: Account{
		// 	login:    login,
		// 	url:      urlString,
		// 	password: password,
		// },
	}
	// Вывод мета данных
	// field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	// fmt.Println(string(field.Tag))
	if password == "" {
		newAcc.generatePassword(12)
		// newAcc.account.generatePassword(12)
	}

	return newAcc, nil
}

// func newAccount(login, password, urlString string) (*account, error) {
// 	if login == "" {
// 		return nil, errors.New("INVALID_LOGIN")
// 	}
// 	_, err := url.ParseRequestURI(urlString)

// 	if err != nil {
// 		return nil, errors.New("INVALID_URL")
// 	}

// 	newAcc := &account{
// 		login:    login,
// 		url:      urlString,
// 		password: password,
// 	}

// 	if password == "" {
// 		newAcc.generatePassword(12)
// 	}

// 	return newAcc, nil
// }
