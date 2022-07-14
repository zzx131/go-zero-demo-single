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

type User struct {
	Id   int
	Name string
	Age  int
}

// 收集id转成切片
func TestListStream_Map(t *testing.T) {
	userList := make([]interface{}, 10)

	var userIdList []int

	userList = append(userList, User{Id: 1, Age: 12, Name: "张三"}, User{Id: 2, Age: 13, Name: "李四"}, User{Id: 3, Age: 10, Name: "赵六"})
	fx.Just(userList...).Map(func(item interface{}) interface{} {
		us := item.(User)
		return us.Id
	}).Reduce(func(pipe <-chan interface{}) (interface{}, error) {
		for item := range pipe {
			userIdList = append(userIdList, item.(int))
		}
		return userIdList, nil
	})
	fmt.Println(userIdList)
}
