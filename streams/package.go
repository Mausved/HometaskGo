package streams

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func GetReadyInput() io.Reader {
	inputFile := flag.Arg(0)
	if inputFile != "" {
		f, err := os.Open(inputFile)
		if err != nil {
			fmt.Println("Error opening file: err:", err)
			os.Exit(1)
		}
		return f
	}
	return os.Stdin
}

func GetReadyOutput() io.Writer {
	outputFile := flag.Arg(1)
	if outputFile != "" {
		f, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Error creating file: err:", err)
			os.Exit(1)
		}
		return f
	}
	return os.Stdout
}

func ReadAllFile(buffer *bufio.Scanner) []string {
	var output []string
	for buffer.Scan() {
		output = append(output, buffer.Text())
	}
	return output
}
