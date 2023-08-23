package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var fName string
	flag.StringVar(&fName, "f", "problems.csv", "The name of the csv that should be used")
	flag.Parse()

	f, err := os.Open(fName)
	check(err)
	defer f.Close()

	lines := readCSV(f)

	launchQuiz(lines)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readCSV(f *os.File) [][]string {
	r := csv.NewReader(f)
	entries, err := r.ReadAll()
	check(err)

	return entries
}

func launchQuiz(lines [][]string) {
	points := 0

	for _, v := range lines {
		fmt.Println(v[0])

		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		err := sc.Err()
		check(err)

		if sc.Text() == v[1] {
			points++
		}
	}

	fmt.Printf("Score: %v/%v", points, len(lines))
}
