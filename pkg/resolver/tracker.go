package resolver

import (
	"time"
)

const (
	statusUnused    = "unused"
	statusResolving = "resolving"
	statusResolved  = "resolved"
	statusFailed    = "failed"

	timeFormat = time.RFC3339
)

type TrackerConfig struct {
	IncludeError bool
	IncludeData  bool
	TrackTime    bool
}

// Track wraps the resolver in a tracker to record durration
func Track(resolver Resolver, config TrackerConfig) Resolver {
	return &tracker{
		Typ:      resolverTypeTracker,
		Resolver: resolver,
		Status:   statusUnused,

		config: config,
	}
}

type tracker struct {
	Typ        string      `json:"type" yaml:"type"`
	Resolver   Resolver    `json:"data" yaml:"data"`
	Status     string      `json:"status" yaml:"status"`
	Start      string      `json:"start,omitempty" yaml:"start,omitempty"`
	DurationMS int64       `json:"duration_ms,omitempty" yaml:"duration_ms,omitempty"`
	Error      string      `json:"error,omitempty" yaml:"error,omitempty"`
	Response   interface{} `json:"response,omitempty" yaml:"response,omitempty"`

	config TrackerConfig
}

func (t *tracker) Get() (interface{}, error) {
	t.Status = statusResolving
	var start time.Time
	if t.config.TrackTime {
		start = time.Now()
		t.Start = start.Format(timeFormat)
	}
	data, err := t.Resolver.Get()
	if err != nil {
		t.Status = statusFailed
		if t.config.IncludeError {
			t.Error = err.Error()
		}
	} else {
		t.Status = statusResolved
		if t.config.IncludeData {
			t.Response = data
		}
	}
	if t.config.TrackTime {
		t.DurationMS = time.Since(start).Milliseconds()
	}
	return data, err
}
