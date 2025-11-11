package handlers

import (
	"house-store/internal/dto"
	"house-store/internal/enum"
	"testing"
)

func TestValidateInputData(t *testing.T) {
	tests := []struct {
		Name        string
		UserRole    string
		ExpectError bool
	}{
		{
			"Test client role",
			enum.USER_ROLE_CLIENT,
			false,
		},
		{
			"Test moderator role",
			enum.USER_ROLE_MODERATOR,
			false,
		},
		{
			"Test client role with capitalization",
			"Client",
			true,
		},
		{
			"Test moderator role with capitalization",
			"Moderator",
			true,
		},
		{
			"Test unknown role",
			"unknown userRole",
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req := dto.DummyLoginRequest{UserType: test.UserRole}
			err := validateInputData(req)
			if err == nil && test.ExpectError {
				t.Errorf("expect error; got nothing.\n req = %+v", req)
			}
			if err != nil && !test.ExpectError {
				t.Errorf("expected NO error.\n got: %v\n req = %+v", err.Error(), req)
			}
		})
	}
}
