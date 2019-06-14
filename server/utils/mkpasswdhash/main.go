package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var (
	cost = flag.Int("cost", bcrypt.DefaultCost, "bcrypt cost")
)

func ane(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	_, err := fmt.Fprint(os.Stderr, "password: ")
	ane(err)
	password, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	ane(err)
	password = bytes.TrimSpace(password)
	hash, err := bcrypt.GenerateFromPassword(password, *cost)
	ane(err)
	_, err = fmt.Printf("%s\n", hash)
	ane(err)
}
