package week1

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("下标超出范围")

// DeleteAt 删除指定位置的元素
// 如果下标不是合法的下标，返回 ErrIndexOutOfRange
func DeleteAt[T any](s []T, idx int) ([]T, error) {
	fmt.Printf("%p, %v\n", s, s)
	if idx < 0 || idx >= len(s) {
		return nil, ErrIndexOutOfRange
	}
	result := s[:idx]
	fmt.Printf("%p, %v\n", result, result)
	result = append(result, s[idx+1:]...)
	fmt.Printf("%p, %v\n", result, result)
	fmt.Printf("%d, %d\n", len(result), cap(result))
	return result, nil
}
