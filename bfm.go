package main

import (
    "os"
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
    PricingInfo     PricingInfo     `xml:"AirItineraryPricingInfo"`
    FlightInfo      FlightInfo      `xml:"AirItinerary"`
}

// see: http://stackoverflow.com/questions/11126793/golang-json-and-dealing-with-unexported-fields
func (a Airport) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct{
        Name string `json:"name"`
    }{
        Name: a.Name,
    })
}

func (f Flight) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct{
        StartDt time.Time `json:"start_dt"`
        EndDt time.Time `json:"end_dt"`
        DestanationTimezone int `json:"destination_timezone"`
        OriginTimezone int `json:"origin_timezone"`
        ElapsedTime int `json:"elapsed_time"`
        Eticket bool `json:"eticket"`
        Cls string `json:"cls"`
        Number int `json:"number"`
        Equipment string `json:"equipment"`
        OperatingAirline string `json:"operating_airline"`
        MarketingAirline string `json:"marketing_airline"`
        Origin Airport `json:"origin"`
        Destination Airport `json:"destination"`
    }{
        StartDt: f.StartDt,
        EndDt: f.EndDt,
        DestanationTimezone: f.DestanationTimezone.Offset,
        OriginTimezone: f.OriginTimezone.Offset,
        ElapsedTime: f.ElapsedTime,
        Eticket: f.Eticket.Ind == "true",
        Cls: f.Cls,
        Number: f.Number,
        Equipment: f.Equipment.Equip,
        OperatingAirline: f.OperatingAirline.Code,
        MarketingAirline: f.MarketingAirline.Code,
        Origin: f.Origin,
        Destination: f.Destination,
    })
}

func (r Route) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct{
        Flights []Flight `json:"flights"`
    }{
        Flights: r.Flights,
    })
}

func (i BfmItinerary) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct{
        Price string `json:"price"`
        LastTicketingDate string `json:"last_tiketing_date"`
        DirectionType string `json:"direction_type"`
        Routes []Route `json:"routes"`
    }{
        Price: i.PricingInfo.Price.Amount,
        LastTicketingDate: i.PricingInfo.LastTicketingDate,
        DirectionType: i.FlightInfo.DirectionType,
        Routes: i.FlightInfo.Routes,
    })
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
    spew.Dump(v.Itineraries[0])

    json, jerr := json.Marshal(v.Itineraries)
    if jerr != nil {
        fmt.Println("error:", err)
    }
    os.Stdout.Write(json)
}
