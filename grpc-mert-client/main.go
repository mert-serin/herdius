package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"

	"mertserin2/proto"

	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "localhost:50051"
)

func main() {

	// Set up a connection to the server.
	creds, _ := credentials.NewClientTLSFromFile("../server-cert.pem", "")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewHerdiusServerClient(conn)

	stream, err := c.CheckMax(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	ctx := stream.Context()
	done := make(chan bool)

	//First goroutine to read input from user
	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter value: ")
			text, err := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if err != nil {
				fmt.Println("Wrong input, please try again")
			}
			i, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println("Can not convert to integer, please try again")
			}
			err = stream.Send(&proto.MaxRequest{Val: int32(i)})
			if err != nil {
				log.Println("Error 1: ", err)
				return
			}
		}
	}()

	//Second goroutine to check if message received from grpc-mert-server
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			fmt.Println("\nCurrent max is: ", resp.Val)
		}
	}()

	//Third goroutinte to check done message from chan
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
}
