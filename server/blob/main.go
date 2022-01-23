package main

import (
	"context"
	"log"
	"notion/blob/blob"
	"notion/blob/blob/dao"
	"notion/blob/cos"
	"notion/shared/server"

	"github.com/namsral/flag"

	blobpb "notion/blob/api/gen/v1"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8084", "address to listen")
var mongoURI = flag.String("mongo_uri", "mongodb://localhost:27017", "mongo uri")

var cosAddr = flag.String("cos_addr", "<URL>", "cos address")
var cosSecID = flag.String("cos_sec_id", "<SEC_ID>", "cos secret id")
var cosSecKey = flag.String("cos_sec_key", "<SEC_KEY>", "cos secret key")
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
	// cos_addr like: "https://notion-1505000900.cos.ap-shenzhen-fsi.myqcloud.com"
	st, err := cos.NewService(
		*cosAddr,
		*cosSecID,
		*cosSecKey)
	if err != nil {
		logger.Fatal("cannot create cos service", zap.Error(err))
	}
	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "blob",
		Addr:   *addr,
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobServiceServer(s, &blob.Service{
				Logger:  logger,
				Storage: st,
				Mongo:   dao.NewMongo(db),
			})
		},
		AuthPublicKeyFile: *authPublicKeyFile,
	}))
}
