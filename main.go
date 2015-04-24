package main

import (
	"errors"
    "encoding/json"
	"github.com/jawher/mow.cli"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
)

// we want to get n bytes from a random keystream for the stream ciphers
type RandomKeyStreamer interface {
	RandomKeyStream(numBytes int) []byte
}

func main() {

	// set GOMAXPROCS to the number of CPU cores available
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.App("unfair", "chart potential biases in stream cipher keystreams")

	format := app.StringOpt("f format", "html", "output format (html or json)")

	app.Command("generate", "generate data", func(cmd *cli.Cmd) {

        cipher := cmd.StringOpt("c cipher", "", "which cipher are we testing?")
        samples := cmd.StringOpt("s samples", "1000", "how many samples should we take?")
        length := cmd.StringOpt("l length", "256", "how many positions in the keystream do we care about?")

        cmd.Action = func() {
            var streamer RandomKeyStreamer

            // add new ciphers here
            switch *cipher {
            case "rc4":
                streamer = NewRc4Streamer()
            case "hc128":
                streamer = NewHc128Streamer()
            case "sosemanuk":
                streamer = NewSosemanukStreamer()
            default:
                log.Fatal(errors.New("I do not know that cipher."))
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

            var wg sync.WaitGroup

            counters := make(chan *StreamCounter, 100)

            // generate a goroutine for each CPU
            for i := 0; i < runtime.NumCPU(); i++ {
                wg.Add(1)

                go func() {

                    defer wg.Done()

                    sc := NewStreamCounter(streamLength)

                    // loop through our share of the samples
                    for j := int64(0); j <= (numSamples / int64(runtime.NumCPU())); j++ {
                        sc.AddBytes(streamer.RandomKeyStream(streamLength))
                    }

                    counters <- sc
                }()
            }

            sc := NewStreamCounter(streamLength)

            // read our samples from the channel
            for i := 0; i < runtime.NumCPU(); i++ {
                sc.AddCounter(<-counters)
            }

            // wait for the goroutines to generate all of their samples
            wg.Wait()
            close(counters)

            // print out the data as either html charts or json
            page := NewChartPage(sc)

            switch *format {
            case "html":
                page.PrintHtml()
            case "json":
                page.PrintJson()
            }
        }
	})

    app.Command("import", "import data from json files", func(cmd *cli.Cmd) {

        filesArg := cmd.StringsArg("FILES", nil, "the files to import")

        cmd.Spec = "FILES..."

        cmd.Action = func() {
            files := *filesArg

            // read the first file, this will be our master counter
            //
            // doing the first outside the loop because we need to read
            // in the length OR ask the user. I prefer to just read it
            // in from the first file.

            file, err := os.Open(files[0])
            if err != nil {
                log.Fatal(err)
            }

            // get the length of the file
            fileInfo, err := file.Stat()
            if err != nil {
                log.Fatal(err)
            }

            // read the file into data
            data := make([]byte, fileInfo.Size())
            _, err = file.Read(data)
            if err != nil {
                log.Fatal(err)
            }

            var masterCounter StreamCounter
            err = json.Unmarshal(data, &masterCounter)
            if err != nil {
                log.Fatal(err)
            }

            // now we add any other files
            for i := 1; i < len(files); i++ {

                file, err := os.Open(files[1])
                if err != nil {
                    log.Fatal(err)
                }

                // get the length of the file
                fileInfo, err := file.Stat()
                if err != nil {
                    log.Fatal(err)
                }

                // read the file into data
                data := make([]byte, fileInfo.Size())
                _, err = file.Read(data)
                if err != nil {
                    log.Fatal(err)
                }

                var newCounter StreamCounter
                err = json.Unmarshal(data, &newCounter)
                if err != nil {
                    log.Fatal(err)
                }

                masterCounter.AddCounter(&newCounter)
            }

            // print out the data as either html charts or json
            page := NewChartPage(&masterCounter)

            switch *format {
            case "html":
                page.PrintHtml()
            case "json":
                page.PrintJson()
            }
        }
    })

	app.Run(os.Args)
}
