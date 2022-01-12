package main

import (
	"os"
	"encoding/csv"
	"fmt"
)

func main() {
file, err := os.Open("problems.csv")
if(err != nil){
	print(err)
}
reader := csv.NewReader(file)
data, err := reader.ReadAll()
fmt.Printf("%+v\n", data)
}
