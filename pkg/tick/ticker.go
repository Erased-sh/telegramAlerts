package tick

import (
	"log/slog"
	"time"
)

func GenerateTicker(amount int64, timeTag string, action func() error) error {
	duration, err := convertDuration(time.Duration(amount), timeTag)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err = action(); err != nil {
					slog.Error(errMalformedAction.Error())
					return
				}
			}
		}
	}()
	ticker.Stop()
	return nil
}

func convertDuration(amount time.Duration, timeTag string) (duration time.Duration, err error) {
	switch timeTag {
	case "ms":
		duration = amount * time.Millisecond
	case "s":
		duration = amount * time.Second
	case "m":
		duration = amount * time.Minute
	case "h":
		duration = amount * time.Hour
	default:
		err = errInvalidTag
	}
	return
}
