package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	data     = flag.String("data", "", "input data")
	dataType = flag.String("type", "", "type of input data. enum: struct, string, int")
)

func main() {
	flag.Parse()
	if *data == "" || *dataType == "" {
		fmt.Fprintln(os.Stderr, errors.New("please specify an argument. use -h for help"))
		os.Exit(1)
	}

	if err := run(*data, *dataType); err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}

func run(data, dataType string) error {
	var output string
	var err error

	switch dataType {
	case "struct":
		output, err = encodeJSON([]byte(data))
	case "string":
		v := wrapperspb.String(data)
		output, err = encode(v)
	case "int":
		i, err := strconv.ParseInt(data, 10, 32)
		if err != nil {
			return err
		}
		v := wrapperspb.Int32(int32(i))
		output, err = encode(v)
	}

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", output)
	return nil
}

func encodeJSON(contents []byte) (string, error) {
	var m map[string]interface{}

	err := json.Unmarshal(contents, &m)
	if err != nil {
		return "", err
	}

	aStruct, err := structpb.NewStruct(m)
	if err != nil {
		return "", err
	}

	return encode(aStruct)
}

func encode(src protoreflect.ProtoMessage) (string, error) {
	any, err := anypb.New(src)
	if err != nil {
		return "", err
	}

	anyJSON, err := json.MarshalIndent(any, "", "\t")
	if err != nil {
		return "", err
	}

	return string(anyJSON), nil
}
