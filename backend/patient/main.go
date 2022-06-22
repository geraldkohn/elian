package main

import (
	"log"
	"net"

	"github.com/geraldkohn/elian/config"
	pb "github.com/geraldkohn/elian/proto/patient"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPatientServiceServer
}

func main() {
	var err error

	lis, err := net.Listen("tcp", config.Host+":"+config.PatientPort)
	if err != nil {
		log.Fatalf("patient component failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPatientServiceServer(s, &server{})
	log.Printf("patient component serve listening at: %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("patient component failed to serve: %v", err)
	}
}
