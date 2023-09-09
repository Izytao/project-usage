package tools

import "regexp"

// IsMobile 判断userAgent是不是手机
func IsMobile(userAgent string) bool {
	mobileRe, _ := regexp.Compile("(?i:Mobile|iPod|iPhone|Android|Opera Mini|BlackBerry|webOS|UCWEB|Blazer|PSP)")
	str := mobileRe.FindString(userAgent)
	if str != "" {
		return true
	}
	return false
}
