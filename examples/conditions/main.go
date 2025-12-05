// Package main demonstrates condition and iteration methods
package main

import (
	"fmt"
	"strings"

	"github.com/qiuapeng921/collections"
)

// User 用户结构体
type User struct {
	ID     int
	Name   string
	Age    int
	Active bool
	Score  float64
}

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("条件操作和迭代方法示例")
	fmt.Println(strings.Repeat("=", 60))

	numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	users := collections.New([]User{
		{ID: 1, Name: "张三", Age: 25, Active: true, Score: 85.5},
		{ID: 2, Name: "李四", Age: 30, Active: true, Score: 92.0},
		{ID: 3, Name: "王五", Age: 35, Active: false, Score: 78.5},
	})

	// ============== 条件检查 ==============
	fmt.Println("\n【1. Contains - 检查是否存在匹配元素】")
	hasEven := numbers.Contains(func(n int) bool { return n%2 == 0 })
	fmt.Printf("是否有偶数: %v\n", hasEven)

	hasOlder := users.Contains(func(u User) bool { return u.Age > 40 })
	fmt.Printf("是否有>40岁: %v\n", hasOlder)

	hasActive := users.Contains(func(u User) bool { return u.Active })
	fmt.Printf("是否有活跃用户: %v\n", hasActive)

	fmt.Println("\n【2. Some - Contains的别名】")
	hasHighScore := users.Some(func(u User) bool { return u.Score > 90 })
	fmt.Printf("是否有高分(>90): %v\n", hasHighScore)

	fmt.Println("\n【3. Every - 检查所有元素是否满足条件】")
	allPositive := numbers.Every(func(n int) bool { return n > 0 })
	fmt.Printf("是否全为正数: %v\n", allPositive)

	allActive := users.Every(func(u User) bool { return u.Active })
	fmt.Printf("是否全为活跃: %v\n", allActive)

	allAdult := users.Every(func(u User) bool { return u.Age >= 18 })
	fmt.Printf("是否全为成年: %v\n", allAdult)

	// ============== 条件执行 ==============
	fmt.Println("\n【4. When - 条件为true时执行】")
	shouldDouble := true
	result := numbers.When(shouldDouble, func(c *collections.Collection[int]) *collections.Collection[int] {
		return collections.Map(c, func(n int, i int) int { return n * 2 })
	})
	fmt.Printf("When(true)后: %v\n", result.All())

	shouldDouble = false
	result = numbers.When(shouldDouble, func(c *collections.Collection[int]) *collections.Collection[int] {
		return collections.Map(c, func(n int, i int) int { return n * 2 })
	})
	fmt.Printf("When(false)后: %v\n", result.All())

	fmt.Println("\n【5. Unless - 条件为false时执行】")
	isDebug := false
	result = numbers.Unless(isDebug, func(c *collections.Collection[int]) *collections.Collection[int] {
		return c.Filter(func(n int) bool { return n > 5 })
	})
	fmt.Printf("Unless(false)后: %v\n", result.All())

	fmt.Println("\n【6. WhenEmpty - 集合为空时执行】")
	empty := collections.Empty[int]()
	filled := empty.WhenEmpty(func(c *collections.Collection[int]) *collections.Collection[int] {
		c.Push(1, 2, 3)
		return c
	})
	fmt.Printf("WhenEmpty后: %v\n", filled.All())

	fmt.Println("\n【7. WhenNotEmpty - 集合非空时执行】")
	result = numbers.WhenNotEmpty(func(c *collections.Collection[int]) *collections.Collection[int] {
		return c.Take(3)
	})
	fmt.Printf("WhenNotEmpty后: %v\n", result.All())

	// ============== 迭代操作 ==============
	fmt.Println("\n【8. Each - 遍历每个元素】")
	fmt.Print("遍历: ")
	numbers.Take(5).Each(func(n int, i int) {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("[%d]=%d", i, n)
	})
	fmt.Println()

	fmt.Println("遍历用户:")
	users.Each(func(u User, i int) {
		fmt.Printf("  %d. %s (%d岁)\n", i+1, u.Name, u.Age)
	})

	fmt.Println("\n【9. EachSpread - 遍历(可中断)】")
	fmt.Print("遍历到>5停止: ")
	numbers.EachSpread(func(n int, i int) bool {
		fmt.Printf("%d ", n)
		return n < 5 // 返回false时中断
	})
	fmt.Println()

	fmt.Println("\n【10. Tap - 调试链】")
	result = numbers.
		Filter(func(n int) bool { return n%2 == 0 }).
		Tap(func(c *collections.Collection[int]) {
			fmt.Printf("  过滤偶数后: %v\n", c.All())
		}).
		Take(3).
		Tap(func(c *collections.Collection[int]) {
			fmt.Printf("  取前3个后: %v\n", c.All())
		})
	fmt.Printf("最终结果: %v\n", result.All())

	fmt.Println("\n【11. Pipe - 管道处理】")
	sum := collections.Pipe(numbers, func(c *collections.Collection[int]) int {
		return collections.Sum(c)
	})
	fmt.Printf("管道求和: %d\n", sum)

	avg := collections.Pipe(numbers, func(c *collections.Collection[int]) float64 {
		return collections.Avg(c)
	})
	fmt.Printf("管道平均: %.1f\n", avg)

	// ============== 修改操作 ==============
	fmt.Println("\n【12. Push - 添加到末尾】")
	list := collections.Make(1, 2, 3)
	list.Push(4, 5)
	fmt.Printf("Push后: %v\n", list.All())

	fmt.Println("\n【13. Pop - 弹出末尾元素】")
	popped := list.Pop()
	fmt.Printf("Pop: %d, 剩余: %v\n", popped, list.All())

	fmt.Println("\n【14. Prepend - 添加到开头】")
	list.Prepend(0, -1)
	fmt.Printf("Prepend后: %v\n", list.All())

	fmt.Println("\n【15. Shift - 弹出开头元素】")
	shifted := list.Shift()
	fmt.Printf("Shift: %d, 剩余: %v\n", shifted, list.All())

	fmt.Println("\n【16. Put - 设置指定索引】")
	list.Put(1, 100)
	fmt.Printf("Put(1,100)后: %v\n", list.All())

	fmt.Println("\n【17. Forget - 删除指定索引】")
	list.Forget(1)
	fmt.Printf("Forget(1)后: %v\n", list.All())

	fmt.Println("\n【18. Pull - 获取并删除】")
	pulled := list.Pull(0)
	fmt.Printf("Pull(0): %d, 剩余: %v\n", pulled, list.All())

	// ============== 字符串操作 ==============
	fmt.Println("\n【19. Implode/Join - 连接字符串】")
	words := collections.Make("Hello", "World", "Go")

	joined := collections.ImplodeStrings(words, " ")
	fmt.Printf("ImplodeStrings: %s\n", joined)

	joinedCustom := collections.JoinStrings(words, ", ", " and ")
	fmt.Printf("JoinStrings带最后分隔符: %s\n", joinedCustom)

	// 自定义转换
	userNames := collections.Implode(users, func(u User) string { return u.Name }, ", ")
	fmt.Printf("用户名连接: %s\n", userNames)

	// ============== 序列化 ==============
	fmt.Println("\n【20. ToJSON - 转为JSON】")
	jsonBytes, _ := numbers.Take(5).ToJSON()
	fmt.Printf("ToJSON: %s\n", string(jsonBytes))

	jsonStr := users.ToJSONString()
	fmt.Printf("ToJSONString: %s\n", jsonStr)

	fmt.Println("\n【21. String - 字符串表示】")
	fmt.Printf("String: %s\n", numbers.Take(5).String())

	// ============== 调试 ==============
	fmt.Println("\n【22. Dump - 调试输出】")
	collections.Make(1, 2, 3).Dump()
}
