package webhooks

import "net/http"

type Auth struct {
}

func (a *Auth) addAuthHeader(r *http.Request, token string) {
	r.Header.Add("Authorization", token)
}
