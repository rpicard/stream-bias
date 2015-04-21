package main

import (
    "github.com/jawher/mow.cli"
    "os"
    "log"
    "errors"
    "strconv"
)

// we want to get n bytes from a random keystream for the stream ciphers
type RandomKeyStreamer interface {
    RandomKeyStream(numBytes int) []byte
}

func main() {

    app := cli.App("unfair", "chart potential biases in stream cipher keystreams")

    format := app.StringOpt("f format", "html", "output format (html or json)")
    cipher := app.StringOpt("c cipher", "", "which cipher are we testing?")
    samples := app.StringOpt("s samples", "1000", "how many samples should we take?")
    length := app.StringOpt("l length", "256", "how many positions in the keystream do we care about?")

    app.Spec = "--cipher [--format] [--samples] [--length]"

    app.Action = func() {
        sc := NewStreamCounter(256)

        var streamer RandomKeyStreamer

        // add new ciphers here
        switch *cipher {
            case "rc4": streamer = NewRc4Streamer()
            default: log.Fatal(errors.New("I do not know that cipher."))
        }

        // parse the number of samples and stream length into int64s
        numSamples, err := strconv.ParseInt(*samples, 10, 64)
        if err != nil {
            log.Fatal(err)
        }

        streamLength, err := strconv.Atoi(*length)
        if err != nil {
            log.Fatal(err)
        }

        for i := int64(0); i < numSamples; i++ {
            sc.AddBytes(streamer.RandomKeyStream(streamLength))
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
