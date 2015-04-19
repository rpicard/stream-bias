package main

import (
    "crypto/rc4"
    "crypto/rand"
    "log"
)

type Rc4Streamer struct {
    *StreamCipher
}

func NewRc4Streamer() *Rc4Streamer {
    return &Rc4Streamer{Name: "RC4"}
}

// generate a random key and return the first numBytes of keystream
func (sc *Rc4Streamer) RandomKeyStream(numBytes int) {

    // hard-coding the key length for now, probably should be an option
    randomKey := make([]byte, 128)

    _, err := rand.Read(randomKey)

    if err != nil {
        log.Fatal(err)
    }

    cipher, err := rc4.NewCipher(randomKey)

    if err != nil {
        log.Fatal(err)
    }

    // make a buffer of zeroes to XOR against the keystream so we can
    // get the actual keystream
    zeroBuffer := make([]byte, numBytes)
    keyStream := make([]byte, numBytes)

    cipher.XORKeyStream(keyStream, zeroBuffer)

    return keyStream
}

