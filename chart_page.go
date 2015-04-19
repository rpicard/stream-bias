package main

import (
    "text/template"
    "log"
    "os"
)

type ChartPage struct {
    JsonData    string
}

func NewChartPage(jsn string) *ChartPage {
    return &ChartPage{JsonData: jsn}
}

// print the generated HTML for this page
func (page *ChartPage) PrintHtml() {

    tmpl := template.Must(template.New("bias-charts").Parse(BiasChartsTemplate))

    err := tmpl.Execute(os.Stdout, page)

    if err != nil {
        log.Fatal(err)
    }
}

const BiasChartsTemplate = `
<!DOCTYPE html>
<html lang="en">
    <head>
    <!-- lots of credit for the d3 stuff goes to http://bl.ocks.org/benjchristensen/2580640 -->
        <style>

body {
    font-family: "Helvetica Neue", Helvetica;
}

/* tell the SVG path to be a thin blue line without any area fill */
path {
    stroke-width: 1;
    fill: none;
}

.data {
    stroke: green;
}

.axis {
  shape-rendering: crispEdges;
}

.x.axis line {
  stroke: lightgrey;
}

.x.axis .minor {
  stroke-opacity: .5;
}

.x.axis path {
  display: none;
}

.x.axis text {
    font-size: 10px;
}

.y.axis line, .y.axis path {
  fill: none;
  stroke: #000;
}

.y.axis text {
    font-size: 12px;
}

        </style>

    </head>

    <body>

        <script src="https://cdnjs.cloudflare.com/ajax/libs/d3/3.5.5/d3.min.js"></script>

        <script type="text/javascript">

var DATA = {{ .JsonData }};

for (var i = 0; i < DATA.length; i++) {

    var m = [80, 80, 80, 80]; // margins
    var w = 1000 - m[1] - m[3]; // width
    var h = 400 - m[0] - m[2]; // height

    var x = d3.scale.linear()
        .domain([0, DATA[i].length])
        .range([0, w]);

    var y = d3.scale.linear()
        .domain([0, d3.max(DATA[i])])
        .range([h, 0]);

    var line = d3.svg.line()
        .x(function(d,i) {
            return x(i);
        })
        .y(function(d) {
            return y(d);
        });

    var graph = d3.select("body")
        .append("div")
        .attr("class", "graph" + i)

    d3.select(".graph" + i)
        .insert("h2")
        .text("Position " + (i + 1))

    graph = graph.append("svg:svg")
        .attr("width", w + m[1] + m[3])
        .attr("height", h + m[0] + m[2])
        .append("svg:g")
        .attr("transform", "translate(" + m[3] + "," + m[0] + ")");

    var xAxis = d3.svg.axis()
        .scale(x)
        .tickSize(-h)
        .tickSubdivide(1);

    graph.append("svg:g")
        .attr("class", "x axis")
        .attr("transform", "translate(0," + h + ")")
        .call(xAxis);

    var yAxis = d3.svg.axis()
        .scale(y)
        .ticks(6)
        .orient("left");

    graph.append("svg:g")
        .attr("class", "y axis")
        .attr("transform", "translate(-10,0)")
        .call(yAxis);

    graph.append("svg:path")
        .attr("d", line(DATA[i]))
        .attr("class", "data");

}

        </script>

    </body>
</html>
`
