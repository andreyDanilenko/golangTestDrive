package second

import "fmt"

// В цикле спрашиваем ввод транзакйций: -10, 10, 40.5
// Добавлять каждую в массив транзакций
// Вывести массив

func second() {
	// transactions := [5]int{5, 10, -7, 2, 1}
	// banks := [2]string{"Тинькофф", "Альфа"}
	transactions := []float64{}

	// tr1 := []int{1, 2, 3}
	// tr2 := []int{4, 5, 6}
	// tr1 = append(tr1, 3, 4, 5)
	// tr2 = append(tr2, tr1...)

	// fmt.Println(tr1)
	// fmt.Println(tr2)

	// fmt.Println(transactions)
	// fmt.Println(banks)
	// part := transactions[2:4]
	// fmt.Println(part)
	fmt.Print("Введите транзакцию (n для выхода) \n")

	for {
		transaction := scanTransaction()

		if transaction == 0 {
			break
		}

		transactions = append(transactions, transaction)
	}

	sum := calculateBalance(transactions)
	fmt.Printf("Ваш индекс массы тела: %.04f \n", sum)

	// tr1 := []int{1, 2, 3}
	// tr2 := []int{4, 5, 6}
	// tr1 = append(tr1, 3, 4, 5)
	// tr2 = append(tr2, tr1...)

	// for index, value := range tr1 {
	// 	fmt.Println(index, value)
	// }

	tr := make([]string, 2, 10)
	tr[0] = "1"
	tr = append(tr, "1")
	tr = append(tr, "2")

}

func scanTransaction() float64 {
	var transaction float64
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	var sum float64 = 0
	for _, value := range transactions {
		sum += value
	}

	return sum
}
