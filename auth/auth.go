package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func GenerateJWT() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":        time.Date(2022, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"authorized": true,
		"user":       "username",
	})

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "[ERROR] Signing Error", err
	}

	return tokenString, nil
}

func VerifyJWT(_ http.ResponseWriter, request *http.Request) (string, error) {
	if request.Header["Authorization"] != nil {
		tokenString := request.Header["Authorization"][0]
		tokenString = strings.Split(tokenString, "Bearer ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("[ERROR] There's an error with the signing method")
			}
			return sampleSecretKey, nil
		})
		if err != nil {
			return "[ERROR] Error Parsing Token: ", err
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			username := claims["user"].(string)
			return username, nil
		} else {
			fmt.Println(err)
			return "", err
		}
	}

	return "[ERROR] Unable to extract claims.", nil
}

func AuthPage() string {
	token, _ := GenerateJWT()
	fmt.Println("\n Token: ", token)
	return token
}
