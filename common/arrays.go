package common

import (
	"strconv"
	"strings"
)

func To2DArray(arrayStr string) [][]int {
	if arrayStr == "[]" {
		return [][]int{}
	}
	// example: [[0,7],[0,8],[6,1]]
	// 移除开头和结尾的方括号
	s := arrayStr[1 : len(arrayStr)-1]
	// 将子数组之间的分隔符替换为分号，以便切分出子数组
	s = strings.ReplaceAll(s, "],", "];")
	subStrs := strings.Split(s, ";")
	twoDim := make([][]int, 0)
	for _, subStr := range subStrs {
		itemStrs := strings.Split(subStr[1:len(subStr)-1], ",")
		first, _ := strconv.Atoi(itemStrs[0])
		second, _ := strconv.Atoi(itemStrs[1])
		twoDim = append(twoDim, []int{first, second})
	}
	return twoDim
}
