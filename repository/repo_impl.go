package repository

import "context"

type AuthCoreRepository struct {
	ctx context.Context
	
	db     struct{}
	cache  struct{}
	memory struct{}
}

func NewAuthCoreRepository(ctx context.Context) *AuthCoreRepository {
	return &AuthCoreRepository{
		ctx: ctx,
	}
}

func (acr *AuthCoreRepository) CreateAccount() (err error) {

	return
}

func (acr *AuthCoreRepository) BlockAccount() (err error) {

	return
}

func (acr *AuthCoreRepository) VerifyAccount() (err error) {

	return
}
