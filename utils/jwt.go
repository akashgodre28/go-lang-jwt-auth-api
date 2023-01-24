package utils

import (
	"UserAuth/dto"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func GenerateJWT(username string) (*dto.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["username"] = username
	claims["authorized"] = true
	tokenString, err := token.SignedString(sampleSecretKey)
	CheckNilErr(err)
	tokenDto := dto.Token{Token: &tokenString, Type: "Bearer"}
	return &tokenDto, err
}

func VerifyJWT(tokenString string) (bool, jwt.MapClaims) {
	claims := *new(jwt.MapClaims)
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return sampleSecretKey, nil
	})
	CheckNilErr(err)
	return token.Valid, claims
}

func extractClaims(_ http.ResponseWriter, request *http.Request) (string, error) {
	if request.Header["Token"] != nil {
		tokenString := request.Header["Token"][0]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("there's an error with the signing method")
			}
			return sampleSecretKey, nil
		})
		if err != nil {
			return "Error Parsing Token: ", err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			username := claims["username"].(string)
			return username, nil
		}
	}

	return "unable to extract claims", nil
}
