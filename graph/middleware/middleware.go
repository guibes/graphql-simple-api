package middleware

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/guibes/graphql-simple-api/graph/model"
	"github.com/guibes/graphql-simple-api/graph/utils"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type UserAuth struct {
	UserID    string
	Name      string
	IPAddress string
	Token     string
}

func JwtMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := TokenFromHTTPRequestgo(r)

			userID := UserIDFromHTTPRequestgo(token)
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			userAuth := UserAuth{
				UserID:    userID,
				IPAddress: ip,
			}

			ctx := context.WithValue(r.Context(), userCtxKey, &userAuth)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func TokenFromHTTPRequestgo(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	var tokenString string

	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) > 1 {
		tokenString = splitToken[1]
	}
	return tokenString
}

func UserIDFromHTTPRequestgo(tokenString string) string {
	token, err := utils.DecodeJwt(tokenString)
	if err != nil {
		return ""
	}
	if claims, ok := token.Claims.(*model.UserClaims); ok && token.Valid {
		if claims == nil {
			return ""
		}
		return claims.UserID
	}
	return ""
}

func GetAuthFromContext(ctx context.Context) *UserAuth {
	raw := ctx.Value(userCtxKey)
	fmt.Println(raw)
	if raw == nil {
		return nil
	}

	return raw.(*UserAuth)
}
