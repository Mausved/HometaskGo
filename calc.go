package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/Mausved/calculator/calculator"
	"os"
)

func main() {
	flag.Parse()
	output := bufio.NewWriter(os.Stdout)
	result := calculator.Calculator(flag.Arg(0))
	_, err := output.WriteString(fmt.Sprint(result) + "\n")
	if err != nil {
		fmt.Println("Error to write to output: err:", err)
		return
	}
	flushErr := output.Flush()
	if flushErr != nil {
		fmt.Println("Error to flush: err:", flushErr)
		return
	}

}
