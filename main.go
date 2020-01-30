package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/meso-org/meso-service-reference/database"
	pb "github.com/meso-org/meso-service-reference/proto/animal"
	"google.golang.org/grpc"
)

/**
https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2
*/
type Server struct {
}

func (s *Server) GetAnimals(ctx context.Context, req *pb.Animal) (*pb.Response, error) {

	return &pb.Response{Created: true, pb.Animal: {}}, nil
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": database.PrintHello(),
		})
	})

	r.GET("/animal/:name", func(c *gin.Context) {
		animal, err := database.GetAnimal(c.Param("name"))
		if err != nil {
			c.String(404, err.Error())
			return
		}
		c.JSON(200, animal)
	})

	// use the grpc library to create a new server and then register the protobuf server.
	s := grpc.NewServer()

	pb.RegisterAnimalServiceServer(s)
	r.Run(":3000")
}
