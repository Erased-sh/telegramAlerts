package tick

import (
	"context"
	"fmt"
	"time"
)

func GenerateTicker(ctx context.Context, amount int64, timeTag string, action func() error) {
	duration, err := convertDuration(time.Duration(amount), timeTag)
	if err != nil {
		return
	}
	ticker := time.NewTicker(time.Duration(amount) * duration)
	go func() {
		var err error
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				if err = action(); err != nil {
					return
				}
			}
		}
	}()
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
		err = fmt.Errorf("invalid tag")
	}
	return
}
