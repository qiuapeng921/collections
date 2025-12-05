// Package main demonstrates collection creation methods
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
	Email  string
	Active bool
}

// Product 产品结构体
type Product struct {
	ID       int
	Name     string
	Price    float64
	Category string
	Stock    int
}

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Collections 创建方法示例")
	fmt.Println(strings.Repeat("=", 60))

	// ============================================================
	// 1. New - 从切片创建集合
	// ============================================================
	fmt.Println("\n【1. New - 从切片创建集合】")

	// 整数切片
	intSlice := []int{1, 2, 3, 4, 5}
	intCol := collections.New(intSlice)
	fmt.Printf("整数集合: %v\n", intCol.All())

	// 字符串切片
	strSlice := []string{"apple", "banana", "cherry"}
	strCol := collections.New(strSlice)
	fmt.Printf("字符串集合: %v\n", strCol.All())

	// 结构体切片
	users := []User{
		{ID: 1, Name: "Alice", Age: 25, Email: "alice@example.com", Active: true},
		{ID: 2, Name: "Bob", Age: 30, Email: "bob@example.com", Active: false},
		{ID: 3, Name: "Charlie", Age: 35, Email: "charlie@example.com", Active: true},
	}
	userCol := collections.New(users)
	fmt.Printf("用户集合数量: %d\n", userCol.Count())

	// ============================================================
	// 2. Make - 从可变参数创建集合
	// ============================================================
	fmt.Println("\n【2. Make - 从可变参数创建集合】")

	// 直接传入多个值
	numCol := collections.Make(10, 20, 30, 40, 50)
	fmt.Printf("数字集合: %v\n", numCol.All())

	names := collections.Make("张三", "李四", "王五", "赵六")
	fmt.Printf("姓名集合: %v\n", names.All())

	// 结构体
	products := collections.Make(
		Product{ID: 1, Name: "iPhone", Price: 999.99, Category: "Electronics", Stock: 50},
		Product{ID: 2, Name: "MacBook", Price: 1999.99, Category: "Electronics", Stock: 30},
		Product{ID: 3, Name: "iPad", Price: 799.99, Category: "Electronics", Stock: 100},
	)
	fmt.Printf("产品集合数量: %d\n", products.Count())

	// ============================================================
	// 3. Range - 创建数字序列
	// ============================================================
	fmt.Println("\n【3. Range - 创建数字序列】")

	// 递增序列
	ascending := collections.Range(1, 10)
	fmt.Printf("1到10: %v\n", ascending.All())

	// 递减序列
	descending := collections.Range(10, 1)
	fmt.Printf("10到1: %v\n", descending.All())

	// 负数序列
	negative := collections.Range(-5, 5)
	fmt.Printf("-5到5: %v\n", negative.All())

	// ============================================================
	// 4. Times - 重复生成
	// ============================================================
	fmt.Println("\n【4. Times - 重复生成】")

	// 生成平方数
	squares := collections.Times(5, func(i int) int {
		return i * i
	})
	fmt.Printf("平方数: %v\n", squares.All())

	// 生成用户列表
	generatedUsers := collections.Times(3, func(i int) User {
		return User{
			ID:     i,
			Name:   fmt.Sprintf("用户%d", i),
			Age:    20 + i,
			Email:  fmt.Sprintf("user%d@example.com", i),
			Active: i%2 == 1,
		}
	})
	fmt.Println("生成的用户:")
	generatedUsers.Each(func(u User, idx int) {
		fmt.Printf("  %d. %s (年龄: %d)\n", idx+1, u.Name, u.Age)
	})

	// ============================================================
	// 5. Empty - 创建空集合
	// ============================================================
	fmt.Println("\n【5. Empty - 创建空集合】")

	emptyInts := collections.Empty[int]()
	fmt.Printf("空整数集合是否为空: %v\n", emptyInts.IsEmpty())

	emptyUsers := collections.Empty[User]()
	fmt.Printf("空用户集合数量: %d\n", emptyUsers.Count())

	// ============================================================
	// 6. Collect - 从可变参数创建（helpers.go）
	// ============================================================
	fmt.Println("\n【6. Collect - 从可变参数创建】")

	collected := collections.Collect(100, 200, 300, 400)
	fmt.Printf("Collect创建: %v\n", collected.All())

	// ============================================================
	// 7. CollectSlice - 从切片创建（helpers.go）
	// ============================================================
	fmt.Println("\n【7. CollectSlice - 从切片创建】")

	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	floatCol := collections.CollectSlice(floatSlice)
	fmt.Printf("浮点数集合: %v\n", floatCol.All())
}
