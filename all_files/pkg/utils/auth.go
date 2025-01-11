package utils

import (
	// "errors"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)
var jwtKey []byte

func init() {
	// Initialize the JWT secret key globally
	err:=godotenv.Load()
	if err!=nil{
		log.Fatal("Error loading env")
	}
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("Error: JWT_SECRET_KEY environment variable is not set")
	}
	jwtKey = []byte(secret)
}

type Claims struct{
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

func GenerateToken(userId string)(string,error){
	expiration_time:=time.Now().Add(24*time.Hour)
	payload := &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration_time), 
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenstr string)(string,error){
	claims:=&Claims{}
	token,err:=jwt.ParseWithClaims(tokenstr,claims,func(t *jwt.Token) (interface{}, error) {return jwtKey,nil})
	if err!=nil{
		if err==jwt.ErrSignatureInvalid{
			return "",errors.New("invalid token signature")
		}
		return "",errors.New("failed to parse token")
	}
	if !token.Valid{
		return "",errors.New("invalid token")
	}
	

	return claims.UserID,nil
}
