package api

import (
	"encoding/json"
	"fmt"
	"modules/internal/server/GetData/providers"
	helper2 "modules/internal/server/api/helper"
	"modules/internal/userModel"
	"net/http"
	"strings"
)

type Service struct {
	service *Api
}

type Api interface {
	GetProfileInfo(w http.ResponseWriter, r *http.Request)
}

func (service *Service) GetProfileInfo(w http.ResponseWriter, r *http.Request) {
	method := providers.Service{}
	helper := helper2.Service{}
	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	var data userModel.ProfileData
	var err error
	data, err, userModel.UserID = helper.QueryProfileData(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
	fmt.Println(userModel.UserID)

	method.XMLGet(userModel.UserID)

}
