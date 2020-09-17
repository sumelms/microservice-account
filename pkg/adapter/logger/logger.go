package logger

import (
	"os"

	"github.com/sumelms/sumelms/microservice-user/pkg/config"

	"github.com/go-kit/kit/log"
)

func NewLogger(cfg *config.Config) log.Logger {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", cfg.Service,
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	return logger
}
