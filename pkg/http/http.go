package http

import (
	"net/http"
	"strings"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"github.com/mvrilo/app-poc/pkg/config"
	"github.com/mvrilo/app-poc/pkg/logger"
)

type Server struct {
	*http.Server
	Router *gin.Engine
}

func NewServer() *Server {
	router := gin.New()
	router.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, true))

	httpServer := &http.Server{
		Addr:    config.HttpAddress(),
		Handler: router,
	}

	s := &Server{
		Server: httpServer,
		Router: router,
	}

	return s
}

func (s *Server) AddGrpcGateway(grpcPathPrefix string, grpcHandler http.Handler) {
	s.Router.Any(grpcPathPrefix+"/*any", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.Replace(r.URL.Path, grpcPathPrefix, "", -1)
		grpcHandler.ServeHTTP(w, r)
	}))
}

func (s *Server) Start() error {
	return s.Server.ListenAndServe()
}

func (s *Server) Stop() {
	s.Close()
}
