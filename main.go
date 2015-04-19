package main

import (
    "log"
    "encoding/json"
)


func main() {
    sc := NewStreamCounter(5)
    sc.AddBytes([]byte{0x00, 0x01, 0x03, 0x05, 0x01})

    jsn, err := json.Marshal(sc.Count)

    if err != nil {
        log.Fatal(err)
    }

    page := NewChartPage(string(jsn))

    page.PrintHtml()
}

