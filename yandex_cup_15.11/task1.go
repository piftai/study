package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	patients := make(map[int]float64)
	var totalTemperature float64
	var count int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "!" {
			break
		}
		var command string
		var id int

		if line == "?" {
			if count > 0 {
				avg := totalTemperature / float64(count)
				fmt.Printf("%.9f\n", avg)
			} else {
				fmt.Printf("n/a")
			}
			continue
		}
		var temp float64
		fmt.Sscanf(line, "%s", &command)

		switch command {
		case "+":
			fmt.Sscanf(line, "+ %d %f", &id, &temp)
			if _, exists := patients[id]; !exists {
				patients[id] = temp
				totalTemperature += temp
				count++
			}

		case "~":
			fmt.Sscanf(line, "~ %d %f", &id, &temp)
			if oldTemp, exists := patients[id]; exists {
				totalTemperature += temp - oldTemp
				patients[id] = temp
			}

		case "-":
			fmt.Sscanf(line, "- %d", &id)
			if oldTemp, exists := patients[id]; exists {
				totalTemperature -= oldTemp
				delete(patients, id)
				count--
			}
		}
	}
}
