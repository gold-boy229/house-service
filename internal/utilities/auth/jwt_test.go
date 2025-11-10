package auth

import (
	"house-store/internal/enum"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	USER_ROLE_UNKNOWN = "unknown"
)

func TestCreateAndValidateToken(t *testing.T) {
	tests := []struct {
		name         string
		userRole     string
		secretKey    string
		expectError  bool
		expectedRole string
	}{
		{
			name:         "Valid Client Role",
			userRole:     enum.USER_ROLE_CLIENT,
			secretKey:    "test_secret_key_very_secure",
			expectError:  false,
			expectedRole: enum.USER_ROLE_CLIENT,
		},
		{
			name:         "Valid Moderator Role",
			userRole:     enum.USER_ROLE_MODERATOR,
			secretKey:    "test_secret_key_very_secure",
			expectError:  false,
			expectedRole: enum.USER_ROLE_MODERATOR,
		},
		{
			name:         "Unknown Role",
			userRole:     USER_ROLE_UNKNOWN,
			secretKey:    "test_secret_key_very_secure",
			expectError:  true, // validateTokenWithKey should return an error for unknown roles
			expectedRole: USER_ROLE_UNKNOWN,
		},
		{
			name:         "Empty Secret Key for Creation",
			userRole:     enum.USER_ROLE_CLIENT,
			secretKey:    "",
			expectError:  true,
			expectedRole: "", // expect error on token creation
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test token creation
			tokenString, err := createTokenWithKey(tt.userRole, tt.secretKey)
			if tt.expectError && err == nil {
				t.Fatal("expected error during token creation, but got none")
			}
			if !tt.expectError && err != nil {
				t.Fatalf("did not expect error during token creation: %v", err)
			}
			if tt.expectError {
				return // Skip validation if creation failed as expected
			}

			// Test token validation
			validatedRole, err := validateTokenWithKey(tokenString, tt.secretKey)
			if err != nil {
				t.Fatalf("did not expect error during token validation: %v", err)
			}

			if validatedRole != tt.expectedRole {
				t.Errorf("expected role %q, but got %q", tt.expectedRole, validatedRole)
			}
		})
	}
}

func TestValidateToken_InvalidScenarios(t *testing.T) {
	secretKey := "super_secret_key"
	validRole := enum.USER_ROLE_CLIENT

	t.Run("Expired Token", func(t *testing.T) {
		// Create a token that expired 1 hour ago
		tokenString, err := createTokenWithCustomExpiration(validRole, secretKey, -1*time.Hour)
		if err != nil {
			t.Fatal(err)
		}

		_, err = validateTokenWithKey(tokenString, secretKey)
		if err == nil {
			t.Error("expected an error for expired token, but got none")
		}
	})

	t.Run("Incorrect Secret Key", func(t *testing.T) {
		tokenString, err := createTokenWithKey(validRole, secretKey)
		if err != nil {
			t.Fatal(err)
		}

		wrongKey := "wrong_secret_key"
		_, err = validateTokenWithKey(tokenString, wrongKey)
		if err == nil {
			t.Error("expected an error for incorrect secret key, but got none")
		}
	})

	t.Run("Token signed with different algorithm (HS256 vs HS512)", func(t *testing.T) {
		claims := jwt.RegisteredClaims{
			Subject:   validRole,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		}
		// Manually create with HS256
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			t.Fatal(err)
		}

		_, err = validateTokenWithKey(tokenString, secretKey)
		if err == nil {
			t.Error("expected an error for wrong signing method, but got none")
		}
	})

	t.Run("Token signed with different algorithm (HS384 vs HS512)", func(t *testing.T) {
		claims := jwt.RegisteredClaims{
			Subject:   validRole,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		}
		// Manually create with HS256
		token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			t.Fatal(err)
		}

		_, err = validateTokenWithKey(tokenString, secretKey)
		if err == nil {
			t.Error("expected an error for wrong signing method, but got none")
		}
	})
}

// Helper function to create a token with a specific expiration time for testing
func createTokenWithCustomExpiration(userRole, secretKey string, expiration time.Duration) (string, error) {
	currentTime := time.Now()
	expiresAt := currentTime.Add(expiration)
	claims := jwt.RegisteredClaims{
		Subject:   userRole,
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		claims,
	)

	return token.SignedString([]byte(secretKey))
}
