package main

import (
    "log"
    "encoding/json"
)

// we want to get n bytes from a random keystream for the stream ciphers
type RandomKeyStreamer interface {
    RandomKeyStream(numBytes int)
}

// the stream ciphers will use this type
type StreamCipher struct {
    Name    string
}

func main() {
    sc := NewStreamCounter(5)
    sc.AddBytes([]byte{0x00, 0x01, 0x03, 0x05, 0x01})

    jsn, err := json.Marshal(sc.Count)

    if err != nil {
        log.Fatal(err)
    }

    page := NewChartPage(string(jsn))

    page.PrintHtml()
}

