package results

import "time"

type AuthTokenResult struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
