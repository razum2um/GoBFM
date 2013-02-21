package main

import (
    "fmt"
    "flag"
    "time"
    "io/ioutil"
    "encoding/xml"
    "encoding/json"
    "github.com/davecgh/go-spew/spew"
)

// XML types

type Airport struct {
    Name string `xml:"LocationCode,attr"`
}

type Airline struct {
    Code string `xml:",attr"`
}

type Timezone struct {
    Offset  int `xml:"GMTOffset,attr"`
}

type Flight struct {
    XMLName             xml.Name  `xml:"FlightSegment"`
    StartDt             time.Time `xml:"DepartureDateTime,attr"`
    EndDt               time.Time `xml:"ArrivalDateTime,attr"`
    DestanationTimezone Timezone  `xml:"DepartureTimeZone"`
    OriginTimezone      Timezone  `xml:"ArrivalTimeZone"`
    ElapsedTime         int       `xml:"ElapsedTime,attr"`
    Eticket             struct {
        Ind      string    `xml:"Ind,attr"`
    }                             `xml:"TPA_Extensions>eTicket"`
    Cls                 string    `xml:"ResBookDesigCode,attr"`
    Number              int       `xml:"FlightNumber,attr"`
    Equipment           struct {
        Equip string `xml:"AirEquipType,attr"`
    }
    OperatingAirline    Airline   `xml:"OperatingAirline"`
    MarketingAirline    Airline   `xml:"MarketingAirline"`
    Origin              Airport   `xml:"DepartureAirport"`
    Destination         Airport   `xml:"ArrivalAirport"`
}

type Route struct {
    Flights []Flight `xml:"FlightSegment"`
}

type FlightInfo struct {
    XMLName         xml.Name    `xml:"AirItinerary"`
    DirectionType   string      `xml:"DirectionInd,attr"`
    Routes          []Route     `xml:"OriginDestinationOptions>OriginDestinationOption"`
}

type Price struct {
    XMLName xml.Name `xml:"TotalFare"`
    Amount   string `xml:"Amount,attr"`
    Currency string `xml:"CurrencyCode,attr"`
}

type PricingInfo struct {
    XMLName xml.Name `xml:"AirItineraryPricingInfo"`
    Price Price `xml:"ItinTotalFare>TotalFare"`
    LastTicketingDate string `xml:"LastTicketDate,attr"`
}

type BfmItinerary struct {
    XMLName xml.Name `xml:"PricedItinerary"`
    PricingInfo     PricingInfo
    FlightInfo      FlightInfo
}

type Response struct {
    XMLName     xml.Name `xml:"OTA_AirLowFareSearchRS"`
    Itineraries []BfmItinerary `xml:"PricedItineraries>PricedItinerary"`
}

// end of XML types

var xmlFileName = flag.String("file", "bfm.xml", "Input file path")

func main() {
    flag.Parse()
    v := Response{}

    content, err := ioutil.ReadFile(*xmlFileName)
    if err != nil {
        fmt.Println("Error opening file: %v\n", err)
        return
    }

    xerr := xml.Unmarshal([]byte(content), &v)
    if xerr != nil {
        fmt.Printf("Error parsing file: %v\n", err)
        return
    }
    spew.Dump(v)

    t := v.Itineraries[0].PricingInfo.Price.Amount
    fmt.Println("XMLName %v", t)
    json, jerr := json.Marshal(t)
    if jerr != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(json)
}
