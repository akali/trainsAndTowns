package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"transAndTowns/cmd"
)

var file = flag.String("file", "", "Input file")

func main() {
	flag.Parse()
	var reader *bufio.Reader = nil

	if *file == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(*file)
		if err != nil {
			panic(fmt.Errorf("can't read file %s", *file))
		}
		defer f.Close()
		reader = bufio.NewReader(f)
	}
	if err := cmd.Run(reader); err != nil {
		panic(err)
	}
}
