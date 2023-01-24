package middleware

import (
	"UserAuth/dto"
	"UserAuth/service"
	"UserAuth/utils"
	"encoding/json"
	"net/http"
	"strings"
)

func CheckAuthentication(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := "No Auth"
		headers := r.Header
		val, ok := headers["Authorization"]
		if ok {
			authorization = val[0]
			valid, claims := utils.VerifyJWT(strings.Split(authorization, " ")[1])
			if (service.UserService{}.CheckUserByUserName(claims["username"].(string))) && valid {
				w.Header().Set("username", claims["username"].(string))
				h.ServeHTTP(w, r)
			}
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(dto.Response{Message: "FORBIDDEN", Status: 403})
		}
	})
}
