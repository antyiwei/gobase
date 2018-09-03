package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {

	//flag.CommandLine = flag.NewFlagSet("question", flag.ExitOnError)
	//flag.CommandLine.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	//
	//flag.StringVar(&name, "name", "everyone", "The greeting object.")

	var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

	cmdLine.Parse(os.Args[1:])

}

func main() {

	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}

	flag.Parse()

	fmt.Printf("Hello, %s!\n", name)

}
