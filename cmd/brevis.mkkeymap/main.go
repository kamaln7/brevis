package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var alphaRegex = regexp.MustCompile(`(\d+)\s+=\s+([a-z])\s+`)

func main() {
	// read
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	keymapAll := strings.Split(string(data), "\n")

	// poor man's filter & map
	keymap := make(map[int]string, 26)
	for _, key := range keymapAll {
		if alphaRegex.MatchString(key) {
			matches := alphaRegex.FindStringSubmatch(key)
			keyCode, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			keymap[keyCode] = matches[2]
		}
	}

	// serialize
	buffer := new(bytes.Buffer)
	gobber := gob.NewEncoder(buffer)
	err = gobber.Encode(keymap)
	if err != nil {
		panic(err)
	}

	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	buffer.WriteTo(encoder)
	encoder.Close()
}
