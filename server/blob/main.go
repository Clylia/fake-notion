package main

import (
	"flag"
	"log"
	"notion/blob/blob"
	"notion/blob/cos"
	"notion/shared/server"

	blobpb "notion/blob/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8084", "address to listen")
var cosAddr = flag.String("cos_addr", "<URL>", "cos address")
var cosSecID = flag.String("cos_sec_id", "<SEC_ID>", "cos secret id")
var cosSecKey = flag.String("cos_sec_key", "<SEC_KEY>", "cos secret key")

func main() {
	flag.Parse()

	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	// "https://notion-1305000400.cos.ap-shenzhen-fsi.myqcloud.com",
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
			})
		},
	}))
}
