package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	count uint
	help  bool
)

func init() {
	flag.UintVar(&count, "c", 1, "number of times to repeat the input")
	flag.BoolVar(&help, "h", false, "help")
}

func main() {
	flag.Parse()
	if !flag.Parsed() {
		log.Printf("unable to parse command line flags")
		os.Exit(1)
	}
	if help {
		flag.PrintDefaults()
		return
	}
	pipe := make(chan string)

	go func() {
		err := input(pipe)
		if err != nil {
			log.Printf("trouble reading from stdin: %v", err)
		}
		close(pipe)
	}()
	repeat(pipe, count)
}

func repeat(pipe <-chan string, count uint) {
	for line := range pipe {
		for n := uint(0); n < count; n++ {
			fmt.Printf("%s", line)
		}
	}
}

func input(pipe chan<- string) error {
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		pipe <- line
	}
	return nil
}
