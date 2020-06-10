package grpc

import (
	"context"
	"net"
	"time"

	"github.com/mvrilo/app-poc/pkg/config"
	"github.com/mvrilo/app-poc/pkg/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type ClientConnInterface = grpc.ClientConnInterface

type Server struct {
	GatewayMux *runtime.ServeMux
	net.Listener
	*grpc.Server
}

type Client struct {
	*grpc.ClientConn
}

func NewClient() (*Client, error) {
	conn, err := grpc.Dial(config.GrpcAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	cli := &Client{
		ClientConn: conn,
	}
	return cli, err
}

func NewServer() (*Server, error) {
	listener, err := net.Listen("tcp", config.GrpcAddress())
	if err != nil {
		return nil, err
	}

	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, rec interface{}) (err error) {
			logger.Logger.Sugar().Error("gRPC server recovered in", rec)
			return status.Errorf(codes.Internal, "Recovered in %v\n", rec)
		}),
	}

	grpc_zap.ReplaceGrpcLoggerV2(logger.Logger)

	zapOpts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ms", duration.Milliseconds())
		}),
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(recoveryOpts...),
			// grpc_zap.StreamServerInterceptor(logger.Logger, zapOpts...),
			grpc_ctxtags.StreamServerInterceptor(),
			// 	grpc_opentracing.StreamServerInterceptor(),
			// 	grpc_prometheus.StreamServerInterceptor,
			// 	grpc_auth.StreamServerInterceptor(myAuthFunction),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
			grpc_zap.UnaryServerInterceptor(logger.Logger, zapOpts...),
			grpc_ctxtags.UnaryServerInterceptor(),
			// 	grpc_opentracing.UnaryServerInterceptor(),
			// 	grpc_prometheus.UnaryServerInterceptor,
			// 	grpc_auth.UnaryServerInterceptor(myAuthFunction),
		)),
	)

	reflection.Register(grpcServer)
	return &Server{
		GatewayMux: runtime.NewServeMux(),
		Listener:   listener,
		Server:     grpcServer,
	}, nil
}

func (s *Server) GatewayOpts() []grpc.DialOption {
	return []grpc.DialOption{grpc.WithInsecure()}
}

func (s *Server) GatewayAddr() string {
	return config.GrpcAddress()
}

func (s *Server) Start() error {
	return s.Server.Serve(s.Listener)
}

func (s *Server) Stop() error {
	return s.Listener.Close()
}
