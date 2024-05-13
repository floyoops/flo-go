package core

import "time"

type UtcTimeImmutable struct {
	value time.Time
}

func NewUtcTimeImmutableNow() UtcTimeImmutable {
	return UtcTimeImmutable{value: time.Now().UTC()}
}

func (t UtcTimeImmutable) ToMysqlDateTime() string {
	return t.value.Format("2006-01-02 15:04:05")
}
