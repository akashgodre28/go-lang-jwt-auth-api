package controller

import (
	"UserAuth/entities"
	"encoding/json"
	"net/http"
)

type AuthController struct {
}

func (AuthController) GenerateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	var user entities.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	token, err := userService.AuthenticateUser(&user)
	if err != nil {
		if err.Error() == "user not found" {
			w.WriteHeader(403)
			resp := make(map[string]interface{})
			resp["message"] = "User Not found !!"
			resp["status"] = 403
			json.NewEncoder(w).Encode(resp)
			return
		}
		if err.Error() == "invalid password" {
			w.WriteHeader(403)
			resp := make(map[string]interface{})
			resp["message"] = "Invalid username or password"
			resp["status"] = 403
			json.NewEncoder(w).Encode(resp)
			return
		}
	}
	json.NewEncoder(w).Encode(token)
}
