package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/damnn/tulahack/your-auth-service/tools"
	"github.com/jmoiron/sqlx"
)

type AuthRepositoryCore struct {
	db *sqlx.DB
}

func NewAuthRepositoryCore() *AuthRepositoryCore {
	return &AuthRepositoryCore{
		db: tools.DB,
	}
}

func (arc AuthRepositoryCore) FlowUser(ctx context.Context, user *User) (err error) {
	err = arc.db.GetContext(ctx, &user.Id, getUserById, user.Login)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		var stmt *sqlx.NamedStmt
		stmt, err = arc.db.PrepareNamed(createUser)
		if err != nil {
			return 
		}
	
		if err = stmt.GetContext(ctx, user, user); err != nil {
			return 
		}
	}

	return

}
