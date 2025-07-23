package regexp

import "regexp"

// StringValidation
//
//	@Description: 字符串验证
//	@param str
//	@param compile
//	@return bool
func StringValidation(str, compile string) bool {
	re := regexp.MustCompile(compile)

	return re.MatchString(str)
}
