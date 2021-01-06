package main

import (
	"context"
	"flag"
	"strconv"
	"test_projects/grpc-server/pkg/api"

	"github.com/reconquest/pkg/log"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	connection, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewAdderClient(connection)
	response, err := client.Add(
		context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof(nil, "result: %d", response.GetResult())
}
