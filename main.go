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

    app := cli.App("unfair", "chart potential biases in stream cipher keystreams")

    format := app.StringOpt("F format", "html", "output format (html or json)")

    app.Action = func() {
        sc := NewStreamCounter(256)

        streamer := NewRc4Streamer()

        for i := 0; i < 1000; i++ {
            sc.AddBytes(streamer.RandomKeyStream(256))
        }

        page := NewChartPage(sc)

        switch *format {
            case "html": page.PrintHtml()
            case "json": page.PrintJson()
        }
    }

    app.Run(os.Args)
}
