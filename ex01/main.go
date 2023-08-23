package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var points = 0
var remainingTime int

func main() {
	var name string
	var time int
	flag.StringVar(&name, "f", "problems.csv", "The name of the csv that should be used")
	flag.IntVar(&time, "t", 30, "The total time for the quiz")
	flag.Parse()

	f, err := os.Open(name)
	check(err)
	defer f.Close()

	remainingTime = time
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
	sc := bufio.NewScanner(os.Stdin)

	fmt.Print("Press enter when you are ready.")

	sc.Scan()
	err := sc.Err()
	check(err)

	go timer(len(lines))

	for _, v := range lines {
		fmt.Println(v[0])

		sc.Scan()
		err := sc.Err()
		check(err)

		if sc.Text() == v[1] {
			points++
		}
	}

	fmt.Printf("Score: %v/%v", points, len(lines))
}

func timer(len int) {
	for {
		time.Sleep(1 * time.Second)
		remainingTime--

		if remainingTime == 0 {
			break
		}
	}

	fmt.Printf("Score: %v/%v", points, len)
	os.Exit(0)
}
