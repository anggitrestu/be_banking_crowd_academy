package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretJWT string
}

func NewService(secretJWT string) *jwtService {
	return &jwtService{secretJWT}
}

func (s *jwtService) GenerateToken(ID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = ID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(s.secretJWT))
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

		return []byte(s.secretJWT), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
