package main

import (
    "fmt"
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
    //Body    string `xml:",innerxml"`
    //Data    string `xml:",chardata"`
    //Any     string `xml:",any"`

    Amount   string `xml:"Amount,attr"`
    Currency string `xml:"CurrencyCode,attr"`
}

type BfmItinerary struct {
    XMLName xml.Name `xml:"PricedItinerary"`
    //Body    string `xml:",innerxml"`
    //Data    string `xml:",chardata"`
    //Any     string `xml:",any"`

    Price Price `xml:"AirItineraryPricingInfo>ItinTotalFare>TotalFare"`
}

func main() {
    v := BfmItinerary{}

    data := `
        <PricedItinerary SequenceNumber="1" MultipleTickets="false">
          <AirItinerary DirectionInd="Return">
            <OriginDestinationOptions>
               <FlightSegment></FlightSegment>
               <FlightSegment></FlightSegment>
            </OriginDestinationOptions>
            <OriginDestinationOptions>
               <FlightSegment></FlightSegment>
               <FlightSegment></FlightSegment>
            </OriginDestinationOptions>
          </AirItinerary>
          <AirItineraryPricingInfo>
            <ItinTotalFare>
              <BaseFare Amount="2.00" CurrencyCode="EUR"/>
              <EquivFare Amount="85" CurrencyCode="RUB"/>
              <Taxes>
                <Tax TaxCode="TOTALTAX" Amount="8164" CurrencyCode="RUB"/>
              </Taxes>
              <TotalFare Amount="8249" CurrencyCode="RUB"/>
            </ItinTotalFare>
          </AirItineraryPricingInfo>
        </PricedItinerary>

        <Person>
            <FullName>Grace R. Emlin</FullName>
            <Company>Example Inc.</Company>
            <Email where="home">
                <Addr>gre@example.com</Addr>
            </Email>
            <Email where='work'>
                <Addr>gre@work.com</Addr>
            </Email>
            <Group>
                <Value>Friends</Value>
                <Value>Squash</Value>
            </Group>
            <City>Hanga Roa</City>
            <State>Easter Island</State>
        </Person>
    `
    err := xml.Unmarshal([]byte(data), &v)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    spew.Dump(v)
}
