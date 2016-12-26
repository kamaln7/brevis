package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"os"
	"regexp"
	"strconv"

	"github.com/kamaln7/brevis/circularslice"

	"github.com/aybabtme/log"
)

var (
	keyRegexp = regexp.MustCompile(`^key\s*release\s*(\d+)\s*$`)
	keymap    map[int]string
	input     *circularslice.Slice
)

func parseKeymap(serialized string) {
	buffer := new(bytes.Buffer)
	data, err := base64.StdEncoding.DecodeString(serialized)
	if err != nil {
		log.Err(err).Fatal("could not decode keymap")
	}
	buffer.Write(data)
	ungobber := gob.NewDecoder(buffer)
	err = ungobber.Decode(&keymap)
	if err != nil {
		log.Err(err).Fatal("could not decode keymap")
	}
}

func initInputRing() {
	input = circularslice.New(4)
}

func main() {
	// read keymap
	if len(os.Args) != 2 {
		log.Fatal("please pass the serialized keymap as the first argument")
	}

	parseKeymap(os.Args[1])
	initInputRing()

	// listen for key events

	reader := bufio.NewReader(os.Stdin)
	for {
		keyEvent, err := reader.ReadString('\n')
		if err != nil {
			log.Err(err).Fatal("could not read event")
		}
		if keyRegexp.MatchString(keyEvent) {
			matches := keyRegexp.FindStringSubmatch(keyEvent)
			keyCode, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Err(err).Error("could not convert keycode to int")
				continue
			}

			process(keyCode)
		}
	}
}

func process(keyCode int) {
	char, exists := keymap[keyCode]
	if !exists {
		input.Clear()
		return
	}
	log.KV("code", keyCode).KV("character", char).Info("got key")
	input.Insert(char)

	expand()
}

func expand() {
	var text string
	for _, c := range input.Get() {
		text += c.(string)
	}

	if text == "test" {
		log.Info("test")
	}
}
