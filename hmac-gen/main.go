package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	secret  = flag.String("secret", "", "secret")
	ts      = flag.String("ts", "", "timestamp")
	payload = flag.String("payload", "", "payload")
)

func main() {
	flag.Parse()
	if *secret == "" && (*ts == "" || *payload == "") {
		fmt.Fprintln(os.Stderr, errors.New("please specify an argument. use -h for help"))
		os.Exit(1)
	}

	if err := gen(*secret, *ts, *payload); err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}

func gen(secret string, ts string, payload string) error {

	h := hmac.New(sha256.New, []byte(secret))

	if ts != "" {
		h.Write([]byte(ts))
	}

	if payload != "" {
		h.Write([]byte(payload))
	}

	sha := base64.StdEncoding.EncodeToString(h.Sum(nil))

	fmt.Println("Result: " + sha)

	return nil
}
