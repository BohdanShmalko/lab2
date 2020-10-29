package main

import (
	"flag"
	"fmt"
	"strings"
	"io"
	"os"
	lab2 "github.com/BohdanShmalko/lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "The path to the file that stores the expression")
	outputFile = flag.String("o", "", "The path to the file that stores the calculation result")
)

func main() {
	flag.Parse()
	var (r io.Reader; w io.Writer; err error)
	if *inputFile != "" {
		r, err = os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR : ", err)
			os.Exit(2)
		}
	} else if *inputExpression == "" {
			fmt.Fprintln(os.Stderr, "ERROR : use -e or -f flag for input data (not empty)")
			os.Exit(2)
	}else {
		r = strings.NewReader(*inputExpression)
	}

	if *outputFile != "" {
		w, err = os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR : ", err)
			os.Exit(2)
		}
	}


	handler := &lab2.ComputeHandler{
	           Input: r,
	           Output: w,
	       }

	err = handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR : ", err)
	}
}
