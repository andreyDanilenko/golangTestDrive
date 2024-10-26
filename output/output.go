package output

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(value any) {
	intVal, ok := value.(int)
	if ok {
		color.Red("Код ошибки %d", intVal)
		return
	}
	strVal, ok := value.(string)
	if ok {
		color.Red(strVal)
		return
	}

	errVal, ok := value.(error)
	if ok {
		color.Red(errVal.Error())
		return
	}

	color.Red("Неизвестный тип ошибки")

	// switch t := value.(type) {
	// case string:
	// 	color.Red(t)
	// case int:
	// 	color.Red("Код ошибки %d", t)
	// default:
	// 	color.Red("Неизвестный тип ошибки")
	// }
}

func sum[T int | string](a, b T) T {
	switch d := any(a).(type) {
	case string:
		fmt.Println(d)
	}
	return a + b
}

type List[T any] struct {
	elements []T
}

func (l *List[T]) addElement() {

}

// // Дженерик обощенный тип позволяет использовать функцию с обобщенными типами
// func sumInt(a, b int) int {
// 	return a + b
// }

// func sumFloat32(a, b float32) float32 {
// 	return a + b
// }

// func sumFloat64(a, b float64) float64 {
// 	return a + b
// }
