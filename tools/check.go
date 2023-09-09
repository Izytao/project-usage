package tools

import "regexp"

// CheckHanZiValid 检查某个汉字有效性
func CheckHanZiValid(hanZi string) bool {
	matched, err := regexp.MatchString("^[\u4e00-\u9fa5]$", hanZi)
	if err == nil && matched {
		return true
	}
	return false
}

// IsValidPassword4OR3 密码匹配 数字，大、小写，特殊符号4选3
func IsValidPassword4OR3(password string) bool {
	// 过滤掉这4类字符以外的密码串，直接判断不合法
	re := regexp.MustCompile(`^[0-9a-zA-Z~!@#$%^&*()\[\]{};:'",.<>/?\-_+=\\|]{8,}$`)
	if !re.MatchString(password) {
		return false
	}
	var level = 1
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*()\[\]{};:'",.<>/?\-_+=\\|]`}

	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, password)
		if match {
			level++
		}
	}
	if level < 3 {
		return false
	}

	return true
}

// IsValidAccount 账号匹配 20位大小写和数字
func IsValidAccount(account string) bool {
	re := regexp.MustCompile(`^[0-9a-zA-Z]{1,20}$`)
	return re.MatchString(account)
}
