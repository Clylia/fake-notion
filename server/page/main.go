package main

import (
	"context"
	"flag"
	"log"
	"notion/page/page/dao"
	"notion/shared/server"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pagepb "notion/page/api/gen/v1"
	"notion/page/page"
)

var addr = flag.String("addr", ":8083", "address to listen")
var mongoURI = flag.String("mongo_uri", "mongodb://localhost:27017", "mongo uri")
var authPublicKeyFile = flag.String("auth_public_key_file", "shared/auth/public.key", "public key file")

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
	db := mongoClient.Database("notion")
	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "page",
		Addr:   *addr,
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			pagepb.RegisterPageServiceServer(s, &page.Service{
				Mongo:  dao.NewMongo(db),
				Logger: logger,
			})
		},
		AuthPublicKeyFile: *authPublicKeyFile,
	}))
}
