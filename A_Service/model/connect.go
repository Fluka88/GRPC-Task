package model

import (
	"A_Service/views"
	"context"
	"log"
	pb "my_task/my_task"
	"os"
	"time"

	"google.golang.org/grpc"
)

type DAO struct {
	client pb.DBServiceClient
	conn   grpc.ClientConn
}

type DAOInt interface {
	Connect()
	Close()
	Read(title string, c *Cache) ([]views.Movie, error)
	Create(title, year string) error
	Delete(title string) error
	Ping() bool
}

func NewDAO() DAOInt {
	return &DAO{}
}

var DB DAOInt

//connect to server and tell it to connect to db
func (d *DAO) Connect() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	d.client = pb.NewDBServiceClient(conn)

	// Contact the server and print out its response.
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := d.client.Connect(ctx, &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	log.Printf("Success: %s", r.GetMessage())
}

func (d *DAO) Close() {
	d.conn.Close()
}

func (d *DAO) Ping() bool {
	request := pb.Request{Name: "ping"}
	_, err := d.client.Ping(context.Background(), &request)
	if err != nil {
		return false
	}
	return true
}
