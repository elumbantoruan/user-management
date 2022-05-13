package crypto

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(email string) (string, error) {
	signKey := []byte("mysupersecretkey")
	claims := &jwt.MapClaims{
		"email":      email,
		"expiration": time.Now().Add(1 * time.Hour).UTC().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signKey)
	return ss, err
}

func ValidateEmailInToken(val string) (string, error) {
	tok, _ := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if tok == nil {
		err := errors.New("invalid token")
		return "", err
	}

	var email string

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("unable to find claims in the token")
	}
	email = claims["email"].(string)
	if len(email) == 0 {
		return "", errors.New("unable to find email in the claims")
	}

	// check for expiration
	exp, ok := claims["expiration"]
	if !ok {
		return "", errors.New("unable to verify expiration of the token")
	}

	ts := int64(exp.(float64))
	if time.Now().UTC().Unix() > ts {
		return "", errors.New("token has expired")
	}

	return email, nil
}

type AuthResponse struct {
	Token string `json:"token"`
}

func CreateResponseWithToken(email string) (*AuthResponse, error) {
	token, err := CreateToken(email)
	if err != nil {
		return nil, err
	}
	return &AuthResponse{
		Token: token,
	}, nil
}

func Hash(password string) string {
	var salt = "!pwd:"
	saltedPassword := []byte(salt + password)
	return fmt.Sprintf("%x", sha256.Sum256(saltedPassword))
}
