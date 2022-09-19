package server

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"modules/configs/server"
	"modules/internal/server/api"
	"modules/internal/userModel"
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
	service *Server
}

type Server interface {
	Handler()
}

func (service *Service) Handler() {
	method := api.Service{}

	http.HandleFunc("/profileInfo/", method.GetProfileInfo)
	fmt.Println("Server is going on!")
	http.ListenAndServe(":8080", nil)
}

var UserID int

func queryProfileData(id string) (userModel.ProfileData, error, int) {
	service := server.Service{}
	var userID int

	apiConfig, err := service.LoadApiConfig("configs/server")
	if err != nil {
		return userModel.ProfileData{}, err, userID
	}

	resp, err := http.Get("https://api.vk.com/method/users.get?user_ids=" + id + "&access_token=" + apiConfig.VKToken + "&v=5.131")
	if err != nil {
		return userModel.ProfileData{}, err, userID
	}

	defer resp.Body.Close()

	var d userModel.ProfileData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return userModel.ProfileData{}, err, userID
	}
	for i := 0; i < len(d.Response); i++ {
		fmt.Println(d.Response[i].ID)
		userID = d.Response[i].ID
	}

	return d, nil, userID
}
