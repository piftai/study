package main

import (
	"fmt"
	"log"
)

func main() {
	var x, y float64
	fmt.Print("Введите значение x: ")
	_, err := fmt.Scan(&x)
	if err != nil {
		log.Fatalf("error trying put uncorrect X value: %v", err)
	}
	fmt.Print("Введите значение y: ")
	_, err = fmt.Scan(&y)
	if err != nil {
		log.Fatalf("error trying put uncorrect Y value: %v", err)
	}
	fmt.Printf("Точка лежит %s\n", half(x, y))
}

func half(x float64, y float64) string {
	response := ""
	if x > 0 && y > 0 {
		response = "на I четверти"
	}
	if x < 0 && y > 0 {
		response = "на II четверти"
	}
	if x < 0 && y < 0 {
		response = "на III четверти"
	}
	if x > 0 && y < 0 {
		response = "на IV четверти"
	}
	if x == 0 {
		response = "на оси Y"
	}
	if y == 0 {
		response = "на оси X"
	}
	return response
}
