package main

import (
	pb "POC-rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	PORT = ":3000"
)

func callSayHello(client pb.GreetPackageClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.HelloRequest{
		Name: "Iresh",
	})
	if err != nil {
		log.Fatalf("Shit happened %v", err)
	}
	log.Println(res.Greeting)
}

func main() {
	conn, err := grpc.Dial("localhost"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not start %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetPackageClient(conn)
	callSayHello(client)

}
