package auth

import (
	"strings"
)

func TokenFromHeader(hd map[string]string) (*Token, bool) {
	var t string
	var ok bool

	// range possible auth headers
	for _, key := range []string{"authorization", "Authorization"} {
		t, ok = hd[key]
		if ok {
			break
		}
	}

	// no token
	if !ok {
		return nil, false
	}

	parts := strings.Split(t, " ")
	if len(parts) != 2 {
		return nil, false
	}
	return &Token{
		AccessToken: parts[1],
		TokenType:   parts[0],
	}, true
}

func HeaderWithToken(hd map[string]string, t *Token) map[string]string {
	// we basically only store access token
	hd["authorization"] = t.TokenType + " " + t.AccessToken
	return hd
}
