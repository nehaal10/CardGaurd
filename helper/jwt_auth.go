package helper

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nehaal10/CardGaurd/models"
	"github.com/nehaal10/CardGaurd/utils"
)

type JwtClaim struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func JwtAuth(user models.UserDatabase) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaim{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		},
	})
	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	utils.CheckError(err)

	return token
}

func GetCookie(c *gin.Context) (string, error) {
	cook, err := c.Request.Cookie("user")
	if err != nil {
		return "", err
	}
	token, err := jwt.ParseWithClaims(cook.Value, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	utils.CheckError(err)
	claim := token.Claims.(*JwtClaim)
	username := claim.Username
	return username, err
}
