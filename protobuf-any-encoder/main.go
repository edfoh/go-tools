package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

var (
	dataInput = flag.String("data-input", "", "input data")
	dataFile  = flag.String("data-file", "", "input data file")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}

func run() error {
	var contents []byte
	if *dataFile != "" {
		d, err := readFile(*dataFile)
		if err != nil {
			return err
		}
		contents = d
	} else if *dataInput != "" {
		contents = []byte(*dataInput)
	} else {
		return fmt.Errorf("please specify an argument. use -h for help")
	}

	output, err := encode(contents)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", output)
	return nil
}

func encode(contents []byte) (string, error) {
	var m map[string]interface{}

	err := json.Unmarshal(contents, &m)
	if err != nil {
		return "", err
	}

	aStruct, err := structpb.NewStruct(m)
	if err != nil {
		return "", err
	}

	any, err := anypb.New(aStruct)
	if err != nil {
		return "", err
	}

	anyJSON, err := json.MarshalIndent(any, "", "\t")
	if err != nil {
		return "", err
	}

	return string(anyJSON), nil
}

func readFile(fileLoc string) ([]byte, error) {
	if _, err := os.Stat(fileLoc); os.IsNotExist(err) {
		return []byte{}, errors.New("unable to read file")
	}

	b, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		return []byte{}, fmt.Errorf("open input: %v", err)
	}
	return b, nil
}
