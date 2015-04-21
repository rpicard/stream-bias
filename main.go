package main

import (
    "github.com/jawher/mow.cli"
    "os"
    "log"
    "errors"
)

// we want to get n bytes from a random keystream for the stream ciphers
type RandomKeyStreamer interface {
    RandomKeyStream(numBytes int) []byte
}

func main() {

    app := cli.App("unfair", "chart potential biases in stream cipher keystreams")

    format := app.StringOpt("f format", "html", "output format (html or json)")
    cipher := app.StringOpt("c cipher", "", "which cipher are we testing?")

    app.Spec = "--cipher [--format]"

    app.Action = func() {
        sc := NewStreamCounter(256)

        var streamer RandomKeyStreamer

        // add new ciphers here
        switch *cipher {
            case "rc4": streamer = NewRc4Streamer()
            default: log.Fatal(errors.New("I do not know that cipher."))
        }

        for i := 0; i < 1000; i++ {
            sc.AddBytes(streamer.RandomKeyStream(256))
        }

        // print out the data as either html charts or json
        page := NewChartPage(sc)

        switch *format {
            case "html": page.PrintHtml()
            case "json": page.PrintJson()
        }
    }

    app.Run(os.Args)
}
