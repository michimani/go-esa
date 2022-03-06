package gesa

import "time"

type Timestamp int64

func (ts *Timestamp) Time() *time.Time {
	if ts == nil {
		return nil
	}

	t := time.Unix(int64(*ts), 0)
	return &t
}

func (ts *Timestamp) SafeTimestamp() int64 {
	if ts == nil {
		return 0
	}

	return int64(*ts)
}

type RateLimitInformation struct {
	Limit     int
	Remaining int
	Reset     *Timestamp
}
