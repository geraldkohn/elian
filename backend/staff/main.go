package main

import (
	"log"
	"net"

	"github.com/geraldkohn/elian/config"
	pb "github.com/geraldkohn/elian/proto/staff"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStaffServiceServer
}

func main() {
	var err error

	lis, err := net.Listen("tcp", config.Host + ":" + config.StaffPort)
	if err != nil {
		log.Fatalf("Staff component failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStaffServiceServer(s, &server{})
	log.Printf("staff component serve listening at: %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("staff component failed to serve: %v", err)
	}
}