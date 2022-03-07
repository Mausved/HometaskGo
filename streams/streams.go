package streams

import (
	"bufio"
	"flag"
	"io"
	"os"
)

func GetReadyInput() (io.Reader, func() error, error) {
	inputFile := flag.Arg(0)
	if inputFile != "" {
		f, err := os.Open(inputFile)
		if err != nil {
			f.Close()
			return nil, nil, err
		}
		return f, f.Close, nil
	}
	return os.Stdin, nil, nil
}

func GetReadyOutput() (io.Writer, func() error, error) {
	outputFile := flag.Arg(1)
	if outputFile != "" {
		f, err := os.Create(outputFile)
		if err != nil {
			f.Close()
			return nil, nil, err
		}
		return f, f.Close, nil
	}
	return os.Stdout, nil, nil
}

func ReadAllFile(buffer *bufio.Scanner) []string {
	var output []string
	for buffer.Scan() {
		output = append(output, buffer.Text())
	}
	return output

}
