package main

import (
	"flag"
	"fmt"
)

func main() {

	wrdptr := flag.Bool("w", false, "for words")
	flag.Parse()
	fmt.Println(*wrdptr)
	fmt.Println(flag.Args())

}
