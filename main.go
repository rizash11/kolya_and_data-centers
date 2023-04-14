package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f_input, err := os.Open("input.txt")
	errorCheck(err)
	defer f_input.Close()

	scanner := bufio.NewScanner(f_input)
	scanner.Scan()

	nmq := strings.Split(scanner.Text(), " ")
	n, err := strconv.Atoi(nmq[0])
	errorCheck(err)
	m, err := strconv.Atoi(nmq[1])
	errorCheck(err)
	q, err := strconv.Atoi(nmq[2])
	errorCheck(err)

	// initializing data_centers

	DisabledServers := make([]bool, m)
	data_centers := make([][]bool, n)
	for i := range data_centers {
		data_centers[i] = append(data_centers[i], DisabledServers...)
	}
	R := make([]int, n)

	// running the commands
	f_output, err := os.Create("output.txt")
	errorCheck(err)
	defer f_output.Close()
	w := bufio.NewWriter(f_output)

	for k := 0; k < q; k++ {
		scanner.Scan()
		command := strings.Split(scanner.Text(), " ")

		switch {
		case command[0] == "DISABLE":
			i, err := strconv.Atoi(command[1])
			errorCheck(err)
			j, err := strconv.Atoi(command[2])
			errorCheck(err)

			data_centers[i-1][j-1] = true
		case command[0] == "RESET":
			i, err := strconv.Atoi(command[1])
			errorCheck(err)

			data_centers[i-1] = make([]bool, m)
			R[i-1]++
		case command[0] == "GETMAX":
			getmax(w, R, data_centers)
		case command[0] == "GETMIN":
			getmin(w, R, data_centers)
		}

		// fmt.Println(data_centers)
	}

	w.Flush()
}

func errorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getmax(w *bufio.Writer, R []int, data_centers [][]bool) {
	var max int
	var maxI int

	for i := len(data_centers) - 1; i >= 0; i-- {
		A := running_servers(data_centers[i])

		if R[i]*A >= max {
			max = R[i] * A
			maxI = i + 1
		}
	}

	output := strconv.Itoa(maxI) + "\n"
	_, err := w.WriteString(output)
	errorCheck(err)
}

func getmin(w *bufio.Writer, R []int, data_centers [][]bool) {

	min := running_servers(data_centers[0]) * R[0]
	var minI int

	for i := len(data_centers) - 1; i >= 0; i-- {
		A := running_servers(data_centers[i])

		if R[i]*A <= min {
			min = R[i] * A
			minI = i + 1
		}
	}

	output := strconv.Itoa(minI) + "\n"
	_, err := w.WriteString(output)
	errorCheck(err)
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
