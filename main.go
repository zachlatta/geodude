package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/kellydunn/golang-geo"
)

type geocodeResult struct {
	Address string
	Point   *geo.Point
}

var (
	prettyPrint bool
)

func main() {
	flag.BoolVar(&prettyPrint, "p", false, "pretty print")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	query := strings.Join(args, " ")

	result, err := geocode(query, new(geo.GoogleGeocoder))
	if err != nil {
		printErr(err)
	}

	if prettyPrint {
		tmpl(os.Stdout, prettyResultTmpl, result)
	} else {
		tmpl(os.Stdout, resultTmpl, result)
	}
}

const usageTmpl = `Geodude is a tiny command-line utility for geocoding addresses.

Usage:

  geodude [OPTION] [address]

Options:

  -p
        pretty print output

`

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

func printUsage(w io.Writer) {
	tmpl(w, usageTmpl, nil)
}

func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func printErr(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(2)
}

const resultTmpl = "{{.Point.Lat}}\t{{.Point.Lng}}\t{{.Address}}\n"

const prettyResultTmpl = `Address: {{.Address}}
Coordinates: {{.Point.Lat}}, {{.Point.Lng}}

`

func geocode(query string, geocoder geo.Geocoder) (*geocodeResult, error) {
	point, err := geocoder.Geocode(query)
	if err != nil {
		return nil, err
	}

	addr, err := geocoder.ReverseGeocode(point)
	if err != nil {
		return nil, err
	}

	result := &geocodeResult{
		Address: addr,
		Point:   point,
	}

	return result, nil
}
