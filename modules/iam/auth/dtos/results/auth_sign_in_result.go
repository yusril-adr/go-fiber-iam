package results

type AuthSignInResult struct {
	AccessToken  AuthTokenResult `json:"access_token"`
	RefreshToken AuthTokenResult `json:"refresh_token"`
}
