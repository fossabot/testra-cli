package datetime

import (
	"time"
	"github.com/testra-tech/testra-cli/internal/constants/strs"
)

const DATE_TIME_FORMAT = "_2 Jan 2006 15:04:05"

func ConvertToDateTime(timeInMillis *int64) string {
	if timeInMillis == nil {
		return strs.EmptyStr
	} else {
		return time.Unix(*timeInMillis/1000, 0).Format(DATE_TIME_FORMAT)
	}
}