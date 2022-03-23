package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/Mausved/uniq/streams"
	"github.com/Mausved/uniq/unique"
	"strings"
)

var opts unique.Options

func main() {
	flag.Parse()

	// prepare buffers
	input, closeInput, getInputErr := streams.GetReadyInput()
	defer func() {
		if closeInput != nil {
			closeInput()
		}
	}()
	if getInputErr != nil {
		fmt.Println("Error to get input buffer: err:", getInputErr)
		return
	}

	output, closeOutput, getOutputErr := streams.GetReadyOutput()
	defer func() {
		if closeOutput != nil {
			closeOutput()
		}
	}()
	if getOutputErr != nil {
		fmt.Println("Error to get input buffer: err:", getOutputErr)
		return
	}

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	text := streams.ReadAllFile(scanner)
	result, errorFlags := unique.Uniq(text, opts)

	if errorFlags != nil {
		flag.PrintDefaults()
		return
	}

	if _, err := writer.WriteString(strings.Join(result, "\n") + "\n"); err != nil {
		fmt.Println("Error to writing to writer: err: ", err)
	}

	if errFlush := writer.Flush(); errFlush != nil {
		fmt.Println("Error to flush: err: ", errFlush)
		return
	}
}

func init() {
	const (
		usageC        = "counting row."
		usageD        = "print only repeating rows."
		usageU        = "print only unique rows. "
		usageNumField = "search only after first 'num_field' fields in row."
		usageNumChars = "search only after first 'num_chars' chars in row."
		usageI        = "not sensitive mode."
	)
	flag.BoolVar(&opts.Count, "c", false, usageC)
	flag.BoolVar(&opts.Double, "d", false, usageD)
	flag.BoolVar(&opts.Uniq, "u", false, usageU)
	flag.IntVar(&opts.NumFields, "f", 0, usageNumField)
	flag.IntVar(&opts.NumChars, "s", 0, usageNumChars)
	flag.BoolVar(&opts.Insensitive, "i", false, usageI)
}
