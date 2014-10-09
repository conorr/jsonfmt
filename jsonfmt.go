package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"jsonfmt/decode"
	"jsonfmt/indent"
	"jsonfmt/util"
	"os"
	"regexp"
)

const JSONP_RE string = "(?s)^([\n]?[A-Za-z_0-9.]+[(]{1}[\n]?)(.*)([)]{1}[\n]?)$"

func main() {

	var opts struct {
		Sort bool `short:"s" long:"sort" description:"Sort keys alphabetically"`
	}

	args, _ := flags.Parse(&opts)
	if len(args) < 1 {
		fmt.Println("Usage: jsonfmt [-s] [--sort] [file]")
		os.Exit(1)
	}
	filename := args[0]

	inBuf := util.ReadFile(filename)
	outBuf := JSONFmt(inBuf, opts.Sort)

	err := util.WriteFile(filename, outBuf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func JSONFmt(body *bytes.Buffer, sortKeys bool) *bytes.Buffer {

	var head bytes.Buffer
	var tail bytes.Buffer

	// Try parsing JSONP.
	if parts, err := ParseJSONP(body.Bytes()); err == nil {
		head.Write(parts[0])
		body.Reset()
		body.Write(parts[1])
		tail.Write(parts[2])
	}

	// Make a new buffer of indented JSON.
	// TODO: need to initialize like this?
	indentedBody := bytes.NewBufferString("")
	obj, err := decode.DecodeJSON(body.Bytes())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	indent.Indent(indentedBody, obj, "    ", sortKeys)

	var result bytes.Buffer

	result.Write(head.Bytes())
	result.Write(indentedBody.Bytes())
	result.Write(tail.Bytes())

	return &result
}

func ParseJSONP(contents []byte) ([][]byte, error) {
	re, _ := regexp.Compile(JSONP_RE)
	matches := re.FindAllSubmatch(contents, -1)
	if len(matches) == 0 {
		return nil, errors.New("Could not parse into JSONP")
	}
	parts := matches[0]
	if len(parts) < 3 {
		return nil, errors.New("Could not parse into JSONP")
	}
	return parts[1:], nil
}
