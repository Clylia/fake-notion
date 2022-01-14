package main

import (
	"context"
	accountpb "notion/account/api/gen/v1"
	authpb "notion/auth/api/gen/v1"
	"notion/shared/auth"
	"notion/shared/server"

	"github.com/namsral/flag"

	"log"
	"net/http"
	"net/textproto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

var addr = flag.String("addr", ":8080", "address to listen")
var accountAddr = flag.String("account_addr", "localhost:8081", "address for account service")
var authAddr = flag.String("auth_addr", "localhost:8082", "address for auth service")

func main() {
	flag.Parse()
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create zap logger: %v", err)
	}

	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	// mux: multiplexer
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers: true,
				UseProtoNames:  true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}), runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
		if key == textproto.CanonicalMIMEHeaderKey(runtime.MetadataHeaderPrefix+auth.ImpersonateAccountHeader) {
			return "", false
		}
		return runtime.DefaultHeaderMatcher(key)
	}))

	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "account",
			addr:         *accountAddr,
			registerFunc: accountpb.RegisterAccountServiceHandlerFromEndpoint,
		},
		{
			name: "auth",
			addr: *authAddr,
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
	}

	for _, s := range serverConfig {
		option := grpc.WithTransportCredentials(insecure.NewCredentials())
		err := s.registerFunc(
			c,
			mux,
			s.addr,
			[]grpc.DialOption{option})
		if err != nil {
			logger.Sugar().Fatalf("cannot register service %s: %v", s.name, err)
		}
	}

	logger.Sugar().Infof("grpc gateway started at %s", *addr)
	logger.Sugar().Fatal(http.ListenAndServe(*addr, mux))
}
