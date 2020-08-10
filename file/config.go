package file

import "time"

var (
	defaultConfig = inputConfig{
		scanFrequency: 60 * time.Second,
	}
)

type inputConfig struct {
	scanFrequency time.Duration
}
