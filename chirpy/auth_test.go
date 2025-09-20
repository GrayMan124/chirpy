package main

import (
	"fmt"
	"github.com/GrayMan124/chirpy/internal/auth"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	cases := []struct {
		user    uuid.UUID
		secret  string
		expires time.Duration
	}{
		{
			user:    uuid.New(),
			secret:  "Takjakpanjezuspowiedzial",
			expires: time.Minute,
		},
		{
			user:    uuid.New(),
			secret:  "TegoTamtego",
			expires: time.Second,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Running case %v", i), func(t *testing.T) {
			strToken, err := auth.MakeJWT(c.user, c.secret, c.expires)
			if err != nil {
				t.Errorf("Failed to make JWT")
			}
			validatedID, err := auth.ValidateJWT(strToken, c.secret)
			if err != nil {
				t.Errorf("Failed to validate string")
			}
			if validatedID != c.user {
				t.Errorf("Got wrong user!")
			}
			time.Sleep(time.Second * 2)
			validatedID, err = auth.ValidateJWT(strToken, c.secret)
			if c.expires <= time.Second {
				if err == nil {
					t.Errorf("This token should fail!")
				}
			} else {
				if err != nil {
					t.Errorf("This token should pass!")
				} else if validatedID != c.user {
					t.Errorf("Got wrong user!")
				}
			}
		})

	}

}
