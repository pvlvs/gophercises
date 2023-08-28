package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func startTerminal() {
	sm := storyMap()

	nextArc(sm, "intro")
}

func nextArc(sm map[string]storyArc, arc string) {
	sa := sm[arc]
	options := []string{}

	fmt.Println(sa.Title)

	for _, v := range sa.Story {
		fmt.Printf("\n%s\n", v)
	}

	fmt.Println()

	if arc == "home" {
		os.Exit(0)
	}

	for i, v := range sa.Options {
		options = append(options, v.Arc)
		fmt.Printf("%d) %s\n", i+1, v.Text)
	}

	pick := readOption(options)

	nextArc(sm, options[pick-1])
}

func readOption(options []string) int {
	fmt.Println()
    fmt.Print("Your pick: ")
	sc.Scan()
	err := sc.Err()
	if err != nil {
		log.Fatal(err)
	}

	res := sc.Text()
	res = strings.TrimSpace(res)

	pick, err := strconv.Atoi(res)
	if err != nil {
		fmt.Printf("'%s' is not a valid option! Try again.\n", res)
		pick = readOption(options)
	}

	if pick > len(options) || pick <= 0 {
		fmt.Printf("'%s' is not a valid option! Try again.\n", res)
		pick = readOption(options)
	}

	return pick
}
