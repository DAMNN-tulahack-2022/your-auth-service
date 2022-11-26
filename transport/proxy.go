package transport

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/damnn/tulahack/your-auth-service/auth"
	"github.com/damnn/tulahack/your-auth-service/gen/proto"
	"github.com/damnn/tulahack/your-auth-service/tools"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

type (

	AuthProxy struct {
		proto.UnimplementedYourAdminServiceServer
		// login
		log    *tools.YourLogger
		config *tools.GenericEnvAppConfig
		service auth.AuthorizationCore
		// parent
		ctx context.Context

		sig chan os.Signal
	}
)

func NewAuthProxy(ctx context.Context, log *tools.YourLogger, config *tools.GenericEnvAppConfig, service auth.AuthorizationCore) *AuthProxy {
	return &AuthProxy{
		ctx:     ctx,
		service: service,
		log:     log,
		config:  config,
		sig:     make(chan os.Signal),
	}
}

func (ap *AuthProxy) Run() error {
	// rest api server init
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{EmitUnpopulated: true},
		}),
	)

	if err := proto.RegisterYourAdminServiceHandlerServer(ap.ctx, mux, ap); err != nil {
		log.Panic(err)
	}

	if err := http.ListenAndServe(ap.config.AuthProxyPort, cors(mux)); err != nil {
		return err
	}

	return nil
}

func (ap *AuthProxy) GracefullNotify() error {
	signal.Notify(ap.sig, os.Interrupt, syscall.SIGTERM)
	<-ap.sig

	return nil
}
