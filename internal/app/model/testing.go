package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "test@test.tt",
		Password: "password",
	}
}
