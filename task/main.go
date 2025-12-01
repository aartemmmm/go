package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Использование: go run . <температура_в_C>")
		os.Exit(1)
	}

	celsius, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка: аргумент должен быть числом")
		os.Exit(1)
	}

	fahrenheit := celsius*9/5 + 32
	fmt.Printf("%.1f °F\n", fahrenheit)
}
