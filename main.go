package main

import (
	"fmt"
	"os"

	a "github.com/ktateish/go-module-exp-lib-a"
	b "github.com/ktateish/go-module-exp-lib-b"
)

func main() {
	if err := Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}

func Run(args []string) error {
	fmt.Println(a.F("foo"))
	fmt.Println(b.F("foo"))
	return nil
}
