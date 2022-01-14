package main

import (
	"os"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"
	"math"
)

func main() {
	fileToRead := selectFile()
	file, err := os.Open(fileToRead)
	if err != nil {
		err := fmt.Errorf("file: %q not found ", fileToRead)
		fmt.Print(err)
		os.Exit(1)
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	totalProblems := len(data)
	if err != nil {
		os.Exit(1)
	}

	correctAnswers, incorrectAnswers := 0, 0

	start := time.Now()

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
	t := time.Now()
	elapsed := t.Sub(start)
	checkFinalResults(correctAnswers, incorrectAnswers, elapsed, totalProblems)
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

func checkFinalResults(correctAnswers int, incorrectAnswers int, time time.Duration, totalProblems int){
	if incorrectAnswers != 0 {
		percentage := math.Round((float64(correctAnswers) / float64(totalProblems)) * 100)
		fmt.Printf("You got %v%% right!\n", percentage)
		fmt.Printf("It took you %v\n", time)
		os.Exit(0)
	} else {
			fmt.Printf("YOU GOT THEM ALL RIGHT!")
			fmt.Printf("It took you %v\n", time)
	}
}

func selectFile() string {
	var filePath string		
	fmt.Println("What file do you want to read?")
	fmt.Scanf("%v", &filePath)
	return filePath
}