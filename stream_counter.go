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
    Count   [](map[byte]int64) // each position has a map of bytes => count
}

func NewStreamCounter(length int) *StreamCounter {
    // create a new instance
    sc := new(StreamCounter)

    // set the length
    sc.Length = length

    // initialize the counter
    count := make([](map[byte]int64), sc.Length)

    // for each position, make a positionMap
    for i := 0; i < len(count); i++ {

        // make a map of 256 keys to the initial count of 0
        positionMap := make(map[byte]int64, 256)

        for j := 0; j < 256; j++ {
            positionMap[byte(j)] = 0
        }

        count[i] = positionMap
    }


    sc.Count = count

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
