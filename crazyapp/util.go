/**
 * Auth :   liubo
 * Date :   2021/6/29 14:18
 * Comment:
 */

package crazyapp

import "regexp"

// 以字母开头，只包含字母和数字的字符串
var IsLetter = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]+$`).MatchString

