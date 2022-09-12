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
	XMLName xml.Name `xml:"xml_name"`
}

type Person struct {
	Name string `xml:"name"`
}
