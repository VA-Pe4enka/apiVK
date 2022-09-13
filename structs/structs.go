package structs

import "encoding/xml"

type ApiConfig struct {
	VkOpenApiToken string `json:"VkOpenApiToken"`
}
type ResponseList struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Closed    bool   `json:"is_closed"`
}

type ProfileData struct {
	Response []ResponseList `json:"response"`
}

type Result struct {
	XMLName xml.Name `xml:"RDF"`
	Person  Person   `xml:"Person>created"`
}

type Person struct {
	XMLName xml.Name `xml:"created" json:"-"`
	Data    string   `xml:"date,attr" json:"date"`
}

var UserID int
