// Для закрепления практических навыков программирования, напишите программу,
// которая создаёт один миллион пустых файлов в известной, пустой директории файловой системы используя вызов os.Create.
// Ввиду наличия определенных ограничений операционной системы на число открытых файлов,
// такая программа должна выполнять аварийную остановку. Запустите программу и дождитесь полученной ошибки.
// Используя отложенный вызов функции закрытия файла, стабилизируйте работу приложения.
// Критерием успешного выполнения программы является успешное создание миллиона пустых файлов в директории.

package main

import (
	// "fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	for idx := 0; idx < 1000000; idx++ {
		name := strconv.Itoa(idx)
		crateAndClose(name)
	}
}

func crateAndClose(n string) {
	file, err := os.Create("sandbox/" + n + ".txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
}
