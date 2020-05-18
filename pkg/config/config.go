package config

import "os"

func DatabaseAdapter() string {
	return os.Getenv("DATABASE_ADAPTER")
}

func DatabaseURI() string {
	return os.Getenv("DATABASE_URI")
}

func GrpcAddress() string {
	return os.Getenv("GRPC_ADDRESS")
}
