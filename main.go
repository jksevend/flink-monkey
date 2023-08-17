package main

import (
	"flink_monkey/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("FlinkMonkey v%s\n", VERSION)
	repl.Start(os.Stdin, os.Stdout)
}
