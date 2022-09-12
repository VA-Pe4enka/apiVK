package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"v2/structs"
)

type Service struct{}

type Server interface {
	Handler()
}

func (service *Service) Handler() {

	http.HandleFunc("/profileInfo/", GetProfileInfo)
	fmt.Println("Server is going on!")
	http.ListenAndServe(":8080", nil)
}

func loadApiConfig(filename string) (structs.ApiConfig, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return structs.ApiConfig{}, err
	}

	var c structs.ApiConfig

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return structs.ApiConfig{}, err
	}

	return c, nil
}

func queryProfileData(id string) (structs.ProfileData, error, int) {
	var userID int

	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return structs.ProfileData{}, err, userID
	}

	resp, err := http.Get("https://api.vk.com/method/users.get?user_ids=" + id + "&access_token=" + apiConfig.VkOpenApiToken + "&v=5.131")
	if err != nil {
		return structs.ProfileData{}, err, userID
	}

	defer resp.Body.Close()

	var d structs.ProfileData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return structs.ProfileData{}, err, userID
	}
	for i := 0; i < len(d.Response); i++ {
		fmt.Println(d.Response[i].ID)
		userID = d.Response[i].ID
	}

	return d, nil, userID
}

func GetProfileInfo(w http.ResponseWriter, r *http.Request) {
	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	data, err, userID := queryProfileData(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
	fmt.Println(userID)
}