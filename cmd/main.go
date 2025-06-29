package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vadkvadrad/track-service/internal/config"
	"github.com/vadkvadrad/track-service/internal/handler"
	"github.com/vadkvadrad/track-service/internal/repository"
	"github.com/vadkvadrad/track-service/internal/service"
	db "github.com/vadkvadrad/track-service/pkg"
	pb "github.com/vadkvadrad/track-service/proto"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("Config err: " + err.Error())
	}

	dbConfig := db.NewDbConfig(cfg.Db.Dsn)
	db := db.InitDB(dbConfig)

	repo := repository.NewTrackRepository(db)
    svc := service.NewTrackService(repo)
    hdl := handler.NewTrackHandler(svc)

	grpcServer := grpc.NewServer()
    pb.RegisterTrackServiceServer(grpcServer, hdl)

    lis, err := net.Listen("tcp", ":"+cfg.App.Port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    log.Println("Track Service is running on port :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }

	fmt.Println("TRACK SERVICE started successfully on port: ", cfg.App.Port)
}