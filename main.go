package main

import (
    "github.com/jawher/mow.cli"
    "os"
)

// we want to get n bytes from a random keystream for the stream ciphers
type RandomKeyStreamer interface {
    RandomKeyStream(numBytes int)
}

func main() {

    app := cli.App("stream-bias", "chart potential biases in stream cipher keystreams")

    app.Command("generate" func() {
        sc := NewStreamCounter(256)

        streamer := NewRc4Streamer()

        for i := 0; i < 1000000; i++ {
            sc.AddBytes(streamer.RandomKeyStream(256))
        }

        page := NewChartPage(sc)

        page.PrintHtml()
    }

    app.Run(os.Args)
}

