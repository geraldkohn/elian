package main

import (
	"log"
	"net"

	"github.com/geraldkohn/elian/config"
	pb "github.com/geraldkohn/elian/proto/agency"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAgencyServiceServer
}

func main() {
	var err error

	lis, err := net.Listen("tcp", config.Host+":"+config.AgencyPort)
	if err != nil {
		log.Fatalf("agency component failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAgencyServiceServer(s, &server{})
	log.Printf("agency component serve listening at: %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("agency component failed to serve: %v", err)
	}
}
