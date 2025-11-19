package auth

import (
	"errors"
	"fmt"
	"house-store/internal/entity"
	"house-store/internal/enum"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthorizedJWTClaims struct {
	UserUUID string `json:"user_uuid"`
	jwt.RegisteredClaims
}

func CreateAuthorizedToken(params entity.AuthorizedTokenParams) (string, error) {
	return createAuthorizedTokenWithKey(params, config.TokenSignSecret)
}

func createAuthorizedTokenWithKey(params entity.AuthorizedTokenParams, secretKey string) (string, error) {
	currentTime := time.Now()
	expiresAt := currentTime.Add(time.Duration(90) * time.Minute)
	claims := AuthorizedJWTClaims{
		UserUUID: params.UserUUID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   params.UserRole,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(
		mySigningMethod,
		claims,
	)

	return token.SignedString([]byte(secretKey))
}

func ValidateAuthorizedToken(tokenString string) (entity.AuthorizedTokenParams, error) {
	return validateAuthorizedTokenWithKey(tokenString, config.TokenSignSecret)
}

func validateAuthorizedTokenWithKey(tokenString, secretKey string) (entity.AuthorizedTokenParams, error) {
	if secretKey == "" {
		return entity.AuthorizedTokenParams{}, errors.New("validateTokenWithKey. secretKey wasn't provided")
	}

	var myClaims AuthorizedJWTClaims
	token, err := jwt.ParseWithClaims(
		tokenString,
		&myClaims,
		func(token *jwt.Token) (any, error) {
			return []byte(secretKey), nil
		},
		jwt.WithValidMethods([]string{mySigningMethod.Alg()}),
	)
	if err != nil {
		return entity.AuthorizedTokenParams{}, err
	}

	fmt.Printf("AuthorizedToken.   token.Claims = %+v\n", token.Claims)
	claims, ok := token.Claims.(*AuthorizedJWTClaims)
	if !ok {
		return entity.AuthorizedTokenParams{}, fmt.Errorf("cannot convert token.Claims to 'AuthorizedJWTClaims' type")
	}

	paramsFromToken := entity.AuthorizedTokenParams{
		UserRole: claims.Subject,
		UserUUID: claims.UserUUID,
	}

	switch paramsFromToken.UserRole {
	case enum.USER_ROLE_CLIENT, enum.USER_ROLE_MODERATOR:
		return paramsFromToken, nil
	default:
		return paramsFromToken, fmt.Errorf("JWT contains unknown userRole = %v", paramsFromToken.UserRole)
	}
}
