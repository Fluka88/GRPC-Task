package model

import (
	"context"
	"github.com/jfarleyx/go-simple-cache"
	"my_task/A_Service/views"
	pb "my_task/my_task"
	"time"
)

const (
	address     = "localhost:50051"
)

type Cache struct {
	myCache *cache.Cache
}

type CacheInt interface {
	SetCache() *Cache
}

func NewCache() CacheInt{
	return &Cache{}
}


func (c *Cache) SetCache() *Cache {
	c.myCache = cache.New(5*time.Minute)
	c.myCache.OnExpired(c.myCache.Flush)
	return c
}

func (d *DAO) Read(title string, c *Cache)( []views.Movie, error){
	var todos []views.Movie
	item, found := c.myCache.Get(title)
	if found {
		todos = item.([]views.Movie)
		//println("found in cache")
	}else{
		request := pb.Request{Name: title}
		movieArr, err := d.client.Read(context.Background(), &request)
		if err != nil {
			panic(err)
		}
		data := views.Movie{}
		for i := 0; i < len(movieArr.Movie); i++ {
			data.Year = movieArr.Movie[i].Year
			data.Title = movieArr.Movie[i].Title
			todos = append(todos, data)
		}
		c.myCache.Set(title, todos)
		//println("added to cache")
	}
	return todos, nil
}