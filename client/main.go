package main

import (
	"fmt"
	"github.com/adrianwg23/grpc-example/proto"
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	client := proto.NewArithmeticServiceClient(conn)

	g := gin.Default()

	g.GET("/add/:a/:b", func(context *gin.Context) {
		a, b, err := parseValues(context)
		if err != nil {
			return
		}

		req := &proto.Request{A: a, B: b}

		if res, err := client.Add(context, req); err == nil {
			context.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Result),
			})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/mult/:a/:b", func(context *gin.Context) {
		a, b, err := parseValues(context)
		if err != nil {
			return
		}

		req := &proto.Request{A: a, B: b}

		if res, err := client.Multiply(context, req); err == nil {
			context.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Result),
			})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func parseValues(context *gin.Context) (int64, int64, error) {
	a, err := strconv.ParseInt(context.Param("a"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
		return -1, -1, err
	}

	b, err := strconv.ParseInt(context.Param("b"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter B"})
		return -1, -1, err
	}

	return a, b, nil
}
