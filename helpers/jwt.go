package helpers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
	"time"
)

func GenerateToken(id uint, email string) (res string) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(os.Getenv("INTERNAL_SECRET")))
	if err != nil {
		return ""
	}

	res = signedString
	return
}

func VerifyToken(c *gin.Context) (result interface{}, err error) {
	authHeader := c.Request.Header.Get("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer") || authHeader == "" {
		log.Println("[check Authorization]")
		err = errors.New("invalid Authorization")
		return
	}
	token := strings.Split(authHeader, " ")[1]

	tokenResult, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("[token.Method.(*jwt.SigningMethodHMAC)]")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("INTERNAL_SECRET")), nil
	})
	if err != nil {
		log.Println("[jwt.Parse]")
		return nil, err
	}

	if _, ok := tokenResult.Claims.(jwt.MapClaims); !ok && !tokenResult.Valid {
		log.Println("[tokenResult.Claims]")
		err = errors.New("invalid Authorization")
		return
	}

	result = tokenResult.Claims.(jwt.MapClaims)
	return
}
