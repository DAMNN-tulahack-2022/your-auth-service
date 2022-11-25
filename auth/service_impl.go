package auth

import (
	"context"
	"github.com/damnn/tulahack/your-auth-service/repository"
)

type AuthCoreService struct {
	ctx  context.Context
	repo repository.AuthRepository
}

func NewAuthService(ctx context.Context, repo repository.AuthRepository) *AuthCoreService {
	return &AuthCoreService{
		ctx:  ctx,
		repo: repo,
	}
}

func (acs *AuthCoreService) SignUp() (err error) {

	return
}

func (acs *AuthCoreService) SignIn() (err error) {

	return
}

func (acs *AuthCoreService) RefreshFlow() (err error) {

	return
}

func (acs *AuthCoreService) BlockFlow() (err error) {

	return
}
