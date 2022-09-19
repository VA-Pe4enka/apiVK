package providers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"log"
	"net/http"
)

type Result struct {
	XMLName xml.Name `xml:"RDF"`
	Person  Person   `xml:"Person>created"`
}

type Person struct {
	XMLName xml.Name `xml:"created" json:"-"`
	Data    string   `xml:"date,attr" json:"date"`
}

type Service struct {
	service *Provider
}

type Provider interface {
	XMLGet(id int)
}

func (service *Service) XMLGet(id int) {
	data := &Result{}
	resp, err := http.Get("https://vk.com/foaf.php?id=" + string(id))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("resp.body: ", resp.Body)
	decoder := xml.NewDecoder(resp.Body)
	fmt.Println("decoder: ", decoder)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("data: ", err)

	result, err := json.Marshal(&data.XMLName)
	if err != nil {
		log.Fatal(err)
	}

	var dict []map[string]string
	if err = json.Unmarshal(result, &dict); err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println("dict: ", dict)

}
