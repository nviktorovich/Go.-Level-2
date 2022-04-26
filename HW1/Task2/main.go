// Дополните функцию из п.1 возвратом собственной ошибки в случае возникновения панической ситуации.
// Собственная ошибка должна хранить время обнаружения панической ситуации.
// Критерием успешного выполнения задания является наличие обработки созданной ошибки в функции main и вывод ее состояния в консоль

package main

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
)

type ErrorWithTrace struct {
	text  string
	trace string
}

func New(text string) error {
	return &ErrorWithTrace{
		text:  text,
		trace: string(debug.Stack()),
	}
}

func (e *ErrorWithTrace) Error() string {
	return fmt.Sprintf("error: %s\ntrace:\n%s", e.text, e.trace)
}

func main() {

	defer fmt.Println("end of pogram")
	defer func() {
		if rcvr := recover(); rcvr != nil {
			fmt.Println("recovered", rcvr)
			err := errors.New("Self error")
			log.Println(err)
		}
	}()

	var s []int = []int{1, 2, 3, 4, 5}
	var userIdx int
	_, err := fmt.Scanln(&userIdx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(GetNumberFromSliceToIndex(userIdx, s))
}

func GetNumberFromSliceToIndex(idx int, s []int) int {
	return s[idx]
}
