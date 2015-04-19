package main

import (
    "os"
    "log"
    "encoding/json"
    "text/template"
)

type ChartPage struct {
    JsonData    string
}

func main() {
    sc := NewStreamCounter(5)
    sc.AddBytes([]byte{0x00, 0x01, 0x03, 0x05, 0x01})

    jsn, err := json.Marshal(sc.Count)

    if err != nil {
        log.Fatal(err)
    }

    tmpl, err := template.ParseFiles("templates/bias-charts.html")

    if err != nil {
        log.Fatal(err)
    }

    _ = tmpl.Execute(os.Stdout, &ChartPage{JsonData: string(jsn)})
}
