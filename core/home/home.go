package home

import (
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/app-poc/pkg/grpc"
	"github.com/mvrilo/app-poc/pkg/server"
)

type Home struct{}

func New(_ *grpc.Client) *Home {
	return &Home{}
}

func (h *Home) Register(s *server.Server) error {
	router := s.HttpServer.Router

	// TODO: panic when using context methods
	router.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("welcome to app-poc"))
	})

	return nil
}
