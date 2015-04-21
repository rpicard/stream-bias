package main

import (
	"errors"
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
	cipher := app.StringOpt("c cipher", "", "which cipher are we testing?")
	samples := app.StringOpt("s samples", "1000", "how many samples should we take?")
	length := app.StringOpt("l length", "256", "how many positions in the keystream do we care about?")

	app.Spec = "-c [-f] [-s] [-l]"

	app.Action = func() {

		var streamer RandomKeyStreamer

		// add new ciphers here
		switch *cipher {
		case "rc4":
			streamer = NewRc4Streamer()
		case "hc128":
			streamer = NewHc128Streamer()
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

		sc := NewStreamCounter(streamLength)

		var wg sync.WaitGroup

		sampleStream := make(chan []byte, 1000)

		// generate a goroutine for each CPU
		for i := 0; i < runtime.NumCPU(); i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				// loop through our share of the samples
				for j := int64(0); j <= (numSamples / int64(runtime.NumCPU())); j++ {
					sampleStream <- streamer.RandomKeyStream(streamLength)
				}
			}()
		}

		// read our samples from the channel
		for i := int64(0); i < numSamples; i++ {
			sc.AddBytes(<-sampleStream)
		}

		// wait for the goroutines to generate all of their samples
		wg.Wait()
		close(sampleStream)

		// print out the data as either html charts or json
		page := NewChartPage(sc)

		switch *format {
		case "html":
			page.PrintHtml()
		case "json":
			page.PrintJson()
		}
	}

	app.Run(os.Args)
}
