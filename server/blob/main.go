package main

import (
	"flag"
	"log"
	"notion/blob/blob"
	"notion/shared/server"

	blobpb "notion/blob/api/gen/v1"

	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8084", "address to listen")

func main() {
	flag.Parse()

	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "blob",
		Addr:   *addr,
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobServiceServer(s, &blob.Service{
				Logger: logger,
			})
		},
	}))
}
