package transport

import (
	"context"
	"github.com/damnn/tulahack/your-auth-service/auth"
	"os"
	"os/signal"
	"syscall"
)

type (
	gatewayProxy struct {
		// interface
		// proto.UnimplementedTasksAdminServiceServer
	}

	AuthProxy struct {
		// login
		service auth.AuthorizationCore

		proxy *gatewayProxy

		// parent
		ctx context.Context

		sig chan os.Signal
	}
)

func NewAuthProxy(ctx context.Context, service auth.AuthorizationCore) *AuthProxy {
	return &AuthProxy{
		ctx:     ctx,
		service: service,
		proxy:   &gatewayProxy{},
	}
}

func (ap *AuthProxy) Run() error {
	/* 	// rest api server init
	   	mux := runtime.NewServeMux(
	   		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
	   			MarshalOptions: protojson.MarshalOptions{EmitUnpopulated: true},
	   		}),
	   		runtime.WithRoutingErrorHandler(helpers.RoutingErrorsHandler),
	   		runtime.WithErrorHandler(errorHandler),
	   	)

	   	if err := proto.RegisterTasksAdminServiceHandlerServer(ctx, mux, adminServer); err != nil {
	   		log.Panic(err)
	   	} */

	/* 	if err := http.ListenAndServe(host, mux); err != nil {
		return err
	} */

	return nil
}

func (ap *AuthProxy) Shotdown() error {

	return nil
}

func (ap *AuthProxy) GracefullNotify() error {
	signal.Notify(ap.sig, os.Interrupt, syscall.SIGTERM)
	<-ap.sig

	return nil
}
