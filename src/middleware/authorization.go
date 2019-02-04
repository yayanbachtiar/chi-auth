package middleware

import (
	"fmt"
	"mama-chi/src/helper"
	"mama-chi/src/models"
	"net/http"
	"time"

	"github.com/gbrlsnchs/jwt"
	"github.com/go-chi/render"
)

// JwtMidleware middleware restricts access to just administrators.
func JwtMidleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		auth := r.Header.Get("Authorization")
		if auth == "" {
			render.Render(w, r, helper.ErrForbiden)
			return
		}
		hs256 := jwt.NewHS256("to_secret")
		payload, sig, err := jwt.Parse(auth)
		if err != nil {
			render.Render(w, r, helper.ErrForbiden)
			return
		}

		if err = hs256.Verify(payload, sig); err != nil {
			render.Render(w, r, helper.ErrForbiden)
			return

		}
		var jot models.Token
		if err = jwt.Unmarshal(payload, &jot); err != nil {
			render.Render(w, r, helper.ErrForbiden)
			return
		}

		// Validate fields.
		iatValidator := jwt.IssuedAtValidator(now)
		expValidator := jwt.ExpirationTimeValidator(now)
		audValidator := jwt.AudienceValidator("admin")
		if err = jot.Validate(iatValidator, expValidator, audValidator); err != nil {
			switch err {
			case jwt.ErrIatValidation:
				// handle "iat" validation error
			case jwt.ErrExpValidation:
				// handle "exp" validation error
			case jwt.ErrAudValidation:
				// handle "aud" validation error
			default:
				render.Render(w, r, helper.ErrForbiden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// BasicAuth middleware
func BasicAuth(realm string, credentials map[string]string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username, password, ok := r.BasicAuth()
			if !ok {
				unauthorized(w, realm)
				return
			}

			validPasswords, userFound := credentials[username]
			if !userFound {
				unauthorized(w, realm)
				return
			}
			if validPasswords == password {
				next.ServeHTTP(w, r)
				return
			}

			unauthorized(w, realm)
		})
	}
}
func unauthorized(w http.ResponseWriter, realm string) {
	helper.LogErr("access without basic-auth")
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}
