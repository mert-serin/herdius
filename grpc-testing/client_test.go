package grpctesting

import (
	"context"
	"mertserin2/proto"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "localhost:50051"
)

func TestClientConnection(t *testing.T) {
	// Set up a connection to the server.
	creds, _ := credentials.NewClientTLSFromFile("../server-cert.pem", "")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		t.Error("Had problem with connection, NOT PASSED")
	}
	defer conn.Close()

	c := proto.NewHerdiusServerClient(conn)

	stream, err := c.CheckMax(context.Background())
	if err != nil {
		t.Error("Had problem with stream, NOT PASSED")
		return
	}

	err = stream.Send(&proto.MaxRequest{Val: int32(10)})
	err = stream.Send(&proto.MaxRequest{Val: int32(12)})
	err = stream.Send(&proto.MaxRequest{Val: int32(13)})
	err = stream.Send(&proto.MaxRequest{Val: int32(9)})

	if err != nil {
		t.Error("Had problem with stream, NOT PASSED")
		return
	}

	return
}
