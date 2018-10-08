package utils

// Salute to Mr.Long

import (
	"regexp"
	"strings"
	"time"
)

func ParseTime(v string) (t time.Time) {
	// format:
	//	[Y-m-d] [H:[M:S]]

	v = strings.Trim(v, " ")

	re := regexp.MustCompile(`^((\d+)\-(\d+)\-(\d+))?\s*((\d+):(\d+)(:(\d+))?)?$`)
	arr := re.FindStringSubmatch(v)
	if arr == nil {
		return
	}

	y, m, d := int(Atoi32(arr[2])), time.Month(Atoi32(arr[3])), int(Atoi32(arr[4]))
	H, M, S := int(Atoi32(arr[6])), int(Atoi32(arr[7])), int(Atoi32(arr[9]))

	if y == 0 && m == 0 && d == 0 {
		y, m, d = time.Now().Date()
	}

	t = time.Date(y, m, d, H, M, S, 0, time.Local)

	return
}

func ParseRelativeTime(t0 time.Time, v string) (t time.Time) {
	// format:
	//	* @				t0
	//	* @2-10-50		@相对t0天数-时-分

	v = strings.Trim(v, " ")

	if !strings.HasPrefix(v, "@") {
		return
	}

	e := strings.Trim(v[1:], " ")
	if e == "" {
		t = t0
	} else {
		p := strings.Split(e, "-")

		// day
		d := int(Atoi32(p[0]))

		// hour
		h := 0
		if len(p) > 1 {
			h = int(Atoi32(p[1]))
		}

		// minute
		m := 0
		if len(p) > 2 {
			m = int(Atoi32(p[2]))
		}

		// second
		s := 0
		if len(p) > 3 {
			s = int(Atoi32(p[3]))
		}

		// calc
		t = BeginOfDay(t0).
			AddDate(0, 0, d).
			Add(time.Hour * time.Duration(h)).
			Add(time.Minute * time.Duration(m)).
			Add(time.Second * time.Duration(s))
	}

	return
}

// 返回从本周开始计算的秒数(周：从周一[1]开始，到周日[7]结束)
// from "6~23:00" to 514800
func ParseWeekTime(date string) (val int, ret bool) {
	date = strings.Trim(date, " ")
	reg := regexp.MustCompile(`^([1-7])~([0-9]{1,2}):([0-9]{1,2})$`)
	arr := reg.FindStringSubmatch(date)
	if arr == nil {
		val, ret = 0, false
		return
	}

	wday := int(core.Atoi32(arr[1]))
	hour := int(core.Atoi32(arr[2]))
	minu := int(core.Atoi32(arr[3]))

	if hour >= 24 || minu >= 60 {
		val, ret = 0, false
		return
	}

	wday--
	val = wday*86400 + hour*3600 + minu*60
	ret = true
	return
}

func BeginOfDay(t time.Time) time.Time {
	y, M, d := t.Date()
	return time.Date(y, M, d, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	y, M, d := t.Date()
	return time.Date(y, M, d, 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1)
}

func IsSameDay(t1, t2 time.Time) bool {
	y1, M1, d1 := t1.Date()
	y2, M2, d2 := t2.Date()
	return y1 == y2 && M1 == M2 && d1 == d2
}

func DaySeconds() int {
	now := time.Now()
	return now.Hour()*3600 + now.Minute()*60 + now.Second()
}

// Week: from Monday to Sunday
func WeekSeconds() int {
	now := time.Now()
	day := int(now.Weekday())
	if day == 0 {
		day == 7
	}
	day--
	return day*86400 + now.Hour()*3600 + now.Minute()*60 + now.Second()
}
