package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "encoding/xml"
    "github.com/davecgh/go-spew/spew"
)

type Email struct {
    Where string `xml:"where,attr"`
    Addr  string
}
type Address struct {
    City, State string
}
type Result struct {
    XMLName xml.Name `xml:"Person"`
    Name    string   `xml:"FullName"`
    Phone   string
    Email   []Email
    Groups  []string `xml:"Group>Value"`
    Address
}

type Price struct {
    XMLName xml.Name `xml:"TotalFare"`

    Amount   string `xml:"Amount,attr"`
    Currency string `xml:"CurrencyCode,attr"`
}

type BfmItinerary struct {
    XMLName xml.Name `xml:"PricedItinerary"`

    Price   Price `xml:"AirItineraryPricingInfo>ItinTotalFare>TotalFare"`
}

type Response struct {
    XMLName     xml.Name `xml:"OTA_AirLowFareSearchRS"`

    Itineraries []BfmItinerary `xml:"PricedItineraries>PricedItinerary"`
}

var xmlFileName = flag.String("file", "bfm.xml", "Input file path")

func main() {
    flag.Parse()
    v := Response{}

    content, err := ioutil.ReadFile(*xmlFileName)
    if err != nil {
        fmt.Println("Error opening file: %v\n", err)
        return
    }

    err = xml.Unmarshal([]byte(content), &v)
    if err != nil {
        fmt.Printf("Error parsing file: %v\n", err)
        return
    }
    spew.Dump(v)
}
