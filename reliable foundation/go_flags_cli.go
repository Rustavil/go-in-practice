package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"strings"
)

var opts struct {
	Name string `short:"n" long:"name" default:"World"`
	Lang string `short:"l" long:"lang" description:"Use Spanish Language"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		return
	}
	switch strings.ToLower(opts.Lang) {
	case "en":
		fmt.Printf("Hello %s\n", opts.Name)
	case "ru":
		fmt.Printf("Привет %s\n", opts.Name)
	default:
		fmt.Printf("Hello %s\n", opts.Name)
	}
}
