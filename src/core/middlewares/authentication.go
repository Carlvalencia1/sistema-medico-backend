package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(name, rol string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"name": name,
		"rol": rol,
		"exp": time.Now().Add(time.Hour * 24).Unix(), 
	})

	tokenString, err := token.SignedString(secretKey);

	if err != nil {
		return "", err; 
	}

	return tokenString, err; 
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err; 
	}

	if !token.Valid {
		return fmt.Errorf("token invalid!"); 
	}

	return nil;
}

func VerifyPassword(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func HashPassword(password string) (string,error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14); 

	if err != nil {
		return "", err; 
	}
	
	return string(bytes), err; 
}