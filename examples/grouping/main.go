// Package main demonstrates collection grouping methods
package main

import (
	"fmt"
	"strings"

	"github.com/qiuapeng921/collections"
)

// User 用户结构体
type User struct {
	ID       int
	Name     string
	Age      int
	Gender   string
	Role     string
	JoinYear int
}

// Order 订单结构体
type Order struct {
	ID       int
	UserID   int
	Product  string
	Amount   float64
	Status   string
	Category string
}

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Collections 分组方法示例")
	fmt.Println(strings.Repeat("=", 60))

	users := collections.New([]User{
		{ID: 1, Name: "张三", Age: 25, Gender: "M", Role: "admin", JoinYear: 2020},
		{ID: 2, Name: "李四", Age: 30, Gender: "M", Role: "user", JoinYear: 2019},
		{ID: 3, Name: "王芳", Age: 28, Gender: "F", Role: "user", JoinYear: 2021},
		{ID: 4, Name: "赵丽", Age: 35, Gender: "F", Role: "moderator", JoinYear: 2020},
		{ID: 5, Name: "钱七", Age: 40, Gender: "M", Role: "user", JoinYear: 2018},
		{ID: 6, Name: "孙红", Age: 22, Gender: "F", Role: "user", JoinYear: 2022},
	})

	orders := collections.New([]Order{
		{ID: 1001, UserID: 1, Product: "iPhone", Amount: 999.99, Status: "completed", Category: "电子"},
		{ID: 1002, UserID: 2, Product: "MacBook", Amount: 1999.99, Status: "pending", Category: "电子"},
		{ID: 1003, UserID: 1, Product: "AirPods", Amount: 199.99, Status: "completed", Category: "配件"},
		{ID: 1004, UserID: 3, Product: "iPad", Amount: 799.99, Status: "cancelled", Category: "电子"},
		{ID: 1005, UserID: 4, Product: "键盘", Amount: 129.99, Status: "completed", Category: "配件"},
		{ID: 1006, UserID: 5, Product: "鼠标", Amount: 79.99, Status: "pending", Category: "配件"},
	})

	// 1. GroupBy - 分组
	fmt.Println("\n【1. GroupBy - 分组】")

	// 按性别分组
	byGender := collections.GroupBy(users, func(u User) string { return u.Gender })
	fmt.Println("按性别分组:")
	byGender.Each(func(k string, v *collections.Collection[User]) {
		gender := "男"
		if k == "F" {
			gender = "女"
		}
		fmt.Printf("  %s (%d人): ", gender, v.Count())
		v.Each(func(u User, i int) {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(u.Name)
		})
		fmt.Println()
	})

	// 按角色分组
	byRole := collections.GroupBy(users, func(u User) string { return u.Role })
	fmt.Println("按角色分组:")
	byRole.Each(func(k string, v *collections.Collection[User]) {
		fmt.Printf("  %s: %d人\n", k, v.Count())
	})

	// 按入职年份分组
	byYear := collections.GroupBy(users, func(u User) int { return u.JoinYear })
	fmt.Println("按入职年份分组:")
	byYear.Each(func(k int, v *collections.Collection[User]) {
		fmt.Printf("  %d年: %d人\n", k, v.Count())
	})

	// 按订单状态分组
	byStatus := collections.GroupBy(orders, func(o Order) string { return o.Status })
	fmt.Println("按订单状态分组:")
	byStatus.Each(func(k string, v *collections.Collection[Order]) {
		fmt.Printf("  %s: %d单\n", k, v.Count())
	})

	// 2. KeyBy - 按键索引
	fmt.Println("\n【2. KeyBy - 按键索引】")

	// 按ID建立索引
	byID := collections.KeyBy(users, func(u User) int { return u.ID })
	fmt.Println("按ID索引:")
	byID.Each(func(k int, v User) {
		fmt.Printf("  ID %d -> %s\n", k, v.Name)
	})

	// 按名字索引
	byName := collections.KeyBy(users, func(u User) string { return u.Name })
	user := byName.Get("张三")
	fmt.Printf("查找张三: ID=%d, 年龄=%d\n", user.ID, user.Age)

	// 3. CountBy - 计数
	fmt.Println("\n【3. CountBy - 计数】")

	// 按性别计数
	genderCount := collections.CountBy(users, func(u User) string { return u.Gender })
	fmt.Println("性别统计:")
	genderCount.Each(func(k string, v int) {
		fmt.Printf("  %s: %d人\n", k, v)
	})

	// 按角色计数
	roleCount := collections.CountBy(users, func(u User) string { return u.Role })
	fmt.Println("角色统计:")
	roleCount.Each(func(k string, v int) {
		fmt.Printf("  %s: %d人\n", k, v)
	})

	// 按订单类别计数
	categoryCount := collections.CountBy(orders, func(o Order) string { return o.Category })
	fmt.Println("订单类别统计:")
	categoryCount.Each(func(k string, v int) {
		fmt.Printf("  %s: %d单\n", k, v)
	})

	// 4. Partition - 分区
	fmt.Println("\n【4. Partition - 分区】")

	// 按年龄分区
	young, old := users.Partition(func(u User) bool { return u.Age < 30 })
	fmt.Printf("年轻人(<30): %d人\n", young.Count())
	fmt.Printf("中年人(>=30): %d人\n", old.Count())

	// 按订单金额分区
	bigOrders, smallOrders := orders.Partition(func(o Order) bool { return o.Amount > 500 })
	fmt.Printf("大订单(>500): %d单\n", bigOrders.Count())
	fmt.Printf("小订单(<=500): %d单\n", smallOrders.Count())

	// 5. ChunkInto - 分块
	fmt.Println("\n【5. ChunkInto - 分块】")
	nums := collections.Range(1, 10)
	chunks := nums.ChunkInto(3)
	fmt.Println("每3个分块:")
	for i, chunk := range chunks {
		fmt.Printf("  块%d: %v\n", i+1, chunk.All())
	}

	// 6. Split - 分成N组
	fmt.Println("\n【6. Split - 分成N组】")
	groups := nums.Split(3)
	fmt.Println("分成3组:")
	for i, group := range groups {
		fmt.Printf("  组%d: %v\n", i+1, group.All())
	}

	// 7. MapWithKeys - 创建键值对映射
	fmt.Println("\n【7. MapWithKeys - 创建键值对映射】")
	idToEmail := collections.MapWithKeys(users, func(u User, i int) (int, string) {
		return u.ID, fmt.Sprintf("%s@example.com", u.Name)
	})
	fmt.Println("ID到邮箱映射:")
	idToEmail.Each(func(k int, v string) {
		fmt.Printf("  %d -> %s\n", k, v)
	})

	// 8. MapToDictionary - 映射到字典
	fmt.Println("\n【8. MapToDictionary - 映射到字典(可重复key)】")
	roleUsers := collections.MapToDictionary(users, func(u User, i int) (string, string) {
		return u.Role, u.Name
	})
	fmt.Println("角色到用户名映射:")
	roleUsers.Each(func(k string, v []string) {
		fmt.Printf("  %s: %v\n", k, v)
	})
}
