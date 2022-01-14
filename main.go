package main

import (
	"os"
	"encoding/csv"
	"fmt"
	// "bufio"
	"strconv"
)

func main() {
	file, err := os.Open("problems.csv")
	if(err != nil){
		print(err)
	}
	reader := csv.NewReader(file)
	// inputReader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadAll()
	if(err != nil) {
		os.Exit(1)
	}

	for i := range data {
		var answer int
		var input int
		answer, _ = strconv.Atoi(string(data[i][1]))
		fmt.Println(data[i][0])
		fmt.Println("Enter answer:")
	  fmt.Scanf("%d", &input)

		if(input == answer){
			fmt.Println("CORRECT!")
		} else {
			fmt.Println("INCORRECT")
		}
	}

// data split by space, answer is at [1]
// randomly select and read from input if it's the correct answer
}

