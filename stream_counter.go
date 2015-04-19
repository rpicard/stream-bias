package main

import (
    "log"
    "errors"
)

// we want to get n bytes from a random keystream for the stream ciphers
type RandomKeyStreamer interface {
    RandomKeyStream(numBytes int)
}

type StreamCounter struct {
    Length    int
    Count   [][256]int64 // each position has a 256 length slice
}

func NewStreamCounter(length int) *StreamCounter {
    // create a new instance
    sc := new(StreamCounter)

    // set the length
    sc.Length = length

    // initialize the counter
    sc.Count = make([][256]int64, sc.Length)

    return sc
}

func (sc *StreamCounter) AddBytes(bytes []byte) {

    // make sure we were given the right keystream length
    if len(bytes) != sc.Length {
        err := errors.New("AddBytes: the provided stream is not the right size")
        log.Fatal(err)
    }

    // iterate over the stream, incrementing our counters
    for position, value := range bytes {
        sc.Count[position][value] += 1
    }
}
