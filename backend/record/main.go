package main

import (
	"log"
	"net"

	"github.com/geraldkohn/elian/config"
	pb "github.com/geraldkohn/elian/proto/record"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	lis, err := net.Listen("tcp", config.Host+":"+config.RecordPort)
	if err != nil {
		log.Fatalf("record component failed to serve: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRecordServiceServer(s, &server{})
	log.Printf("record component serve listening at: %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("record component failed to serve: %v", err)
	}
}
