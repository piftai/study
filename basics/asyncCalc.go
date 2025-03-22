package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	ch := make(chan int)
	var expression string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку: ")

	// Читаем строку до символа новой строки ('\n')

	for {
		fmt.Println("Введите ваше выражение в формате '5 + 2'")
		expression, _ = reader.ReadString('\n')
		go read(expression, ch)
		fmt.Println(<-ch)
		time.Sleep(1 * time.Second)
	}
}

func read(expression string, result chan int) error {

	numbers := strings.Fields(expression)
	fNum, err := strconv.Atoi(numbers[0])
	sNum, err := strconv.Atoi(numbers[2])
	if err != nil {
		return errors.New("not valid numbers")
	}
	if len(numbers) != 3 {
		result <- 3
		return errors.New("not valid expression")
	}
	switch numbers[1] {
	case "+":
		result <- fNum + sNum
	case "-":
		result <- fNum - sNum
	case "*":
		result <- fNum * sNum
	case "/":
		result <- fNum / sNum
	}
	return nil
}
