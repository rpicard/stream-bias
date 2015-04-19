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

        <style>

.chart rect {
    fill: steelblue;
}

.chart text {
    fill: white;
    font: 10px sans-serif;
    text-anchor: end;
}
        </style>

    </head>

    <body>

        <script src="https://cdnjs.cloudflare.com/ajax/libs/d3/3.5.5/d3.min.js"></script>

        <svg class="chart"></svg>

        <script type="text/javascript">

var THE_DATA = {{ .JsonData }};

var width = 960,
    height = 500;

var barWidth = width / THE_DATA[0].length;

var y = d3.scale.linear()
    .range([height, 0]);

var chart = d3.select(".chart")
    .attr("width", width)
    .attr("height", height);

var bar = chart.selectAll("g")
    .data(THE_DATA[0])
    .enter().append("g")
    .attr("transform", function(d, i) { return "translate(" + i * barWidth + ",0)"; });

bar.append("rect")
    .attr("y", function(d) { return y(d); })
    .attr("height", function(d) { return height - y(d) })
    .attr("width", barWidth - 1)

bar.append("text")
    .attr("x", barWidth / 2)
    .attr("y", function(d) { return y(d) + 3; })
    .attr("dy", ".75em")
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

