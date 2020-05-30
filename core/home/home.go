package home

import (
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/pkg/server"
)

type Home struct{}

func New(_ *grpc.Client) *Home {
	return &Home{}
}

func (h *Home) Register(s *server.Server) error {
	router := s.HttpServer.Router

	// TODO: panic when using context methods
	router.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("test home"))
	})

	return nil
}
