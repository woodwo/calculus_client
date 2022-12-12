package main

import (
	"github.com/woodwo/calculus/grpc/proto"
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func catsFromCalculus(c proto.CalculusClient) (string, error) {
	ctx := context.TODO()
	fibValue, err := c.Fibonacci(ctx, &proto.Empty{})
	if err != nil {
		return "", err
	}
	return strings.Repeat("🐈", int(fibValue.Value)), nil
}

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("No calculus")
	}
	defer conn.Close()

	client := proto.NewCalculusClient(conn)
	cats, err := catsFromCalculus(client)
	fmt.Printf(cats)
}
