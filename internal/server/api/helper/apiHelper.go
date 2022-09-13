package helper

import (
	"encoding/json"
	"fmt"
	"modules/configs/server"
	"modules/internal/userModel"
	"net/http"
)

type Service struct {
	service *Helper
}

type Helper interface {
	QueryProfileData(id string) (userModel.ProfileData, error, int)
}

func (service *Service) QueryProfileData(id string) (userModel.ProfileData, error, int) {
	method := server.Service{}
	var userID int

	apiConfig, err := method.LoadApiConfig("configs/server")
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
