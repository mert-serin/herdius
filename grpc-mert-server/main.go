package main

import (
	"log"

	"net"

	"mertserin2/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

var preVal int32

// SayHello implements helloworld.GreeterServer
func (s *server) CheckMax(stream proto.HerdiusServer_CheckMaxServer) error {

	for {

		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
		}

		req, err := stream.Recv()
		if err != nil {
			log.Println(err)
			return err
		}

		if req.Val > preVal {
			preVal = req.Val
			err = stream.Send(&proto.MaxResponse{Val: preVal})
			if err != nil {
				log.Println(err)
				return err
			}
		}

	}
}

func main() {
	creds, _ := credentials.NewServerTLSFromFile("../server-cert.pem", "../server-key.pem")
	s := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterHerdiusServerServer(s, &server{})
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
