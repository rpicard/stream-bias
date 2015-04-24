package main

import (
	"crypto/rand"
	"github.com/rpicard/sosemanuk"
	"log"
)

type SosemanukStreamer struct {
	Name string
}

func NewSosemanukStreamer() *SosemanukStreamer {
	return &SosemanukStreamer{Name: "Sosemanuk"}
}

// generate a random key and return the first numBytes of keystream
func (sc *SosemanukStreamer) RandomKeyStream(numBytes int) []byte {

	randomKey := make([]byte, 16)
	randomIv := make([]byte, 16)

	_, err := rand.Read(randomKey)
	if err != nil {
		log.Fatal(err)
	}

	_, err = rand.Read(randomIv)
	if err != nil {
		log.Fatal(err)
	}

    // initialze the cipher
    var cipher *sosemanuk.Sosemanuk
    cipher.Init(randomKey, randomIv)

	// make a buffer of zeroes to XOR against the keystream so we can
	// get the actual keystream
	zeroBuffer := make([]byte, numBytes)
	keyStream := make([]byte, numBytes)

	cipher.XORKeyStream(keyStream, zeroBuffer)

	return keyStream
}
