package main

import (
	"context"
	"io/ioutil"
	"log"
	"notion/auth/auth"
	"notion/auth/auth/dao"
	"notion/auth/auth/token"
	"notion/shared/enc"
	"notion/shared/server"
	"time"

	authpb "notion/auth/api/gen/v1"

	"github.com/golang-jwt/jwt"
	"github.com/namsral/flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8082", "address to listen")
var mongoURI = flag.String("mongo_uri", "mongodb://localhost:27017", "mongo uri")
var privateKeyFile = flag.String("private_key_file", "auth/private.key", "private key file")

func main() {
	flag.Parse()
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI(*mongoURI))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}

	pkBytes, err := ioutil.ReadFile(*privateKeyFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name: "auth",
		Addr: *addr,
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &auth.Service{
				Monogo:         dao.NewMongo(mongoClient.Database("notion")),
				Logger:         logger,
				TokenExprie:    30 * time.Minute,
				TokenGenerator: token.NewJWTokenGen("notion/auth", privateKey),
				Decryptor:      enc.New(),
			})
		},
		Logger: logger,
	}))
}
