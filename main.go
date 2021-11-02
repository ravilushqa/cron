package main

import (
	"errors"
	"log"
	"os"
	"strings"
)

var errWrongArgs = errors.New("pass cron arguments like: cron \"*/15 0 1,15 * 1-5 /usr/bin/find\"")

func main() {
	if len(os.Args) != 2 {
		log.Fatal(errWrongArgs)
	}

	args := strings.Split(strings.Trim(os.Args[1], " "), " ")

	if len(args) < 6 {
		log.Fatal(errWrongArgs)
	}

	c := newCron(args)
	err := c.parseFields()
	if err != nil {
		log.Fatal(err)
		return
	}

	_ = c.prettyPrint(os.Stdout)
}
