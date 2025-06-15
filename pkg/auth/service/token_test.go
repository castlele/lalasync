package service

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestToken_GetTokenWithUserName(t *testing.T) {
	godotenv.Load()
	cases := []string{"castlelecs", "nikita", "javie"}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("Get Token Test: %v", i), func(t *testing.T) {
			token, err := NewToken(ts)

			if err != nil {
				t.Errorf("No error should raise generating token, but got: %v", err)
			}

			if len(token) == 0 {
				t.Errorf("Token length shouldn't be 0")
			}
		})
	}
}
