package auth_token

import (
	app_config "api-social-media/app/core/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func Createtoken(userID uuid.UUID) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 24).Unix() // expire token
	permissions["userID"] = userID

	//asign using key secret to create and authentify our token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(app_config.SECRET_KEY)) //secret

}
