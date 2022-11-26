package repository

import "context"

type AuthRepository interface {
	FlowUser(ctx context.Context, user *User) (err error)
}