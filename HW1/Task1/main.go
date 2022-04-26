// Для закрепления навыков отложенного вызова функций, напишите программу, содержащую вызов функции, 
// которая будет создавать паническую ситуацию неявно. 
// Затем создайте отложенный вызов, который будет обрабатывать эту паническую ситуацию и, в частности, печатать предупреждение в консоль. 
// Критерием успешного выполнения задания является то, что программа не завершается аварийно ни при каких условиях.


package main

import(
	"fmt"
)

func main() {
	defer fmt.Println("end of pogram")
	defer func() {
     if rcvr := recover(); rcvr != nil {
        fmt.Println("recovered", rcvr)
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

func GetNumberFromSliceToIndex(idx int, s []int) int{
	return s[idx]
}