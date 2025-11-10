package auth

import (
	"errors"
	"fmt"
	"house-store/internal/consts"
	"house-store/internal/enum"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ilyakaznacheev/cleanenv"
)

type configJWT struct {
	TokenSignSecret string `env:"TOKEN_SIGN_SECRET" env-required:"true"`
}

var (
	config          configJWT
	mySigningMethod = jwt.SigningMethodHS512
)

func LoadJWTSecret() error {
	path := fmt.Sprintf("%v.env", consts.PathToEnvFile)
	err := cleanenv.ReadConfig(path, &config)
	if err != nil {
		return fmt.Errorf("LoadTokenSignSecret. Cannot read configJWT. err: %q", err.Error())
	}

	return nil
}

func CreateToken(userRole string) (string, error) {
	return createTokenWithKey(userRole, config.TokenSignSecret)
}

func createTokenWithKey(userRole, secretKey string) (string, error) {
	switch userRole {
	case enum.USER_ROLE_CLIENT, enum.USER_ROLE_MODERATOR:
		//
	default:
		return "", fmt.Errorf("unknown userRole is send to JWT. userRole = %v", userRole)
	}

	if secretKey == "" {
		return "", errors.New("createTokenWithKey. secretKey wasn't provided")
	}

	currentTime := time.Now()
	expiresAt := currentTime.Add(time.Duration(90) * time.Minute)
	claims := jwt.RegisteredClaims{
		Subject:   userRole,
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	token := jwt.NewWithClaims(
		mySigningMethod,
		claims,
	)

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string) (string, error) {
	return validateTokenWithKey(tokenString, config.TokenSignSecret)
}

func validateTokenWithKey(tokenString, secretKey string) (string, error) {
	if secretKey == "" {
		return "", errors.New("validateTokenWithKey. secretKey wasn't provided")
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (any, error) {
			return []byte(secretKey), nil
		},
		jwt.WithValidMethods([]string{mySigningMethod.Alg()}),
	)
	if err != nil {
		return "", err
	}

	fmt.Printf("token.Claims = %+v", token.Claims)
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", fmt.Errorf("cannot convert token.Claims to 'jwt.RegisteredClaims' type")
	}

	userRole := claims.Subject
	switch userRole {
	case enum.USER_ROLE_CLIENT, enum.USER_ROLE_MODERATOR:
		return userRole, nil
	default:
		return userRole, fmt.Errorf("JWT contains unknown userRole = %v", userRole)
	}
}
