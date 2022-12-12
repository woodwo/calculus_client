package main

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/woodwo/calculus/grpc/proto"
	"github.com/woodwo/calculus/grpc/proto/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Test_nextValueFromCalculus(t *testing.T) {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("No calculus")
	}

	client := proto.NewCalculusClient(conn)

	t.Run("Happy Path", func(t *testing.T) {
		n1, _ := catsFromCalculus(client)
		if !strings.Contains(n1, "🐈") {
			t.Error("Fibonacci should be positive")
		}
	})

	t.Run("Another happy path", func(t *testing.T) {
		n1, _ := catsFromCalculus(client)
		n2, _ := catsFromCalculus(client)

		if len(n1) > len(n2) {
			t.Error("Fibonacci should grow")
		}
	})

	t.Run("Yet another happy path", func(t *testing.T) {
		n1, _ := catsFromCalculus(client)
		n2, _ := catsFromCalculus(client)
		n3, _ := catsFromCalculus(client)

		if len(n3) != len(n2)+len(n1) {
			t.Error("This is not a Fibonacci")
		}
	})
}

func Test_nextValueMocked(t *testing.T) {
	// TODO here are mocks
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	client := mock_proto.NewMockCalculusClient(mockCtl)

	gomock.InOrder(
		client.EXPECT().Fibonacci(gomock.Any(), gomock.Any()).Return(&proto.Value{Value: 1}, nil).Times(1),
	)

	t.Run("Happy Path", func(t *testing.T) {
		n1, _ := catsFromCalculus(client)
		if n1 != "🐈" {
			t.Error("Expect just one cat")
		}
	})
}
