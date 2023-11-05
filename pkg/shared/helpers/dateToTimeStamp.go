package helpers

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func DateToTimeStamp(date time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: date.Unix(),
		Nanos:   int32(date.Nanosecond()),
	}
}
