package utils

import "strings"

/**
 * @Author: Hao_pp
 * @Data: 2022-7-30 10:05
 * @Desc: StringUtil
 */

const (
	invaild_chars = " -。‘·；】【’，/、\\|——+="
) //非法字符集

//JudgeStrVaild 判断字符串是否合法
func JudgeStrVaild(s string, len_limit int) bool {

	if len(s) >= len_limit {
		return false
	}

	for _, k := range invaild_chars {
		if find := strings.Contains(s, string(k)); find {
			return false
		}
	}

	return true
}
