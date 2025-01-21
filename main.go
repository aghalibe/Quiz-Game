package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// define flags and tell flag package to parse all of them

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file in the format of 'question,solution'")
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
	questions := parseLines(lines)
	//fmt.Println(questions)

	//keep count of how many questions are correct
	correct := 0
	for i, p := range questions {
		fmt.Printf("question #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(questions))
}

//build a question via csv 2d slice
func parseLines(lines [][]string) []question {
	ret := make([]question, len(lines))

	for i, line := range lines {
		ret[i] = question{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
type question struct {
	q string
	a string
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
