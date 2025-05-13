package main

import "fmt"

func main() {
	var cntTrains int
	if _, err := fmt.Scan(&cntTrains); err != nil {
		return
	}
	consts := make([]int, cntTrains)
	mnozh := make([]int, cntTrains)

	for i := 0; i < cntTrains; i++ {
		if _, err := fmt.Scan(&consts[i], &mnozh[i]); err != nil {
			return
		}
	}
	var q int
	fmt.Scan(&q)
	output := make([]int, 0, q)

	for i := 0; i < q; i++ {
		var ti, di int
		fmt.Scan(&ti, &di)
		ti--

		ai := consts[ti]
		bi := mnozh[ti]

		if di <= ai {
			output = append(output, ai)
			continue
		}

		k := (di - ai + bi - 1) / bi
		nextTrain := ai + k*bi
		output = append(output, nextTrain)
	}
	for _, v := range output {
		fmt.Println(v)
	}
}
