package telemetry

import (
	"errors"
	"sync"

	"github.com/posthog/posthog-go"
)

var once sync.Once

var telemetrySingleton *Telemetry

type Telemetry struct {
	Client *posthog.Client
	ApiKey string
	UserId string
}

// The telemetry config should be configured as a singleton,
// so that it can be accessed from anywhere in the codebase.
func InitTelemetry(Client *posthog.Client,
	ApiKey string,
	UserId string) *Telemetry {
	once.Do(func() {
		client, _ := posthog.NewWithConfig(
			ApiKey,
			posthog.Config{
				Endpoint: "https://us.i.posthog.com",
			},
		)
		defer client.Close()

		telemetrySingleton = &Telemetry{
			Client: &client,
			ApiKey: ApiKey,
			UserId: UserId,
		}
	})
	return telemetrySingleton
}

func GetTelemetry() (*Telemetry, error) {
	if telemetrySingleton == nil {
		return nil, errors.New("Telemetry not initialized")
	}
	return telemetrySingleton, nil
}

func CaptureEvent(telemetry *Telemetry, event string) error {
	userId := telemetry.UserId
	err := (*telemetry.Client).Enqueue(posthog.Capture{
		DistinctId: userId,
		Event:      event,
	})
	return err
}
