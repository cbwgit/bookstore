package date_utils

import (
	"time"
)

const(
	apiDateLayout = "1995-10-27T14:20:23Z"
)

func GetNow() time.Time {

	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)

}
