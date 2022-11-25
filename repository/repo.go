package repository

type AuthRepository interface {
	CreateAccount() error
	VerifyAccount() error
	BlockAccount() error
}

