package timestamp

import "time"

type TimeStamp struct {
	timestamp int64 // UTC - タイムスタンプ
}

func NewTimeStamp() TimeStamp {
	now := time.Now().UTC()

	return TimeStamp{
		timestamp: now.Unix(),
	}
}
func NewTimeStampFromUnixTime(unixTime int64) *TimeStamp {
	datetime := time.Unix(unixTime, 0)
	datetime = datetime.UTC()

	return &TimeStamp{
		timestamp: datetime.Unix(),
	}
}

func (t *TimeStamp) String() string {
	datetime := time.Unix(t.timestamp, 0)
	datetime = datetime.UTC()

	return datetime.String()
}
