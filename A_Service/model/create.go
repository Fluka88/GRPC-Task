package model

import (
	"context"
	pb "my_task/my_task"
)

func (d *DAO) Create(title, year string) error{
	//give id and title to server using Create()

	in := &pb.Movie{
		Title:  title,
		Year:   year,
	}
	reply, err := d.client.Create(context.Background(), in)
	if err != nil {
		return err
	}
	println(reply.Message)
	return nil
}

func (d *DAO) Delete(title string) error{
	in := &pb.Request{
		Name: title,
	}
	reply, err := d.client.Delete(context.Background(), in)
	if err != nil {
		return err
	}
	println(reply.Message)
	return nil
}