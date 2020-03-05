package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Declare the flag using the corresponding package
	//Usage flag.String(flagname, default,help output)
	csvFileName := flag.String("csv", "problems.csv",
		"a csv file in the format of 'question,answer'")

	//parse the flag provided
	flag.Parse()

	//Open the file
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	// I/O reader
	// It is the most common interface
	//It cal also be used for reading HTTP respnses or other byte slices

	r := csv.NewReader(file)

	//We don't need to worry about memory overflow in that project,
	//so we are going to read the whole file at once
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file")
	}

	problems := parseLines(lines)

	correct := 0

	//Range: provides both index and value (i.e. i, p) for each entry
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answer string

		//It is not appropriate for other cases (i.e string answers)
		//Trims the spaces
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			fmt.Println("Correct")
			correct++
		}

	}

	fmt.Printf("You scored %d out of %d", correct, len(problems))
}

func parseLines(lines [][]string) []problem {

	//We declare here the required size of ret and not just the
	//zero value. In the second case, we would append every new
	//entry. However, there is no reason here to add this overhead
	// in the execution time

	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]), // trims also the spaces in the file
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
