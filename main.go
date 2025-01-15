package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// define flags and tell flag package to parse all of them

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file in the format of 'problem,solution'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	//make sure you're using actual value from the strig, and not the pointer
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV File. ")
	}
	fmt.Println(lines)

}
