package stream

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/fx"
	"testing"
)

// TestDistinct 测试流式处理去重
func TestDistinct(t *testing.T) {
	fx.Just(1, 2, 3, 3, 4, 5, 6).Distinct(func(item interface{}) interface{} {
		return item
	}).ForEach(func(item interface{}) {
		fmt.Println(item)
	})

	fmt.Println()
	fx.Just(1, 2, 3, 3, 4, 5, 5).Distinct(func(item interface{}) interface{} {
		uid := item.(int)
		if uid > 3 {
			return 4
		}
		return item
	}).ForEach(func(item interface{}) {
		fmt.Println(item)
	})
}

// map练习
func TestInternalStream_Map(t *testing.T) {
	var result int
	fx.Just(1, 2, 3, 4, 5, 2, 2, 2, 2, 2, 2).Map(func(item interface{}) interface{} {
		return item
	}).Reduce(func(pipe <-chan interface{}) (interface{}, error) {
		for item := range pipe {
			result += item.(int)
		}
		return result, nil
	})
	fmt.Println(result)
}
