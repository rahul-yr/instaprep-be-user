package auth

import (
	"crypto/rand"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	default_token_expiry int = 3
)

type TokenGeneratorService struct{}

type Claims struct {
	Email  string `json:"email"`
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type TokenMapper struct {
	Email  string
	UserId string
	Token  string
}

// This method generates and returns the auth token
func (token *TokenGeneratorService) GenerateAuthToken(email string, userId string) (string, error) {
	jwt_secret_token := []byte(os.Getenv("JWT_TOKEN_GENERATOR_SECRET"))
	if len(jwt_secret_token) < 7 {
		return "", errors.New("token_generator_error : JWT Token is not valid")
	}
	// token validity 3 days
	expirationTime := time.Now().AddDate(0, 0, default_token_expiry)
	// add claims and payload
	claims := &Claims{
		Email:  email,
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	unsigned_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	signed_token, err := unsigned_token.SignedString(jwt_secret_token)
	if err != nil {
		return "", err
	}
	return signed_token, nil
}

// This method verify the authenticity of auth token and
//  returns the Claims object
func (token *TokenGeneratorService) VerifyAuthToken(given_token string) (*TokenMapper, error) {
	jwt_secret_token := []byte(os.Getenv("JWT_TOKEN_GENERATOR_SECRET"))
	if len(jwt_secret_token) < 7 {
		return nil, errors.New("token_verify_error : JWT Token is not valid")
	}
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(given_token, claims, func(jwt_token *jwt.Token) (interface{}, error) {
		return jwt_secret_token, nil
	})

	if err != nil {
		return nil, err
	}

	return &TokenMapper{Email: claims.Email, UserId: claims.UserId, Token: given_token}, nil
}

// This method generates and returns the new auth token
func (token *TokenGeneratorService) RefreshAuthToken(given_token string) (*TokenMapper, error) {
	jwt_secret_token := []byte(os.Getenv("JWT_TOKEN_GENERATOR_SECRET"))
	if len(jwt_secret_token) < 7 {
		return nil, errors.New("token_refresh_error : JWT Token is not valid")
	}
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(given_token, claims, func(jwt_token *jwt.Token) (interface{}, error) {
		return jwt_secret_token, nil
	})

	if err != nil {
		return nil, err
	}
	if time.Until(time.Unix(claims.ExpiresAt, 0)).Seconds() > 30 {
		return nil, errors.New("refresh_token : Sorry! You already have the valid token")
	}
	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().AddDate(0, 0, default_token_expiry)
	claims.ExpiresAt = expirationTime.Unix()
	unsigned_new_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed_new_token, err := unsigned_new_token.SignedString(jwt_secret_token)
	if err != nil {
		return nil, err
	}

	return &TokenMapper{Email: claims.Email, UserId: claims.UserId, Token: signed_new_token}, nil
}

// This method generates and returns the OTP token
func (token *TokenGeneratorService) GenerateOTPToken() (string, error) {
	otp_length := 6
	otp_chars := os.Getenv("OTP_ALLOWED_DIGITS")
	if len(otp_chars) != 10 {
		return "", errors.New("otp_generator_error : OTP allowed digits is not valid")
	}
	otp_chars_len := len(otp_chars)
	buffer := make([]byte, otp_length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	for i := 0; i < otp_length; i++ {
		buffer[i] = otp_chars[int(buffer[i])%otp_chars_len]
	}
	return string(buffer), nil
}
