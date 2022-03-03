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
	if opts.Uniq && opts.Double || opts.Uniq && opts.Count || opts.Double && opts.Count {
		flag.PrintDefaults()
		return
	}

	// prepare buffers
	buffer := bufio.NewScanner(streams.GetReadyInput())
	output := bufio.NewWriter(streams.GetReadyOutput())
	text := streams.ReadAllFile(buffer)

	result := unique.Uniq(text, opts)
	_, err := output.WriteString(strings.Join(result, "\n"))
	if err != nil {
		fmt.Println("Error to writing to output: err: ", err)

	}

	errFlush := output.Flush()
	if errFlush != nil {
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
