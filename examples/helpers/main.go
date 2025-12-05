// Package main demonstrates helper functions and utilities
package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/qiuapeng921/collections"
)

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Helper 帮助函数和工具示例")
	fmt.Println(strings.Repeat("=", 60))

	// ============== 切片工具 ==============
	fmt.Println("\n【1. Head - 获取第一个元素】")
	nums := []int{1, 2, 3, 4, 5}
	strs := []string{"a", "b", "c", "d"}
	fmt.Printf("Head(nums): %d\n", collections.Head(nums))
	fmt.Printf("Head(strs): %s\n", collections.Head(strs))
	fmt.Printf("Head(空切片): %d\n", collections.Head([]int{}))

	fmt.Println("\n【2. Tail - 获取除第一个外的所有元素】")
	fmt.Printf("Tail(nums): %v\n", collections.Tail(nums))
	fmt.Printf("Tail(strs): %v\n", collections.Tail(strs))

	fmt.Println("\n【3. Init - 获取除最后一个外的所有元素】")
	fmt.Printf("Init(nums): %v\n", collections.Init(nums))
	fmt.Printf("Init(strs): %v\n", collections.Init(strs))

	fmt.Println("\n【4. LastItem - 获取最后一个元素】")
	fmt.Printf("LastItem(nums): %d\n", collections.LastItem(nums))
	fmt.Printf("LastItem(strs): %s\n", collections.LastItem(strs))

	// ============== 值检查 ==============
	fmt.Println("\n【5. Blank - 检查是否为空值】")
	fmt.Printf("Blank(0): %v\n", collections.Blank(0))
	fmt.Printf("Blank(1): %v\n", collections.Blank(1))
	fmt.Printf("Blank(\"\"): %v\n", collections.Blank(""))
	fmt.Printf("Blank(\"hello\"): %v\n", collections.Blank("hello"))

	fmt.Println("\n【6. Filled - 检查是否非空值】")
	fmt.Printf("Filled(0): %v\n", collections.Filled(0))
	fmt.Printf("Filled(1): %v\n", collections.Filled(1))
	fmt.Printf("Filled(\"\"): %v\n", collections.Filled(""))
	fmt.Printf("Filled(\"hello\"): %v\n", collections.Filled("hello"))

	// ============== Optional 类型 ==============
	fmt.Println("\n【7. Optional - 安全的可空值处理】")

	// 创建有值的Optional
	some := collections.Some(42)
	fmt.Printf("Some(42).HasValue(): %v\n", some.HasValue())
	fmt.Printf("Some(42).Get(): %v\n", some.Get())
	fmt.Printf("Some(42).GetOr(0): %v\n", some.GetOr(0))

	// 创建空的Optional
	none := collections.None[int]()
	fmt.Printf("None.HasValue(): %v\n", none.HasValue())
	fmt.Printf("None.GetOr(0): %v\n", none.GetOr(0))

	// Map操作
	doubled := some.Map(func(v int) int { return v * 2 })
	fmt.Printf("Some(42).Map(*2): %v\n", doubled.Get())

	// Filter操作
	filtered := some.Filter(func(v int) bool { return v > 50 })
	fmt.Printf("Some(42).Filter(>50).HasValue(): %v\n", filtered.HasValue())

	filtered2 := some.Filter(func(v int) bool { return v > 40 })
	fmt.Printf("Some(42).Filter(>40).HasValue(): %v\n", filtered2.HasValue())

	// ============== 函数工具 ==============
	fmt.Println("\n【8. Value - 返回值本身】")
	fmt.Printf("Value(100): %d\n", collections.Value(100))
	fmt.Printf("Value(\"hello\"): %s\n", collections.Value("hello"))

	fmt.Println("\n【9. Identity - 恒等函数】")
	fmt.Printf("Identity(42): %d\n", collections.Identity(42))
	fmt.Printf("Identity(\"test\"): %s\n", collections.Identity("test"))

	fmt.Println("\n【10. With - 返回值】")
	fmt.Printf("With(123): %d\n", collections.With(123))

	fmt.Println("\n【11. Tap - 执行回调并返回原值】")
	result := collections.Tap(100, func(v int) {
		fmt.Printf("  Tap回调中的值: %d\n", v)
	})
	fmt.Printf("Tap返回值: %d\n", result)

	// ============== 执行控制 ==============
	fmt.Println("\n【12. Once - 只执行一次的函数】")
	counter := 0
	onceFn := collections.Once(func() int {
		counter++
		return counter
	})
	fmt.Printf("第1次调用: %d\n", onceFn())
	fmt.Printf("第2次调用: %d\n", onceFn())
	fmt.Printf("第3次调用: %d\n", onceFn())
	fmt.Printf("counter实际值: %d\n", counter)

	fmt.Println("\n【13. Retry - 重试机制】")
	attempts := 0
	result2, err := collections.Retry(3, func() (string, error) {
		attempts++
		if attempts < 3 {
			return "", errors.New("failed")
		}
		return "success", nil
	})
	fmt.Printf("Retry结果: %s, 尝试次数: %d\n", result2, attempts)
	if err != nil {
		fmt.Printf("Retry错误: %v\n", err)
	}

	// 失败的情况
	_, err = collections.Retry(2, func() (int, error) {
		return 0, errors.New("always fails")
	})
	if err != nil {
		fmt.Printf("Retry失败: %v\n", err)
	}

	fmt.Println("\n【14. Rescue - 异常捕获】")
	rescued := collections.Rescue(func() int {
		panic("something went wrong")
	}, -1)
	fmt.Printf("Rescue捕获异常后返回默认值: %d\n", rescued)

	normal := collections.Rescue(func() int {
		return 42
	}, -1)
	fmt.Printf("Rescue正常执行: %d\n", normal)

	fmt.Println("\n【15. Transform - 条件转换】")
	// 非空值时转换
	transformed := collections.Transform(10, func(v int) string {
		return fmt.Sprintf("值是: %d", v)
	})
	fmt.Printf("Transform(10): %s\n", transformed)

	// 空值时返回默认值
	transformed2 := collections.Transform(0, func(v int) string {
		return fmt.Sprintf("值是: %d", v)
	}, "默认")
	fmt.Printf("Transform(0): %s\n", transformed2)

	// ============== 数据访问 ==============
	fmt.Println("\n【16. DataGet - 从嵌套结构获取数据】")
	nested := map[string]any{
		"user": map[string]any{
			"name": "张三",
			"age":  25,
		},
	}
	fmt.Printf("DataGet(user.name): %v\n", collections.DataGet(nested, "user.name"))
	fmt.Printf("DataGet(user.phone, 默认): %v\n", collections.DataGet(nested, "user.phone", "未设置"))

	fmt.Println("\n【17. DataSet - 设置嵌套数据】")
	data := map[string]any{}
	collections.DataSet(data, "config.debug", true)
	fmt.Printf("DataSet后: %v\n", data)

	fmt.Println("\n【18. DataForget - 删除嵌套数据】")
	forgetData := map[string]any{"a": 1, "b": 2}
	collections.DataForget(forgetData, "a")
	fmt.Printf("DataForget后: %v\n", forgetData)

	// ============== 异常控制 ==============
	fmt.Println("\n【19. ThrowIf - 条件抛出异常】")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("ThrowIf捕获: %v\n", r)
			}
		}()
		collections.ThrowIf(true, "条件为true，抛出异常")
	}()

	fmt.Println("\n【20. ThrowUnless - 条件不满足抛出异常】")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("ThrowUnless捕获: %v\n", r)
			}
		}()
		collections.ThrowUnless(false, "条件为false，抛出异常")
	}()

	// ============== 综合示例 ==============
	fmt.Println("\n【21. 综合示例：安全的用户查询】")

	findUser := func(id int) collections.Optional[string] {
		users := map[int]string{1: "张三", 2: "李四", 3: "王五"}
		if name, ok := users[id]; ok {
			return collections.Some(name)
		}
		return collections.None[string]()
	}

	// 查找存在的用户
	user1 := findUser(1)
	fmt.Printf("查找ID=1: %s\n", user1.GetOr("未找到"))

	// 查找不存在的用户
	user4 := findUser(4)
	fmt.Printf("查找ID=4: %s\n", user4.GetOr("未找到"))

	// 链式操作
	greeting := findUser(2).
		Map(func(name string) string { return "你好, " + name }).
		GetOr("用户不存在")
	fmt.Printf("问候语: %s\n", greeting)
}
