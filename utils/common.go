package utils

import "strconv"

func StringToInt64(s string) (result int64, err error) {
	return strconv.ParseInt(s, 10, 64)
}
func Int64ToString(i int64) (result string) {
	return strconv.FormatInt(i, 10)
}
