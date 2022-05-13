package helpers

import (
	"fmt"
	"strings"
	"time"
)

func IyzipayTimeString(time time.Time) string {
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
}

func Append(key string, value interface{}) string {
	if value == nil {
		return ""
	}
	if value == "" {
		return ""
	}
	return fmt.Sprintf("%v=%v,", key, value)
}

func AppendArray(key string, value interface{}) string {
	if value == nil {
		return ""
	}
	if value == "" {
		return ""
	}
	return fmt.Sprintf("%v=%v,", key, value)
}

func AppendPrice(key, value string) string {
	if value == "" {
		return ""
	}
	value = formatPrice(value)
	return fmt.Sprintf("%v=%v,", key, value)
}

func formatPrice(price string) string {
	if !strings.Contains(price, ".") {
		return fmt.Sprintf("%s.0", price)
	}
	reversePrice := reverseString(price)

	startIndex := 0
	for i := 0; i < len(reversePrice); i++ {
		if string(reversePrice[i]) == "0" {
			startIndex = i + 1
		} else {
			if string(reversePrice[i]) == "." {
				reversePrice = "0" + reversePrice
				break
			}
			break
		}
	}

	return reverseString(reversePrice[startIndex:])
}

func GetRequestString(text string) string {
	text = removeTrailingComma(text)
	text = appendPrefix(text)
	return text
}

func removeTrailingComma(text string) string {
	if text == "" {
		return ""
	}
	return text[:len(text)-1]
}

func appendPrefix(text string) string {
	return "[" + text + "]"
}

func reverseString(s string) string {
	a := []byte(s)
	for i, j := 0, len(s)-1; i < j; i++ {
		a[i], a[j] = a[j], a[i]
		j--
	}
	return string(a)
}
