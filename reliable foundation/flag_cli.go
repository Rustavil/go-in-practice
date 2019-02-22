package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	name = flag.String("name", "World", "Set name for say hello")
	lang string
)

func init() {
	flag.StringVar(&lang, "lang", "en", "language")
	flag.StringVar(&lang, "l", "en", "language")
}

func main() {
	flag.Parse()
	flag.Visit(func(i *flag.Flag) {
		fmt.Printf("Name %s value: %s usage: %s default: %s\n", i.Name, i.Value, i.Usage, i.DefValue)
	})
	switch strings.ToLower(lang) {
	case "en":
		fmt.Printf("Hello %v\n", *name)
	case "ru":
		fmt.Printf("Привет %v\n", *name)
	default:
		fmt.Printf("Hello %s\n", *name)
	}
}
