//go:telemetry

package telemetry

import (
	"runtime"
	"sync"

	"github.com/posthog/posthog-go"
)

var once sync.Once

var telemetrySingleton Telemetry = &BogusTelemetry{}

type BogusTelemetry struct{}

type RealTelemetry struct {
	Client *posthog.Client
	ApiKey string
	UserId string
}

// The telemetry config should be configured as a singleton,
// so that it can be accessed from anywhere in the codebase.
func InitTelemetry(Client *posthog.Client,
	ApiKey string,
	UserId string) Telemetry {
	once.Do(func() {
		client, _ := posthog.NewWithConfig(
			ApiKey,
			posthog.Config{
				Endpoint: "https://us.i.posthog.com",
			},
		)
		runtime.SetFinalizer(client, func(client *posthog.Client) {
			(*client).Close()
		})

		telemetrySingleton = &RealTelemetry{
			Client: &client,
			ApiKey: ApiKey,
			UserId: UserId,
		}
	})
	return telemetrySingleton
}

type Telemetry interface {
	CaptureEvent(event string) error
}

func GetTelemetry() Telemetry {
	return telemetrySingleton
}

func (telemetry *RealTelemetry) CaptureEvent(event string) error {
	userId := telemetry.UserId
	err := (*telemetry.Client).Enqueue(posthog.Capture{
		DistinctId: userId,
		Event:      event,
	})
	return err
}

func (telemetry *BogusTelemetry) CaptureEvent(event string) error {
	return nil
}
