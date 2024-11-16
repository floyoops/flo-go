package middleware

import (
	"github.com/floyoops/flo-go/backend/pkg/bus"
	"github.com/floyoops/flo-go/backend/pkg/logger"
)

func LoggingMiddleware(logger logger.Logger) bus.CommandMiddleware {
	return func(next bus.CommandHandler) bus.CommandHandler {
		return bus.CommandHandlerFunc(func(command bus.Command) ([]bus.Event, error) {
			logger.Infof("Handling command %s", command.Identifier().String())
			events, err := next.Handle(command)
			if err != nil {
				logger.Errorf("Error handling command: %s", err)
			}

			for _, event := range events {
				logger.Infof("Generated event: %s by command %s", event.Identifier().String(), command.Identifier().String())
			}

			return events, err
		})
	}
}
