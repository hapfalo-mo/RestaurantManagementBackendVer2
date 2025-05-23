package service

import (
	dto "RestuarantBackend/models/dto"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("a-string-secret-at-least-256-bits-long")

// Create Struct Claims
type Claims struct {
	UserID   int    `json:"userID"`
	Role     int    `json:"role"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	FullName string `json:"fullName"`
	Point    int    `json:"point"`
	jwt.RegisteredClaims
}

// Create Token
func CreateToken(user *dto.LoginResponse) (string, error) {
	claims := &Claims{
		UserID:   user.Id,
		Email:    user.Email,
		Phone:    user.PhoneNumber,
		FullName: user.FullName,
		Role:     user.Role,
		Point:    user.Point,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "RestuarantBackend",
			Subject:   "Authentication",
		},
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, errors.New("Token Expired")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func CallApiCheckUser(token string) (bool, error) {
	url := "http://localhost:1703/api/v1/users/verify-user-access"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("Failed in veriy user!")
	}
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return false, err
	// }
	// var loginResponse dto.LoginResponse
	// err = json.Unmarshal(body, loginResponse)
	// if err != nil {
	// 	return false, err
	// }
	return true, nil
}
