package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	pb "my_task/my_task"
)

type Server struct {
	pb.UnimplementedDBServiceServer
}

var con *sql.DB

func (s *Server) Connect(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received data source: %v", in.GetName())

	dataSourceName := in.GetName()
	if dataSourceName == "" {
		dataSourceName = "root:root@tcp(localhost:33060)/DATA"
	}

	var err error
	con, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return &pb.Reply{Message: "Connection with database not established"}, err
	}
	if err := con.Ping();err!=nil{
		return &pb.Reply{Message: "Connection with database not established"}, err
	}
	fmt.Println("Connected to Database: " + dataSourceName)
	return &pb.Reply{Message: "Connected to Database: " + dataSourceName}, nil

}

func (s *Server) Read(ctx context.Context, title *pb.Request) (*pb.MovieArray, error) {
	var rows *sql.Rows
	var err error

	if title.Name != "" {
		rows, err = con.Query("SELECT * FROM MOVIES WHERE TITLE=?", title.Name)
	}else{
		rows, err = con.Query("SELECT * FROM MOVIES")
	}
	if err != nil {
		return nil, err
	}
	movies := pb.MovieArray{}
	for rows.Next(){
		data := pb.Movie{}
		rows.Scan(&data.Title, &data.Year)
		movies.Movie = append(movies.Movie, &data)
	}
	return &movies, nil
}

func (s *Server) Create(ctx context.Context, in *pb.Movie) (*pb.Reply, error){
	insertQ, err := con.Query("INSERT INTO MOVIES VALUES(?, ?)", in.Title, in.Year)
	if err != nil {
		fmt.Printf(err.Error())
		return &pb.Reply{Message: "Could not insert values into table."}, err
	}
	defer insertQ.Close()
	return &pb.Reply{Message: "Created."}, nil
}

func (s *Server) Delete(ctx context.Context, in *pb.Request) (*pb.Reply, error){
	insertQ, err := con.Query("DELETE FROM TODO WHERE title=?", in.Name)
	if err != nil {
		fmt.Printf(err.Error())
		return &pb.Reply{Message: "Could not delete."}, err
	}
	defer insertQ.Close()
	return &pb.Reply{Message: "Deletion complete."}, nil
}

func (s *Server) Ping(ctx context.Context, in *pb.Request) (*pb.Reply, error){
	return &pb.Reply{Message: "Pong"}, nil
}