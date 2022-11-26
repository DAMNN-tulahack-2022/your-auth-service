package auth

import "context"

type AuthorizationCore interface {
	GithubLogin(ctx context.Context, login string) (id uint32, err error)
}
