package indent

import (
	"bytes"
	"fmt"
	"io"
	"jsonfmt/decode"
	"log"
	"os"
	"testing"
)

func TestIndentEndtoEnd(t *testing.T) {

	var bufIn bytes.Buffer
	var bufOut bytes.Buffer

	// Open file and read into buffer.
	fi, err := os.Open("../testfiles/test1.json")
	if err != nil {
		log.Fatal(err)
	}
	readBytes := make([]byte, 1024)
	for {
		n, err := fi.Read(readBytes)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		bufIn.Write(readBytes[:n])
	}
	fi.Close()

	obj, err := decode.DecodeJSON(bufIn.Bytes())
	if err != nil {
		t.Errorf("DecodeJSON returned error; possible syntax error")
		return
	}
	Indent(&bufOut, obj, "    ", false)

	// Trim newlines off right of file
	for index, expect := range bytes.TrimRight(bufIn.Bytes(), "\n") {
		if index > len(bufOut.Bytes())+1 {
			t.Errorf("bufOut was smaller than bufIn!")
			break
		}
		result := bufOut.Bytes()[index]
		if expect != result {
			t.Errorf("Expecting '%s', got '%s' at byte %d", string(expect),
				string(result), index)
			fmt.Println(string(bufOut.Bytes()))
			break
		}
	}
}
