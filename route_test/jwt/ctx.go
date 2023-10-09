package jwt

import(
	"context"

	"github.com/golang-jwt/jwt"
)

type middlewareContextkey string

const tokenContextKey middlewareContextkey = "token"

func ContextWithToken(ctx context.Context, token *jwt.Token) context.Context{
	return context.WithValue(ctx,tokenContextKey,token)
}