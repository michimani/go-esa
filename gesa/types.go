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

// PageNumber is int type for nullable value
// included pagination response
type PageNumber int

// NewPageNumber generates *PageNumber.
// If a negative value is passed,
// a value equivalent to 0 is returned.
func NewPageNumber(n int) *PageNumber {
	if n < 0 {
		n = 0
	}
	p := PageNumber(n)
	return &p
}

func (pn *PageNumber) IsNull() bool {
	return pn == nil || *pn == 0
}

func (pn *PageNumber) SafeInt() int {
	if pn == nil {
		return 0
	}

	return int(*pn)
}
