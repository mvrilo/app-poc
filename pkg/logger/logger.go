package logger

// import "github.com/mvrilo/storepoc/pkg/config"
import "go.uber.org/zap"

var Logger *zap.Logger = load()

func load() *zap.Logger {
	// config.Env

	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return zapLogger
}
