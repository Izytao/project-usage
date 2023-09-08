package tools

import "regexp"

// CheckHanZiValid 检查某个汉字有效性
func CheckHanZiValid(hanZi string) bool {
	matched, err := regexp.MatchString("^[\u4e00-\u9fa5]{1}$", hanZi)
	if err == nil && matched {
		return true
	}
	return false
}
