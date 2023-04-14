package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m, q int
	fmt.Scanf("%d %d %d", &n, &m, &q)

	// initializing data_centers
	DisabledServers := make([]bool, m)
	data_centers := make([][]bool, n)
	for i := range data_centers {
		data_centers[i] = append(data_centers[i], DisabledServers...)
	}
	R := make([]int, n)

	// running the commands
	for k := 0; k < q; k++ {
		var command string
		fmt.Fscanf(os.Stdin, "%s", command)

		switch {
		case command == "DISABLE":
			var i, j int
			fmt.Scanf("%d %d", &i, &j)

			data_centers[i-1][j-1] = true
		case command == "RESET":
			var i int
			fmt.Scanf("%d", &i)

			data_centers[i-1] = make([]bool, m)
			R[i-1]++
		case command == "GETMAX":
			getmax(R, data_centers)
		case command == "GETMIN":
			getmin(R, data_centers)
		}

		// fmt.Println(data_centers)
	}
}

// func errorCheck(err error) {
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

func getmax(R []int, data_centers [][]bool) {
	var max int
	var maxI int

	for i := len(data_centers) - 1; i >= 0; i-- {
		A := running_servers(data_centers[i])

		if R[i]*A >= max {
			max = R[i] * A
			maxI = i + 1
		}
	}

	fmt.Println(strconv.Itoa(maxI))
}

func getmin(R []int, data_centers [][]bool) {

	min := running_servers(data_centers[0]) * R[0]
	var minI int

	for i := len(data_centers) - 1; i >= 0; i-- {
		A := running_servers(data_centers[i])

		if R[i]*A <= min {
			min = R[i] * A
			minI = i + 1
		}
	}

	fmt.Println(strconv.Itoa(minI))
}

func running_servers(data_center []bool) int {
	var A int

	for _, disabled := range data_center {
		if !disabled {
			A++
		}
	}

	return A
}
