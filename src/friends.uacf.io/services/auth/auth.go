package auth

import (
	"golang.org/x/net/context"
)

func NewContext(ctx context.Context, userId int64) context.Context {
	return context.WithValue(ctx, authKey(0), &Auth{userId})
}

func FromContext(ctx context.Context) *Auth {
	a, _ := ctx.Value(authKey(0)).(*Auth)
	return a
}

type Auth struct {
	UserId int64
}

// Return true if the user is authenticated.
//
// This checks for a nil auth so it can return a sensible value when not
// logged in.
func (a *Auth) Authenticated() bool {
	return a != nil && a.UserId != 0
}

// Return true if the user is authenticated and the user ID matches the given ID.
//
// This function works via chaining even when there's no Auth because it
// first checks for a nil auth.
func (a *Auth) MatchUserId(id int64) bool {
	return a != nil && id != 0 && id == a.UserId
}

type authKey int
