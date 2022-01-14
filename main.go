package main

import (
	"os"
	"encoding/csv"
	"fmt"
	"strconv"
)

func main() {
	fileToRead := selectFile()
	file, err := os.Open(fileToRead)
	if err != nil {
		print(err)
		os.Exit(1)
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		os.Exit(1)
	}

	correctAnswers, incorrectAnswers := 0, 0

	for i := range data {
		var answer int
		var input int
		answer, _ = strconv.Atoi(string(data[i][1]))
		fmt.Println(data[i][0])
		fmt.Println("Enter answer:")
	  fmt.Scanf("%d", &input)

		result := checkAnswer(input, answer)

		if result {
			correctAnswers++
		} else {
			incorrectAnswers++
		}
	}
	checkFinalResults(correctAnswers, incorrectAnswers)
}

func checkAnswer(input int, answer int) bool {
	if(input == answer){
		fmt.Println("CORRECT!")
		return true
	} else {
		fmt.Println("INCORRECT")
		return false
	}
}

func checkFinalResults(correctAnswers int, incorrectAnswers int){
	if incorrectAnswers != 0 {
		percentage := correctAnswers / incorrectAnswers
		fmt.Print(percentage)
		fmt.Printf("You got %d%% right!", percentage)
		os.Exit(0)
	} else {
			fmt.Printf("YOU GOT THEM ALL RIGHT!")
	}
}

// TODO validate filepath

func selectFile() string {
	var filePath string		
	fmt.Println("What file do you want to read?")
	fmt.Scanf("%v", &filePath)
	return filePath
}