package main

import (
	"context"

	"github.com/damnn/tulahack/your-auth-service/auth"
	"github.com/damnn/tulahack/your-auth-service/repository"
	"github.com/damnn/tulahack/your-auth-service/transport"
	"github.com/nickolation/your-api-tools/your"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := your.SetupLogger()
	config, err := your.LoadDammnAppConfig(logger) 
	if err != nil {
		logger.Log().Fatal().AnErr("load env configs err", err)
	}

	db, err := your.BeginDbInstance(config, logger) 
	if err != nil {
		logger.Log().Fatal().AnErr("db connection err", err)
	}

	cache, err := your.NewExternalCache(ctx, config, logger)
	if err != nil {
		logger.Log().Fatal().AnErr("connect to external cache err", err)
	} 

	// mock
	logger.Log().Debug().Interface("a", db)
	logger.Log().Debug().Interface("a", cache)

	repo := repository.NewAuthCoreRepository(ctx)

	serivce := auth.NewAuthService(ctx, repo)

	proxyServer := transport.NewAuthProxy(ctx, serivce)

	go func() {
		if err := proxyServer.Run(); err != nil {
			//
		}
	}()

	proxyServer.GracefullNotify()
	if err := proxyServer.Shotdown(); err != nil {
		// log

	}

	// closeDeps

}
