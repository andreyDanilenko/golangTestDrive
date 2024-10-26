package main

import "fmt"

type stringMap = map[string]string

func main() {
	// m := map[string]string{
	// 	"MyCollection": "https://garwin.ru/",
	// }
	// fmt.Println(m)

	// m["NewCollection"] = "https://ipr.garwin.ru/"
	// m["YANDEX"] = "https://dzen.ru/"
	// fmt.Println(m)

	// delete(m, "YANDEX")
	// fmt.Println(m)

	// for key, value := range m {
	// 	fmt.Println(key, value)
	// }

	// TYPE_ALIAS

	m := stringMap{}

	for {
		getMenuItem()
		option := chooseMenuOption()

		if option == 1 {
			var isEmpty bool = true
			fmt.Println("Список значений")
			for key, value := range m {
				isEmpty = false
				fmt.Println(key + ": " + value)
			}
			if isEmpty {
				fmt.Println(`		  *** СПИСОК ПУСТ ***		  `)
			}
		} else if option == 2 {
			for {
				ans := 'n'
				key, value := addMapValue()
				m[key] = value
				fmt.Print("Значение добавлено! Желаете продолжить добавление? (y / n)\n")
				fmt.Scan(&ans)
				if ans == 'n' {
					break
				}
			}
		} else if option == 3 {
			ans := 'n'
			var value string
			for {
				var isEmpty bool = true
				fmt.Println("Список значений")
				for key, value := range m {
					isEmpty = false
					fmt.Println(key + ": " + value)
				}
				if isEmpty {
					fmt.Println(`		  *** СПИСОК ПУСТ ***		  `)
					break
				}
				fmt.Print("Введите значение которое хотите удалить \n")
				fmt.Scan(&value)
				delete(m, value)
				fmt.Print("Значение удалено! Желаете продолжить удаление? (y / n)\n")
				for key, value := range m {
					fmt.Println(key + ": " + value)
				}
				fmt.Scan(&ans)
				if ans == 'n' {
					break
				}
			}
		} else if option == 4 {
			break
		}
	}

}

func chooseMenuOption() int {
	var option int
	fmt.Scan(&option)
	return option
}

func addMapValue() (string, string) {
	var key string
	var value string

	fmt.Print("Введите название \n")
	fmt.Scan(&key)
	fmt.Print("Введите значение \n")
	fmt.Scan(&value)

	return key, value
}

func answerContinue() string {
	var question string
	fmt.Scan(&question)
	return question
}

func getMenuItem() {
	fmt.Println(`
			__Меню__

		1. Посмотреть закладки
		2. Добавить закладки
		3. Удалить закладки
		4. Выход
	`)
}
