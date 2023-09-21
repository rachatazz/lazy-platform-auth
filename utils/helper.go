package utils

func Includes(cond interface{}, args ...interface{}) bool {
	result := false
	for _, arg := range args {
		if cond == arg {
			result = true
			break
		}
	}
	return result
}
