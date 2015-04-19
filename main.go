package main

import (
    "fmt"
)

func main() {
    sc := NewStreamCounter(5)
    sc.AddBytes([]byte{0x00, 0x01, 0x03, 0x05, 0x01})

    for position, positionMap := range sc.Count {

        fmt.Printf("Position %d\n\n", position)

        for i := 0; i < 256; i++ {

            fmt.Printf("%x", i)

            for j := 0; int64(j) < positionMap[byte(i)]; j++ {
                fmt.Printf(" ")
            }

            fmt.Printf("|\n")

        }


    }

}
