package middleware

import (
	"github.com/floyoops/flo-go/backend/pkg/bus"
	"github.com/floyoops/flo-go/backend/pkg/logger"
)

func LoggingMiddleware(logger logger.Logger) bus.CommandMiddleware {
	return func(next bus.CommandHandler) bus.CommandHandler {
		return bus.CommandHandlerFunc(func(command bus.Command) error {
			logger.Info("Handling command: %T", command)
			err := next.Handle(command)
			if err != nil {
				logger.Error("Error handling command: %s", err)
			}
			return err
		})
	}
}
