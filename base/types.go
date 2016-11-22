package base

import (
	"encoding/json"
	"time"
)

type NilField struct {
	Set bool
}

type Time string

func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	if s == "" {
		*t = Time(s)
		return nil
	}

	var parsed time.Time
	parsed, err = time.Parse(SERVER_TIME_FORMAT, s)
	if err != nil {
		return err
	}
	*t = Time(parsed.Format(TIME_FORMAT))
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(*t))
}

type BinaryResponse string

type UnixTimeStamp string

func (u *UnixTimeStamp) UnmarshalJSON(b []byte) error {
	n := int64(0)
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	ut := time.Unix(n, 0)
	*u = UnixTimeStamp(ut.String())

	return nil
}

type JSTimeStamp string

func (u *JSTimeStamp) UnmarshalJSON(b []byte) error {
	n := int64(0)
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	//Javascript timestamp is in milliseconds so convert to seconds
	n = n / 1000

	ut := time.Unix(n, 0)
	*u = JSTimeStamp(ut.String())

	return nil
}
