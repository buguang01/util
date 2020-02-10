package util

import (
	"time"
)

var (
	BeiJing     = time.FixedZone("CST", 8*3600)
	UTCLocation = time.UTC
	Location    = UTCLocation
)

//TimeConvert 时间操作类

//GetCurrTime 当前UTC时间
func GetCurrTime() time.Time {
	return time.Now().In(Location)
}

//GetCurrTimeSecond 当前UTC时间精确到秒
func GetCurrTimeSecond() time.Time {
	result := GetCurrTime()
	result = time.Date(
		result.Year(),
		result.Month(),
		result.Day(),
		result.Hour(),
		result.Minute(),
		result.Second(),
		0,
		Location)
	return result
}

//GetCurrDate 当前时间的日期
func GetCurrDate() time.Time {
	result := GetCurrTime()
	result = time.Date(
		result.Year(),
		result.Month(),
		result.Day(),
		0,
		0,
		0,
		0,
		Location)
	return result
}

func GetDate(d time.Time) time.Time {
	result := time.Date(
		d.Year(),
		d.Month(),
		d.Day(),
		0,
		0,
		0,
		0,
		Location)
	return result
}
func GetMinDateTime() time.Time {
	result := time.Date(
		1970,
		1,
		1,
		0,
		0,
		0,
		0,
		Location)
	return result
}

func UnixToTime(d int64) time.Time {
	return time.Unix(d, 0).In(Location)
}

//设置获取的服务器时间的默认时区
func SetLocation(loc *time.Location) {
	Location = loc
}

//Utc时间
type TimeUtc struct {
	time.Time
}

func (this TimeUtc) MarshalJSON() ([]byte, error) {
	buf := NewStringInt64(this.Unix()).ToString()
	return []byte(buf), nil
}

func (this *TimeUtc) UnmarshalJSON(v []byte) error {
	this.Time = time.Unix(NewString(string(v)).ToInt64V(), 0).UTC()
	return nil
}
