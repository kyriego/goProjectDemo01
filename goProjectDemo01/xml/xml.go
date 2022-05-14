package main

import (
	"encoding/xml"
	"fmt"
)

/* type Address struct {
	City, State string
}
*/
/* type Person struct {
	XMLName   xml.Name `xml:"person"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int
	Height    float32 `xml:"height,omitempty"`
	Married   bool    `xml:"-"`
	Address
	Comment     string `xml:",comment"`
	Description string `xml:",innerxml"`
} */

type Person struct {
	Id        int    `xml:"id,attr"`
	FirstName string `xml:"name>firstname"`
	LastName  string `xml:"name>lastname"`
	Addr      string `xml:"addr"`
	Gender    bool   `xml:"gender"`
}

func main() {
	p := Person{
		Id:        1,
		FirstName: "Kyrie",
		LastName:  "Irving",
		Addr:      "USA",
		Gender:    true,
	}
	b, _ := xml.MarshalIndent(p, "", " ")
	p1 := &Person{}
	xml.Unmarshal(b, p1)
	fmt.Printf("%v\n", *p1)
}
