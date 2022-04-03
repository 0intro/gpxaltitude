package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	abs     = flag.Float64("a", 0, "set absolute altitude")
	rel     = flag.Float64("r", 0, "set relative altitude change")
	verbose = flag.Bool("v", false, "verbose")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gpxaltitude [ -v ] [ -a n | -r n ] file\n")
	os.Exit(2)
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		usage()
	}

	if *abs != 0 && *rel != 0 {
		usage()
	}

	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile("<ele>(-?[0-9.]+)</ele>")

	var start *float64
	var count uint64

	fixElevation := func(m string) string {
		s := r.FindStringSubmatch(m)
		if len(s) < 2 {
			return ""
		}

		ele, err := strconv.ParseFloat(s[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		if start == nil {
			start = new(float64)
			*start = ele
		}

		var fix float64
		if *abs != 0 {
			fix = *abs - *start
		}
		if *rel != 0 {
			fix = *rel
		}

		if *verbose {
			fmt.Fprintf(os.Stderr, "%f → %0.2f\n", ele, ele+fix)
		}
		count++

		e := fmt.Sprintf("<ele>%0.2f</ele>", ele+fix)

		return e
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(r.ReplaceAllStringFunc(scanner.Text(), fixElevation))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if *verbose {
		fmt.Fprintf(os.Stderr, "fixed %d altitudes\n", count)
	}
}
