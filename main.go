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

    const chartPageHtml = `
<!DOCTYPE html>
<html lang="en">
    <head>

        <script src="https://cdnjs.cloudflare.com/ajax/libs/d3/3.5.5/d3.min.js"></script>

        <style>

.chart div {
    font: 10px sans-serif;
    background-color: steelblue;
    text-align: right;
    padding: 3px;
    margin: 1px;
    color: white;
}

        <style>

    </head>

    <body>

        <script type="text/javascript">

var THE_DATA = {{ .JsonData }};

// create one chart
var chart = d3.select("body")
    .append("div")
        .attr("class", "chart");

chart.selectAll("div")
    .data(THE_DATA[0])
    .enter().append("div")
        .style("width", function(d) { return d * 10 + "px"; })
        .text(function(d) { return d; });

        </script>
    </body>
</html>
`

    jsn, err := json.Marshal(sc.Count)

    if err != nil {
        log.Fatal(err)
    }

    tmpl := template.Must(template.New("bias-charts").Parse(chartPageHtml))

    _ = tmpl.Execute(os.Stdout, &ChartPage{JsonData: string(jsn)})
}

