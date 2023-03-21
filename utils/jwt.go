package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Claims struct {
	Id      uint   `json:"id"`
	Purpose string `json:"purpose"`
	jwt.StandardClaims
}

func GetJwtSecret() ([]byte, error) {
	JWTRawSecret := os.Getenv("JWT_SECRET")
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return []byte{}, err
	}

	return []byte(JWTRawSecret), nil
}

func GenerateAuthJwt(id uint) (string, error) {
	JwtSecret, err := GetJwtSecret()
	if err != nil {
		log.Println("Error loading .env file")
		return "", err
	}

	expirationTime := time.Now().Add(10 * 12 * 30 * 24 * time.Hour) //10 years
	claims := &Claims{
		Id:      id,
		Purpose: Authentication,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
func GenerateTimedAuthJwt(id uint, hours uint) (string, error) {
	JwtSecret, err := GetJwtSecret()
	if err != nil {
		log.Println("Error loading .env file")
		return "", err
	}

	expirationTime := time.Now().Add(time.Duration(hours) * time.Hour)
	claims := &Claims{
		Id:      id,
		Purpose: FileOperations,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
func GenerateUserVerificationJwt(id uint) (string, error) {
	JwtSecret, err := GetJwtSecret()
	if err != nil {
		log.Fatal("Error loading .env file")
		return "", err
	}

	expirationTime := time.Now().Add(12 * 30 * 24 * time.Hour)
	claims := &Claims{
		Id:      id,
		Purpose: UserVerification,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func GenerateResetPasswordJwt(id uint) (string, error) {
	JwtSecret, err := GetJwtSecret()
	if err != nil {
		log.Fatal("Error loading .env file")
		return "", err
	}

	expirationTime := time.Now().Add(3 * 24 * time.Hour)
	claims := &Claims{
		Id:      id,
		Purpose: ResetPassword,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ParseJwt(tokenString string) (Claims, error) {
	JwtSecret, err := GetJwtSecret()
	if err != nil {
		return Claims{}, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})
	if err != nil {
		return Claims{}, err
	}

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id, ok1 := claims["id"].(float64)
			purpose, ok2 := claims["purpose"].(string)

			if ok1 && ok2 {
				return Claims{
					Id:      uint(id),
					Purpose: purpose,
				}, nil
			} else {
				return Claims{}, errors.New("can not parse the token")
			}
		} else {
			return Claims{}, errors.New("can not parse the token")
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return Claims{}, errors.New("malformed token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {

			return Claims{}, errors.New("token has been expired")

		} else {
			return Claims{}, err

		}
	} else {
		return Claims{}, errors.New("unknown error")
	}

}

const (
	Authentication   = "AUTHENTICATION"
	UserVerification = "USER_VERIFICATION"
	ResetPassword    = "RESET_PASSWORD"
	FileOperations   = "FILE_OPERATIONS"
)
