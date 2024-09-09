package auth

import (
	"context"
	"log"
	"net/http"
)

func (s *Service) IsAuthorized() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "opName", "IsAuthorized")

			cookies := r.Cookies()
			token := ""
			for _, cookie := range cookies {
				if cookie.Name == "token" {
					token = cookie.Value
				}
			}

			if token == "" {
				log.Println("token is empty")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			exists, err := s.token.IsTokenExpired(ctx, token)
			if err != nil {
				log.Println("IsTokenExpired: " + err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !exists {
				log.Println("token expired")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
