package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var points = 0
var remainingTime int

func main() {
	var name string
	var time int
	var shuffle bool
	flag.StringVar(&name, "f", "problems.csv", "The name of the csv that should be used")
	flag.IntVar(&time, "t", 30, "The total time for the quiz")
	flag.BoolVar(&shuffle, "s", false, "If the questions should be shuffled or not")
	flag.Parse()

	f, err := os.Open(name)
	check(err)
	defer f.Close()

	remainingTime = time
	lines := readCSV(f)

	if shuffle {
		rand.Shuffle(len(lines), func(i, j int) {
			lines[i], lines[j] = lines[j], lines[i]
		})
	}

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

		res := sc.Text()
		res = strings.TrimSpace(res)
		res = strings.ToLower(res)

		if res == v[1] {
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
