package main

import (
	"fmt"
	"github.com/WhoIsYgim/swampy_rabbit/internal/config"
	"github.com/WhoIsYgim/swampy_rabbit/internal/pkg/rabbit"
	"github.com/WhoIsYgim/swampy_rabbit/internal/service"
	swampyservice "github.com/WhoIsYgim/swampy_rabbit/pkg/api/swampy"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}
	proucer, consumer, err := rabbit.SetupWorkers(&cfg.RabbitMQConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = consumer.Run()
	if err != nil {
		log.Fatal(err)
	}
	handler := service.NewSwampyService(proucer, consumer)
	srv := grpc.NewServer()
	swampyservice.RegisterSwampyServiceServer(srv, handler)
	fmt.Printf("Port %d", cfg.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
