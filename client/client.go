package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Galionme/hercules/grpcapi"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var (
		opts   []grpc.DialOption
		conn   *grpc.ClientConn
		err    error
		client grpcapi.AdminClient
	)

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	opts = append(opts, grpc.WithInsecure())
	if conn, err = grpc.Dial(fmt.Sprintf("%s:%s", os.Getenv("ADMIN_HOST"), os.Getenv("ADMIN_PORT")), opts...); err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client = grpcapi.NewAdminClient(conn)
	var cmd = new(grpcapi.Command)
	cmd.In = os.Args[1]
	ctx := context.Background()
	cmd, err = client.RunCommand(ctx, cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmd.Out)
}
