package controller

import (
	"UserAuth/service"
	"encoding/json"
	"net/http"
)

type UserController struct {
}

var userService = service.UserService{}

func (UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userService.GetAllUsers())
}
