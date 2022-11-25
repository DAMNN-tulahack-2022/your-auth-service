package db

import (
	"github.com/jmoiron/sqlx"
)

type DbConnectionPool struct {
	Db *sqlx.DB
}

const (
	_postgresGenericDriver = "postgres"
)

func BeginDbInstance() (*sqlx.DB, error) {
	/*
		 	db := new(DbInstance)
			db.Db, err = sqlx.Connect(_postgresGenericDriver, dsn)
			if err != nil {
				log.Panic(err.Error())
			}

			postgres.SetMaxIdleConns(3)
			postgres.SetMaxOpenConns(20)
			postgres.SetConnMaxIdleTime(15 * time.Second)
			postgres.SetConnMaxLifetime(30 * time.Second)

			return
	*/

	return &sqlx.DB{}, nil
}
