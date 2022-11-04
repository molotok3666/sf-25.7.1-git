package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var n string
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	var result interface{}

	if value, err := strconv.Atoi(n); err == nil {
		result = value
	} else if value, err := strconv.ParseBool(n); err == nil {
		result = value
	} else {
		result = n
	}

	fmt.Printf("Вы ввели число: %d\n", result)
}
