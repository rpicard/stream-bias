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
    sc := NewStreamCounter(256)

    streamer := NewRc4Streamer()

    sc.AddBytes(streamer.RandomKeyStream)

    jsn, err := json.Marshal(sc.Count)

    if err != nil {
        log.Fatal(err)
    }

    page := NewChartPage(string(jsn))

    page.PrintHtml()
}

