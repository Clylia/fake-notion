package main

import (
	"context"
	"log"
	"notion/shared/enc"
	"notion/shared/server"

	"github.com/namsral/flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"notion/account/account"
	"notion/account/account/dao"
	accountpb "notion/account/api/gen/v1"
)

var addr = flag.String("addr", ":8081", "address to listen")
var mongoURI = flag.String("mongo_uri", "mongodb://localhost:27017", "mongo uri")

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
	encryption := enc.New()
	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "account",
		Addr:   *addr,
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			accountpb.RegisterAccountServiceServer(s, &account.Service{
				Encryptor: encryption,
				Mongo:     dao.NewMongo(db),
				Logger:    logger,
			})
		},
	}))
}
