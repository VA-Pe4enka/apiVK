package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type apiConfig struct {
	VkOpenApiToken string `json:"VkOpenApiToken"`
}
type ResponseList struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Closed    bool   `json:"is_closed"`
}

type profileData struct {
	Response []ResponseList `json:"response"`
}

func loadApiConfig(filename string) (apiConfig, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return apiConfig{}, err
	}

	var c apiConfig

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfig{}, err
	}

	return c, nil
}

func queryProfileData(id string) (profileData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return profileData{}, err
	}

	resp, err := http.Get("https://api.vk.com/method/users.get?user_ids=" + id + "&access_token=" + apiConfig.VkOpenApiToken + "&v=5.131")
	if err != nil {
		return profileData{}, err
	}

	defer resp.Body.Close()

	var d profileData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		//userID, _ := resp.Body.Read()
		return profileData{}, err
	}

	return d, nil
}

func getProfileInfo(w http.ResponseWriter, r *http.Request) {
	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	data, err := queryProfileData(id)
	userData, _ := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
	fmt.Println(string(userData))
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString(string(userData), -1))
	userID := re.FindAllString(string(userData), -1)
	userIdString := strings.Join(userID, "")
	fmt.Println(userIdString)
}

func main() {
	http.HandleFunc("/profileInfo/", getProfileInfo)

	http.ListenAndServe(":8080", nil)
}

//http://vk.com/foaf.php?id=1
