package resellerclub

import (
	"encoding/json"
	"time"
)

type Bool bool

type Int64 int64

type Float64 float64

type Time time.Time

func (v *Bool) UnmarshalJSON(data []byte) error {
	var err error
	var b bool

	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		err = json.Unmarshal(data[1:len(data)-1], &b)
	} else {
		err = json.Unmarshal(data, &b)
	}

	*v = Bool(b)

	return err
}

func (v *Int64) UnmarshalJSON(data []byte) error {
	var err error
	var i int64

	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		err = json.Unmarshal(data[1:len(data)-1], &i)
	} else {
		err = json.Unmarshal(data, &i)
	}

	*v = Int64(i)

	return err
}

func (v *Float64) UnmarshalJSON(data []byte) error {
	var err error
	var f float64

	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		err = json.Unmarshal(data[1:len(data)-1], &f)
	} else {
		err = json.Unmarshal(data, &f)
	}

	*v = Float64(f)

	return err
}

func (v *Time) UnmarshalJSON(data []byte) error {
	var err error
	var t time.Time
	var i int64

	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		err = json.Unmarshal(data[1:len(data)-1], &i)
		if err != nil {
			t, err = time.Parse("2006-01-02 15:04:05.999999999-07", string(data[1:len(data)-1]))
		}
		if err != nil {
			t, err = time.Parse(time.RFC3339, string(data[1:len(data)-1]))
		}
	} else {
		err = json.Unmarshal(data, &i)
	}
	if err != nil {
		return err
	}

	if t.IsZero() {
		t = time.Unix(i, 0)
	}

	*v = Time(t)

	return nil
}
