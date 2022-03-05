package streams

import (
	"bufio"
	"flag"
	"io"
	"os"
)

func GetReadyInput() (io.Reader, error) {
	inputFile := flag.Arg(0)
	if inputFile != "" {
		f, err := os.Open(inputFile)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	return os.Stdin, nil
}

func GetReadyOutput() (io.Writer, error) {
	outputFile := flag.Arg(1)
	if outputFile != "" {
		f, err := os.Create(outputFile)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	return os.Stdout, nil
}

func ReadAllFile(buffer *bufio.Scanner) []string {
	var output []string
	for buffer.Scan() {
		output = append(output, buffer.Text())
	}
	return output
}
