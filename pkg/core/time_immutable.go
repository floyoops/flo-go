package core

import "time"

type TimeImmutable struct {
	value time.Time
}

func NewTimeImmutableNow() TimeImmutable {
	return TimeImmutable{value: time.Now()}
}

func (t TimeImmutable) ToMysqlDateTimeUTC() string {
	return t.value.UTC().Format("2006-01-02 15:04:05")
}
