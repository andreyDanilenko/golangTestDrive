package first

import (
	"errors"
	"fmt"
	"math"
)

func first() {
	fmt.Println("Калькулятор индекса массы тела")

	for {
		userKg, userHeight := fetUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			// fmt.Println("Не корректные данные")
			// continue
			panic("Не корректные данные")
		}
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()

		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.04f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас недостаток веса")
	case imt < 18.5:
		fmt.Println("У дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас ожирение")
	}
}

func calculateIMT(userKg, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func fetUserInput() (float64, float64) {
	var userHeight float64 // 0.0
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах \n")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес \n")
	fmt.Scan(&userKg)

	return userKg, userHeight
}

func checkRepeatCalculation() bool {
	var userChoose string
	fmt.Print("Вы хотите сделать еще расчет (y/n): ")
	fmt.Scan(&userChoose)
	if userChoose == "y" || userChoose == "Y" {
		return true
	}

	return false
}
