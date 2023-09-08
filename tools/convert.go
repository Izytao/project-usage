package tools

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"strconv"
	"strings"
)

var (
	FORMAT_GBK  = "GB18030"
	FORMAT_UTF8 = "UTF8"
)

// ConvertByte2String GBK格式转换为UTF-8
func ConvertByte2String(byte []byte, charset string) string {
	var str string
	switch charset {
	case FORMAT_GBK:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case FORMAT_UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

// StrToInt 类型转化 string  to int
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// StrToUint64 类型转化 string  to uint64
func StrToUint64(str string) uint64 {
	i, _ := strconv.ParseUint(str, 0, 64)
	return i
}

// StrToFloat64 类型转化 string  to float64
func StrToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

// IntToString 类型转化 int to string
func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}

// Int64ToString 类型转化 int64 to string
func Int64ToString(i int64) string {
	return fmt.Sprintf("%d", i)
}

// Uint64ToString 类型转化 uint64 to string
func Uint64ToString(i uint64) string {
	return fmt.Sprintf("%d", i)
}

// Uint32ToString 类型转化 uint32 to string
func Uint32ToString(i uint32) string {
	return fmt.Sprintf("%d", i)
}

// InterfaceToString 类型转换inerface to string
func InterfaceToString(data interface{}) string {
	return fmt.Sprintf("%s", data)
}

// StringToInt64Arr 将[id,id,id]字符串转换成id数组
func StringToInt64Arr(src string) ([]int64, error) {
	list := make([]int64, 0)
	dec := json.NewDecoder(strings.NewReader(src))
	dec.UseNumber()
	err := dec.Decode(&list)
	return list, err
}

// Float64Decimal 保留几位小数
func Float64Decimal(f float64, decimal int) string {
	format := "%." + fmt.Sprint(decimal) + "f"
	return fmt.Sprintf(format, f)
}
