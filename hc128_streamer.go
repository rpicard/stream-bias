package main

import (
    "github.com/tomfitzhenry/go-hc128"
    "crypto/rand"
    "encoding/binary"
    "log"
)

type Hc128Streamer struct {
    Name    string
}

func NewHc128Streamer() *Hc128Streamer {
    return &Hc128Streamer{Name: "HC-128"}
}

// generate a random key and return the first numBytes of keystream
func (sc *Hc128Streamer) RandomKeyStream(numBytes int) []byte {

    if (numBytes % 4) != 0 {
        log.Fatal("the hc-128 package doesn't implement padding. numBytes has to be a multiple of 4")
    }

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

    var keyParts [4]uint32
    for i := 0; i < len(randomKey)/4; i++ {
        keyParts[i] = binary.BigEndian.Uint32(randomKey[4*i : 4*(i+1)])
    }

    var ivParts [4]uint32
    for i := 0; i < len(randomIv)/4; i++ {
        ivParts[i] = binary.BigEndian.Uint32(randomIv[4*i : 4*(i+1)])
    }

    cipher := hc128.NewCipher(keyParts, ivParts)

    // make a buffer of zeroes to XOR against the keystream so we can
    // get the actual keystream
    zeroBuffer := make([]byte, numBytes)
    keyStream := make([]byte, numBytes)

    cipher.XORKeyStream(keyStream, zeroBuffer)

    return keyStream
}

