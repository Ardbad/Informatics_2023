package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/internal/lab4"
	"isuct.ru/informatics2022/internal/lab5"
)

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	fmt.Println("Ахмадеева Арина Игоревна")

	taskAResults := lab4.TaskA(2.0, 0.95, 1.25, 2.75, 0.3, 2.71828)
	fmt.Println("Результаты TaskA:", taskAResults)

	taskBResults := lab4.TaskB(2.0, 0.95, 2.71828, 2.2, 3.78, 4.51, 6.58, 1.2)
	fmt.Println("Результаты TaskB:", taskBResults)

	Dog, err := lab5.Thedog("Бочонок пива", 7, "01/01/2023")
	checkForError(err)

	fmt.Printf("Dog's age is %d\n", Dog.Getage())
	fmt.Printf("Dog's name is %s\n", Dog.Getname())
	fmt.Printf("Age is %d\n", Dog.Getage())
}
