package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID uint) (string, error) // Note: Changed int to uint for consistency
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct{}

var SECRET_KEY = []byte("User_Management")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID uint) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return SECRET_KEY, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
