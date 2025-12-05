// Package main demonstrates collection filter methods
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
	Email    string
	Active   bool
	Role     string
	Score    float64
	JoinYear int
}

// Order 订单结构体
type Order struct {
	ID       int
	UserID   int
	Product  string
	Amount   float64
	Status   string
	Quantity int
}

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Collections 过滤方法示例")
	fmt.Println(strings.Repeat("=", 60))

	// 准备测试数据
	numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	users := collections.New([]User{
		{ID: 1, Name: "张三", Age: 25, Email: "zhangsan@example.com", Active: true, Role: "admin", Score: 85.5, JoinYear: 2020},
		{ID: 2, Name: "李四", Age: 30, Email: "lisi@example.com", Active: false, Role: "user", Score: 92.0, JoinYear: 2019},
		{ID: 3, Name: "王五", Age: 35, Email: "wangwu@example.com", Active: true, Role: "user", Score: 78.5, JoinYear: 2021},
		{ID: 4, Name: "赵六", Age: 28, Email: "zhaoliu@example.com", Active: true, Role: "moderator", Score: 88.0, JoinYear: 2020},
		{ID: 5, Name: "钱七", Age: 40, Email: "qianqi@example.com", Active: false, Role: "user", Score: 95.5, JoinYear: 2018},
		{ID: 6, Name: "孙八", Age: 22, Email: "sunba@example.com", Active: true, Role: "user", Score: 70.0, JoinYear: 2022},
	})
	orders := collections.New([]Order{
		{ID: 1001, UserID: 1, Product: "iPhone", Amount: 999.99, Status: "completed", Quantity: 1},
		{ID: 1002, UserID: 2, Product: "MacBook", Amount: 1999.99, Status: "pending", Quantity: 1},
		{ID: 1003, UserID: 1, Product: "AirPods", Amount: 199.99, Status: "completed", Quantity: 2},
		{ID: 1004, UserID: 3, Product: "iPad", Amount: 799.99, Status: "cancelled", Quantity: 1},
		{ID: 1005, UserID: 4, Product: "Apple Watch", Amount: 399.99, Status: "completed", Quantity: 1},
		{ID: 1006, UserID: 5, Product: "HomePod", Amount: 299.99, Status: "pending", Quantity: 3},
	})

	// ============================================================
	// 1. Filter - 过滤满足条件的元素
	// ============================================================
	fmt.Println("\n【1. Filter - 过滤满足条件的元素】")

	// 过滤偶数
	evens := numbers.Filter(func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("偶数: %v\n", evens.All())

	// 过滤3的倍数
	multiplesOf3 := numbers.Filter(func(n int) bool {
		return n%3 == 0
	})
	fmt.Printf("3的倍数: %v\n", multiplesOf3.All())

	// 过滤活跃用户
	activeUsers := users.Filter(func(u User) bool {
		return u.Active
	})
	fmt.Println("活跃用户:")
	activeUsers.Each(func(u User, i int) {
		fmt.Printf("  - %s (年龄: %d)\n", u.Name, u.Age)
	})

	// 过滤年龄大于30的用户
	olderUsers := users.Filter(func(u User) bool {
		return u.Age > 30
	})
	fmt.Println("年龄>30的用户:")
	olderUsers.Each(func(u User, i int) {
		fmt.Printf("  - %s (年龄: %d)\n", u.Name, u.Age)
	})

	// 过滤已完成的订单
	completedOrders := orders.Filter(func(o Order) bool {
		return o.Status == "completed"
	})
	fmt.Println("已完成订单:")
	completedOrders.Each(func(o Order, i int) {
		fmt.Printf("  - #%d: %s (%.2f)\n", o.ID, o.Product, o.Amount)
	})

	// ============================================================
	// 2. Reject - 排除满足条件的元素
	// ============================================================
	fmt.Println("\n【2. Reject - 排除满足条件的元素】")

	// 排除小于5的数字
	rejectSmall := numbers.Reject(func(n int) bool {
		return n < 5
	})
	fmt.Printf("排除<5: %v\n", rejectSmall.All())

	// 排除非活跃用户(等同于获取活跃用户)
	rejectInactive := users.Reject(func(u User) bool {
		return !u.Active
	})
	fmt.Printf("排除非活跃用户数量: %d\n", rejectInactive.Count())

	// 排除已取消的订单
	notCancelled := orders.Reject(func(o Order) bool {
		return o.Status == "cancelled"
	})
	fmt.Printf("未取消订单数量: %d\n", notCancelled.Count())

	// ============================================================
	// 3. Take / Take(负数) - 获取前N个/后N个元素
	// ============================================================
	fmt.Println("\n【3. Take - 获取前N个/后N个元素】")

	// 取前5个
	first5 := numbers.Take(5)
	fmt.Printf("前5个: %v\n", first5.All())

	// 取后3个 (负数)
	last3 := numbers.Take(-3)
	fmt.Printf("后3个: %v\n", last3.All())

	// 取前2个用户
	first2Users := users.Take(2)
	fmt.Println("前2个用户:")
	first2Users.Each(func(u User, i int) {
		fmt.Printf("  - %s\n", u.Name)
	})

	// ============================================================
	// 4. Skip - 跳过前N个元素
	// ============================================================
	fmt.Println("\n【4. Skip - 跳过前N个元素】")

	// 跳过前5个
	skip5 := numbers.Skip(5)
	fmt.Printf("跳过前5个: %v\n", skip5.All())

	// 跳过前10个
	skip10 := numbers.Skip(10)
	fmt.Printf("跳过前10个: %v\n", skip10.All())

	// 跳过前3个用户
	skip3Users := users.Skip(3)
	fmt.Println("跳过前3个用户后:")
	skip3Users.Each(func(u User, i int) {
		fmt.Printf("  - %s\n", u.Name)
	})

	// ============================================================
	// 5. TakeWhile - 获取元素直到条件不满足
	// ============================================================
	fmt.Println("\n【5. TakeWhile - 获取元素直到条件不满足】")

	nums := collections.New([]int{1, 2, 3, 4, 5, 4, 3, 2, 1})

	// 取元素直到>=5
	takeWhileLt5 := nums.TakeWhile(func(n int) bool {
		return n < 5
	})
	fmt.Printf("TakeWhile n<5: %v\n", takeWhileLt5.All())

	// ============================================================
	// 6. TakeUntil - 获取元素直到条件满足
	// ============================================================
	fmt.Println("\n【6. TakeUntil - 获取元素直到条件满足】")

	// 取元素直到>=4
	takeUntil4 := nums.TakeUntil(func(n int) bool {
		return n >= 4
	})
	fmt.Printf("TakeUntil n>=4: %v\n", takeUntil4.All())

	// ============================================================
	// 7. SkipWhile - 跳过元素直到条件不满足
	// ============================================================
	fmt.Println("\n【7. SkipWhile - 跳过元素直到条件不满足】")

	// 跳过小于4的元素
	skipWhileLt4 := nums.SkipWhile(func(n int) bool {
		return n < 4
	})
	fmt.Printf("SkipWhile n<4: %v\n", skipWhileLt4.All())

	// ============================================================
	// 8. SkipUntil - 跳过元素直到条件满足
	// ============================================================
	fmt.Println("\n【8. SkipUntil - 跳过元素直到条件满足】")

	// 跳过直到>=4
	skipUntil4 := nums.SkipUntil(func(n int) bool {
		return n >= 4
	})
	fmt.Printf("SkipUntil n>=4: %v\n", skipUntil4.All())

	// ============================================================
	// 9. Slice - 切片操作
	// ============================================================
	fmt.Println("\n【9. Slice - 切片操作】")

	// 从索引2开始取4个元素
	sliced := numbers.Slice(2, 4)
	fmt.Printf("Slice(2, 4): %v\n", sliced.All())

	// 从索引5开始取到末尾
	slicedToEnd := numbers.Slice(5)
	fmt.Printf("Slice(5): %v\n", slicedToEnd.All())

	// 负索引（从末尾开始）
	slicedNeg := numbers.Slice(-3)
	fmt.Printf("Slice(-3): %v\n", slicedNeg.All())

	// ============================================================
	// 10. Partition - 分区（分成满足和不满足条件两组）
	// ============================================================
	fmt.Println("\n【10. Partition - 分区】")

	// 按奇偶分区
	evensP, oddsP := numbers.Partition(func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("偶数分区: %v\n", evensP.All())
	fmt.Printf("奇数分区: %v\n", oddsP.All())

	// 按活跃状态分区用户
	activeP, inactiveP := users.Partition(func(u User) bool {
		return u.Active
	})
	fmt.Printf("活跃用户数: %d\n", activeP.Count())
	fmt.Printf("非活跃用户数: %d\n", inactiveP.Count())

	// ============================================================
	// 11. ChunkInto - 分块
	// ============================================================
	fmt.Println("\n【11. ChunkInto - 分块】")

	// 每3个一组
	chunks := numbers.ChunkInto(3)
	fmt.Println("每3个分块:")
	for i, chunk := range chunks {
		fmt.Printf("  块%d: %v\n", i+1, chunk.All())
	}

	// ============================================================
	// 12. Split - 分成指定数量的组
	// ============================================================
	fmt.Println("\n【12. Split - 分成指定数量的组】")

	// 分成3组
	groups := numbers.Split(3)
	fmt.Println("分成3组:")
	for i, group := range groups {
		fmt.Printf("  组%d: %v\n", i+1, group.All())
	}

	// ============================================================
	// 13. Nth - 每隔N个取一个
	// ============================================================
	fmt.Println("\n【13. Nth - 每隔N个取一个】")

	// 每隔2个取一个
	every2nd := numbers.Nth(2)
	fmt.Printf("每隔2个(从0开始): %v\n", every2nd.All())

	// 每隔3个取一个，从索引1开始
	every3rdFrom1 := numbers.Nth(3, 1)
	fmt.Printf("每隔3个(从1开始): %v\n", every3rdFrom1.All())

	// ============================================================
	// 14. ForPage - 分页
	// ============================================================
	fmt.Println("\n【14. ForPage - 分页】")

	// 每页5条，获取各页
	fmt.Println("每页5条:")
	for page := 1; page <= 3; page++ {
		pageData := numbers.ForPage(page, 5)
		fmt.Printf("  第%d页: %v\n", page, pageData.All())
	}

	// ============================================================
	// 15. Unique - 去重
	// ============================================================
	fmt.Println("\n【15. Unique - 去重】")

	duplicates := collections.New([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})
	unique := collections.UniqueComparable(duplicates)
	fmt.Printf("去重前: %v\n", duplicates.All())
	fmt.Printf("去重后: %v\n", unique.All())

	// 按属性去重用户
	usersWithDups := collections.New([]User{
		{ID: 1, Name: "张三", Role: "admin"},
		{ID: 2, Name: "李四", Role: "user"},
		{ID: 3, Name: "王五", Role: "admin"},
		{ID: 4, Name: "赵六", Role: "user"},
	})
	uniqueByRole := usersWithDups.Unique(func(u User) string {
		return u.Role
	})
	fmt.Println("按角色去重:")
	uniqueByRole.Each(func(u User, i int) {
		fmt.Printf("  - %s (%s)\n", u.Name, u.Role)
	})

	// ============================================================
	// 16. 组合使用示例
	// ============================================================
	fmt.Println("\n【16. 组合使用示例】")

	// 获取活跃用户中，分数大于80的前2名用户
	result := users.
		Filter(func(u User) bool {
			return u.Active
		}).
		Filter(func(u User) bool {
			return u.Score > 80
		}).
		Take(2)

	fmt.Println("活跃且分数>80的前2名用户:")
	result.Each(func(u User, i int) {
		fmt.Printf("  - %s (分数: %.1f)\n", u.Name, u.Score)
	})
}
