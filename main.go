package main

import (
    "log"
    "encoding/json"
)

// we want to get n bytes from a random keystream for the stream ciphers
type RandomKeyStreamer interface {
    RandomKeyStream(numBytes int)
}

func main() {
    sc := NewStreamCounter(256)

    streamer := NewRc4Streamer()

    for i := 0; i < 10000000; i++ {
        sc.AddBytes(streamer.RandomKeyStream(256))
    }

    jsn, err := json.Marshal(sc.Probability)

    if err != nil {
        log.Fatal(err)
    }

    page := NewChartPage(string(jsn))

    page.PrintHtml()
}

