package main

import (
	"errors"
	"log"
)

type StreamCounter struct {
	Length      int
	Count       [][256]int64 // each position has a 256 length slice
	Probability [][256]float32
	Samples     int64
}

func NewStreamCounter(length int) *StreamCounter {
	// create a new instance
	sc := new(StreamCounter)

	// set the length
	sc.Length = length

	// initialize the counter
	sc.Count = make([][256]int64, sc.Length)

	// initialize the probabilities
	sc.Probability = make([][256]float32, sc.Length)

	// initialize the sample counter
	sc.Samples = 0

	return sc
}

func (sc *StreamCounter) AddBytes(bytes []byte) {

	// make sure we were given the right keystream length
	if len(bytes) != sc.Length {
		err := errors.New("AddBytes: the provided stream is not the right size")
		log.Fatal(err)
	}

	// increment our count of samples
	sc.Samples += 1

	// iterate over the stream, incrementing our counter and recalculating probabilities
	for position, value := range bytes {
		sc.Count[position][value] += 1

		sc.Probability[position][value] = float32(sc.Count[position][value]) / float32(sc.Samples)
	}
}
