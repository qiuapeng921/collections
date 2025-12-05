// Package main demonstrates collection access methods
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

// Order 订单结构体
type Order struct {
	ID       int
	UserID   int
	Product  string
	Amount   float64
	Quantity int
}

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Collections 访问方法示例")
	fmt.Println(strings.Repeat("=", 60))

	// 准备测试数据
	numbers := collections.New([]int{10, 20, 30, 40, 50})
	names := collections.Make("Alice", "Bob", "Charlie", "David", "Eve")
	users := collections.New([]User{
		{ID: 1, Name: "张三", Age: 25, Email: "zhangsan@example.com", Active: true},
		{ID: 2, Name: "李四", Age: 30, Email: "lisi@example.com", Active: false},
		{ID: 3, Name: "王五", Age: 35, Email: "wangwu@example.com", Active: true},
		{ID: 4, Name: "赵六", Age: 28, Email: "zhaoliu@example.com", Active: true},
		{ID: 5, Name: "钱七", Age: 40, Email: "qianqi@example.com", Active: false},
	})
	orders := collections.New([]Order{
		{ID: 1001, UserID: 1, Product: "iPhone", Amount: 999.99, Quantity: 1},
		{ID: 1002, UserID: 2, Product: "MacBook", Amount: 1999.99, Quantity: 1},
		{ID: 1003, UserID: 1, Product: "AirPods", Amount: 199.99, Quantity: 2},
		{ID: 1004, UserID: 3, Product: "iPad", Amount: 799.99, Quantity: 1},
	})

	// ============================================================
	// 1. All - 获取所有元素
	// ============================================================
	fmt.Println("\n【1. All - 获取所有元素】")
	fmt.Printf("所有数字: %v\n", numbers.All())
	fmt.Printf("所有名字: %v\n", names.All())

	// ============================================================
	// 2. Count - 获取元素数量
	// ============================================================
	fmt.Println("\n【2. Count - 获取元素数量】")
	fmt.Printf("数字数量: %d\n", numbers.Count())
	fmt.Printf("用户数量: %d\n", users.Count())
	fmt.Printf("订单数量: %d\n", orders.Count())

	// ============================================================
	// 3. IsEmpty / IsNotEmpty - 判断是否为空
	// ============================================================
	fmt.Println("\n【3. IsEmpty / IsNotEmpty - 判断是否为空】")
	empty := collections.Empty[int]()
	fmt.Printf("空集合 IsEmpty: %v\n", empty.IsEmpty())
	fmt.Printf("空集合 IsNotEmpty: %v\n", empty.IsNotEmpty())
	fmt.Printf("数字集合 IsEmpty: %v\n", numbers.IsEmpty())
	fmt.Printf("数字集合 IsNotEmpty: %v\n", numbers.IsNotEmpty())

	// ============================================================
	// 4. First / FirstOr - 获取第一个元素
	// ============================================================
	fmt.Println("\n【4. First / FirstOr - 获取第一个元素】")
	fmt.Printf("第一个数字: %d\n", numbers.First())
	fmt.Printf("第一个名字: %s\n", names.First())

	firstUser := users.First()
	fmt.Printf("第一个用户: %s (年龄: %d)\n", firstUser.Name, firstUser.Age)

	// FirstOr - 空集合时返回默认值
	emptyNums := collections.Empty[int]()
	fmt.Printf("空集合FirstOr(默认值100): %d\n", emptyNums.FirstOr(100))

	// ============================================================
	// 5. Last / LastOr - 获取最后一个元素
	// ============================================================
	fmt.Println("\n【5. Last / LastOr - 获取最后一个元素】")
	fmt.Printf("最后一个数字: %d\n", numbers.Last())
	fmt.Printf("最后一个名字: %s\n", names.Last())

	lastUser := users.Last()
	fmt.Printf("最后一个用户: %s (年龄: %d)\n", lastUser.Name, lastUser.Age)

	// LastOr - 空集合时返回默认值
	fmt.Printf("空集合LastOr(默认值999): %d\n", emptyNums.LastOr(999))

	// ============================================================
	// 6. Get / GetOr - 按索引获取元素
	// ============================================================
	fmt.Println("\n【6. Get / GetOr - 按索引获取元素】")
	fmt.Printf("索引0的数字: %d\n", numbers.Get(0))
	fmt.Printf("索引2的数字: %d\n", numbers.Get(2))
	fmt.Printf("索引4的数字: %d\n", numbers.Get(4))

	// GetOr - 索引超出范围时返回默认值
	fmt.Printf("索引10(超出范围): %d\n", numbers.GetOr(10, -1))

	user2 := users.Get(1)
	fmt.Printf("索引1的用户: %s\n", user2.Name)

	// ============================================================
	// 7. FirstWhere - 条件获取第一个匹配元素
	// ============================================================
	fmt.Println("\n【7. FirstWhere - 条件获取第一个匹配元素】")

	// 找第一个大于25的数字
	num, found := numbers.FirstWhere(func(n int) bool {
		return n > 25
	})
	if found {
		fmt.Printf("第一个大于25的数字: %d\n", num)
	}

	// 找第一个活跃用户
	activeUser, found := users.FirstWhere(func(u User) bool {
		return u.Active
	})
	if found {
		fmt.Printf("第一个活跃用户: %s\n", activeUser.Name)
	}

	// 找第一个年龄大于30的用户
	older, found := users.FirstWhere(func(u User) bool {
		return u.Age > 30
	})
	if found {
		fmt.Printf("第一个年龄>30的用户: %s (年龄: %d)\n", older.Name, older.Age)
	}

	// 找第一个金额大于1000的订单
	bigOrder, found := orders.FirstWhere(func(o Order) bool {
		return o.Amount > 1000
	})
	if found {
		fmt.Printf("第一个金额>1000的订单: %s (金额: %.2f)\n", bigOrder.Product, bigOrder.Amount)
	}

	// ============================================================
	// 8. LastWhere - 条件获取最后一个匹配元素
	// ============================================================
	fmt.Println("\n【8. LastWhere - 条件获取最后一个匹配元素】")

	// 找最后一个偶数
	evenNums := collections.New([]int{1, 2, 4, 5, 6, 8, 9})
	lastEven, found := evenNums.LastWhere(func(n int) bool {
		return n%2 == 0
	})
	if found {
		fmt.Printf("最后一个偶数: %d\n", lastEven)
	}

	// 找最后一个活跃用户
	lastActive, found := users.LastWhere(func(u User) bool {
		return u.Active
	})
	if found {
		fmt.Printf("最后一个活跃用户: %s\n", lastActive.Name)
	}

	// ============================================================
	// 9. ContainsOneItem - 判断是否只有一个元素
	// ============================================================
	fmt.Println("\n【9. ContainsOneItem - 判断是否只有一个元素】")
	single := collections.Make(42)
	fmt.Printf("单元素集合 ContainsOneItem: %v\n", single.ContainsOneItem())
	fmt.Printf("多元素集合 ContainsOneItem: %v\n", numbers.ContainsOneItem())
	fmt.Printf("空集合 ContainsOneItem: %v\n", empty.ContainsOneItem())

	// ============================================================
	// 10. Search - 查找元素索引
	// ============================================================
	fmt.Println("\n【10. Search - 查找元素索引】")

	// 查找数字30的索引
	idx := numbers.Search(func(n int) bool {
		return n == 30
	})
	fmt.Printf("数字30的索引: %d\n", idx)

	// 查找用户"王五"的索引
	userIdx := users.Search(func(u User) bool {
		return u.Name == "王五"
	})
	fmt.Printf("用户'王五'的索引: %d\n", userIdx)

	// 查找不存在的元素
	notFound := numbers.Search(func(n int) bool {
		return n == 999
	})
	fmt.Printf("不存在的元素索引: %d\n", notFound) // 返回 -1

	// ============================================================
	// 11. Random / RandomN - 随机获取元素
	// ============================================================
	fmt.Println("\n【11. Random / RandomN - 随机获取元素】")
	fmt.Printf("随机一个数字: %d\n", numbers.Random())
	fmt.Printf("随机一个名字: %s\n", names.Random())

	randomUsers := users.RandomN(2)
	fmt.Println("随机2个用户:")
	randomUsers.Each(func(u User, i int) {
		fmt.Printf("  - %s\n", u.Name)
	})

	// ============================================================
	// 12. 安全访问方法 (OrFail 系列)
	// ============================================================
	fmt.Println("\n【12. 安全访问方法 (OrFail 系列)】")

	// FirstOrFail
	first, err := numbers.FirstOrFail()
	if err != nil {
		fmt.Printf("FirstOrFail 错误: %v\n", err)
	} else {
		fmt.Printf("FirstOrFail: %d\n", first)
	}

	// 空集合会返回错误
	_, err = empty.FirstOrFail()
	if err != nil {
		fmt.Printf("空集合 FirstOrFail 错误: %v\n", err)
	}

	// GetOrFail
	val, err := numbers.GetOrFail(2)
	if err != nil {
		fmt.Printf("GetOrFail 错误: %v\n", err)
	} else {
		fmt.Printf("GetOrFail(2): %d\n", val)
	}

	// RandomOrFail
	randVal, err := numbers.RandomOrFail()
	if err != nil {
		fmt.Printf("RandomOrFail 错误: %v\n", err)
	} else {
		fmt.Printf("RandomOrFail: %d\n", randVal)
	}
}
