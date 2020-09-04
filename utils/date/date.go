package date

import (
	"github.com/go-zs/gocommon/constant"
	"time"
)

func ParseMysqlDateNanoTs(ds string) (int64, error) {
	ts, err := time.ParseInLocation(constant.MysqlDateFormat, ds, time.Local)
	if err != nil {
		return 0, err
	}

	return ts.UnixNano(), err
}

func ParseMysqlDateMilliTs(ds string) (int64, error) {
	ts,err := ParseMysqlDateNanoTs(ds)

	return ts/int64(time.Millisecond), err
}

func ParseUtcDateNanoTs(ds string) (int64, error)  {
	ts, err := time.ParseInLocation(constant.UtcDateFormat, ds, time.Local)
	if err != nil {
		return 0, err
	}

	return ts.UnixNano(), err
}
func ParseUtcDateMilliTs(ds string) (int64, error) {
	ts, err := ParseUtcDateNanoTs(ds)

	return ts/int64(time.Millisecond), err
}