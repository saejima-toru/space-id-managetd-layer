package timestamp

import "testing"

func TestNewTimeStamp(t *testing.T) {

	now := NewTimeStamp()
	duplicate := NewTimeStampFromUnixTime(now.timestamp)

	if now.String() != duplicate.String() {
		t.Fail()
	}
	println(now.String())
	println(duplicate.String())
}
