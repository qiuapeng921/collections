// Package main 展示 Go Collections 库的完整用法示例
package main

import (
	"fmt"

	"github.com/qiuapeng921/collections"
)

// User 用户结构体示例
type User struct {
	ID     int
	Name   string
	Email  string
	Age    int
	Gender string
}

// Product 产品结构体示例
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func main() {
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║           Go Collections 库完整用法示例                      ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// 1. 创建集合
	creationExamples()

	// 2. 访问元素
	accessExamples()

	// 3. 过滤操作
	filterExamples()

	// 4. 转换操作
	transformExamples()

	// 5. 排序操作
	sortingExamples()

	// 6. 聚合操作
	aggregationExamples()

	// 7. 分组操作
	groupingExamples()

	// 8. 集合操作
	setOperationsExamples()

	// 9. MapCollection 用法
	mapCollectionExamples()

	// 10. 实际业务场景示例
	businessScenarioExamples()

	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println("                         示例完成！")
	fmt.Println("═══════════════════════════════════════════════════════════════")
}

// ============================================
// 1. 创建集合示例
// ============================================
func creationExamples() {
	fmt.Println("【1. 创建集合】")
	fmt.Println("─────────────────────────────────────────────────")

	// 从切片创建
	numbers := collections.New([]int{1, 2, 3, 4, 5})
	fmt.Printf("从切片创建: %v\n", numbers.All())

	// 从可变参数创建
	names := collections.Make("Alice", "Bob", "Charlie")
	fmt.Printf("从可变参数创建: %v\n", names.All())

	// 使用 Range 创建数字序列
	rangeNums := collections.Range(1, 10)
	fmt.Printf("Range(1, 10): %v\n", rangeNums.All())

	// 倒序范围
	reverseRange := collections.Range(5, 1)
	fmt.Printf("Range(5, 1): %v\n", reverseRange.All())

	// 使用 Times 生成
	squares := collections.Times(5, func(i int) int {
		return i * i
	})
	fmt.Printf("Times(5, i²): %v\n", squares.All())

	// 创建空集合
	empty := collections.Empty[string]()
	fmt.Printf("空集合是否为空: %v\n", empty.IsEmpty())

	fmt.Println()
}

// ============================================
// 2. 访问元素示例
// ============================================
func accessExamples() {
	fmt.Println("【2. 访问元素】")
	fmt.Println("─────────────────────────────────────────────────")

	numbers := collections.New([]int{10, 20, 30, 40, 50})

	// 基本访问
	fmt.Printf("集合元素: %v\n", numbers.All())
	fmt.Printf("元素个数: %d\n", numbers.Count())
	fmt.Printf("第一个元素: %d\n", numbers.First())
	fmt.Printf("最后一个元素: %d\n", numbers.Last())
	fmt.Printf("索引2的元素: %d\n", numbers.Get(2))

	// 带默认值的访问
	fmt.Printf("索引10的元素(默认99): %d\n", numbers.GetOr(10, 99))

	// 条件访问
	value, found := numbers.FirstWhere(func(n int) bool { return n > 25 })
	if found {
		fmt.Printf("第一个大于25的元素: %d\n", value)
	}

	lastValue, _ := numbers.LastWhere(func(n int) bool { return n < 35 })
	fmt.Printf("最后一个小于35的元素: %d\n", lastValue)

	fmt.Println()
}

// ============================================
// 3. 过滤操作示例
// ============================================
func filterExamples() {
	fmt.Println("【3. 过滤操作】")
	fmt.Println("─────────────────────────────────────────────────")

	numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("原始集合: %v\n", numbers.All())

	// Filter - 过滤偶数
	evens := numbers.Filter(func(n int) bool { return n%2 == 0 })
	fmt.Printf("Filter(偶数): %v\n", evens.All())

	// Reject - 排除小于5的
	rejected := numbers.Reject(func(n int) bool { return n < 5 })
	fmt.Printf("Reject(n < 5): %v\n", rejected.All())

	// Take 和 Skip
	fmt.Printf("Take(3): %v\n", numbers.Take(3).All())
	fmt.Printf("Take(-3): %v\n", numbers.Take(-3).All())
	fmt.Printf("Skip(7): %v\n", numbers.Skip(7).All())

	// Slice
	fmt.Printf("Slice(2, 4): %v\n", numbers.Slice(2, 4).All())

	// TakeWhile 和 SkipWhile
	fmt.Printf("TakeWhile(n < 5): %v\n", numbers.TakeWhile(func(n int) bool { return n < 5 }).All())
	fmt.Printf("SkipWhile(n < 5): %v\n", numbers.SkipWhile(func(n int) bool { return n < 5 }).All())

	// TakeUntil 和 SkipUntil
	fmt.Printf("TakeUntil(n > 5): %v\n", numbers.TakeUntil(func(n int) bool { return n > 5 }).All())
	fmt.Printf("SkipUntil(n > 5): %v\n", numbers.SkipUntil(func(n int) bool { return n > 5 }).All())

	fmt.Println()
}

// ============================================
// 4. 转换操作示例
// ============================================
func transformExamples() {
	fmt.Println("【4. 转换操作】")
	fmt.Println("─────────────────────────────────────────────────")

	numbers := collections.New([]int{1, 2, 3, 4, 5})
	fmt.Printf("原始集合: %v\n", numbers.All())

	// Map - 映射转换
	doubled := collections.Map(numbers, func(n int, i int) int {
		return n * 2
	})
	fmt.Printf("Map(n * 2): %v\n", doubled.All())

	// 转换为字符串
	strings := collections.Map(numbers, func(n int, i int) string {
		return fmt.Sprintf("数字%d", n)
	})
	fmt.Printf("Map(转字符串): %v\n", strings.All())

	// Reduce - 归约
	sum := collections.Reduce(numbers, func(acc int, n int, i int) int {
		return acc + n
	}, 0)
	fmt.Printf("Reduce(求和): %d\n", sum)

	product := collections.Reduce(numbers, func(acc int, n int, i int) int {
		return acc * n
	}, 1)
	fmt.Printf("Reduce(求积): %d\n", product)

	// FlatMap - 映射并展平
	expanded := collections.FlatMap(numbers, func(n int, i int) []int {
		return []int{n, n * 10}
	})
	fmt.Printf("FlatMap: %v\n", expanded.All())

	// Pluck - 提取字段
	users := collections.New([]User{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	})

	names := collections.Pluck(users, func(u User) string { return u.Name })
	fmt.Printf("Pluck(Name): %v\n", names.All())

	ages := collections.Pluck(users, func(u User) int { return u.Age })
	fmt.Printf("Pluck(Age): %v\n", ages.All())

	fmt.Println()
}

// ============================================
// 5. 排序操作示例
// ============================================
func sortingExamples() {
	fmt.Println("【5. 排序操作】")
	fmt.Println("─────────────────────────────────────────────────")

	numbers := collections.New([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3})
	fmt.Printf("原始集合: %v\n", numbers.All())

	// Sort - 升序排序
	sorted := collections.Sort(numbers)
	fmt.Printf("Sort(升序): %v\n", sorted.All())

	// SortDesc - 降序排序
	sortedDesc := collections.SortDesc(numbers)
	fmt.Printf("SortDesc(降序): %v\n", sortedDesc.All())

	// SortBy - 按字段排序
	users := collections.New([]User{
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 3, Name: "Charlie", Age: 20},
	})

	byAge := collections.SortBy(users, func(u User) int { return u.Age })
	fmt.Printf("SortBy(Age): ")
	byAge.Each(func(u User, i int) {
		fmt.Printf("%s(%d) ", u.Name, u.Age)
	})
	fmt.Println()

	byAgeDesc := collections.SortByDesc(users, func(u User) int { return u.Age })
	fmt.Printf("SortByDesc(Age): ")
	byAgeDesc.Each(func(u User, i int) {
		fmt.Printf("%s(%d) ", u.Name, u.Age)
	})
	fmt.Println()

	// Reverse - 反转
	nums := collections.New([]int{1, 2, 3, 4, 5})
	fmt.Printf("Reverse: %v\n", nums.Reverse().All())

	// Shuffle - 随机打乱
	fmt.Printf("Shuffle: %v\n", nums.Shuffle().All())

	fmt.Println()
}

// ============================================
// 6. 聚合操作示例
// ============================================
func aggregationExamples() {
	fmt.Println("【6. 聚合操作】")
	fmt.Println("─────────────────────────────────────────────────")

	numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("原始集合: %v\n", numbers.All())

	// 基本聚合
	fmt.Printf("Sum(求和): %d\n", collections.Sum(numbers))
	fmt.Printf("Avg(平均): %.2f\n", collections.Avg(numbers))
	fmt.Printf("Min(最小): %d\n", collections.Min(numbers))
	fmt.Printf("Max(最大): %d\n", collections.Max(numbers))
	fmt.Printf("Median(中位数): %.2f\n", collections.Median(numbers))

	// Mode - 众数
	repeated := collections.New([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})
	mode := collections.Mode(repeated)
	fmt.Printf("Mode(众数): %v\n", mode)

	// 按字段聚合
	products := collections.New([]Product{
		{ID: 1, Name: "苹果", Price: 5.5, Stock: 100},
		{ID: 2, Name: "香蕉", Price: 3.0, Stock: 150},
		{ID: 3, Name: "橙子", Price: 4.5, Stock: 80},
	})

	totalPrice := collections.SumBy(products, func(p Product) float64 { return p.Price })
	fmt.Printf("SumBy(Price): %.2f\n", totalPrice)

	avgPrice := collections.AvgBy(products, func(p Product) float64 { return p.Price })
	fmt.Printf("AvgBy(Price): %.2f\n", avgPrice)

	mostExpensive := collections.MaxBy(products, func(p Product) float64 { return p.Price })
	fmt.Printf("MaxBy(Price): %s (%.2f元)\n", mostExpensive.Name, mostExpensive.Price)

	cheapest := collections.MinBy(products, func(p Product) float64 { return p.Price })
	fmt.Printf("MinBy(Price): %s (%.2f元)\n", cheapest.Name, cheapest.Price)

	fmt.Println()
}

// ============================================
// 7. 分组操作示例
// ============================================
func groupingExamples() {
	fmt.Println("【7. 分组操作】")
	fmt.Println("─────────────────────────────────────────────────")

	users := collections.New([]User{
		{ID: 1, Name: "Alice", Age: 30, Gender: "F"},
		{ID: 2, Name: "Bob", Age: 25, Gender: "M"},
		{ID: 3, Name: "Charlie", Age: 35, Gender: "M"},
		{ID: 4, Name: "Diana", Age: 28, Gender: "F"},
		{ID: 5, Name: "Eve", Age: 30, Gender: "F"},
	})

	// GroupBy - 分组
	byGender := collections.GroupBy(users, func(u User) string { return u.Gender })
	fmt.Printf("GroupBy(Gender) - 女性: %d人, 男性: %d人\n",
		byGender.Get("F").Count(), byGender.Get("M").Count())

	// KeyBy - 按键索引
	byID := collections.KeyBy(users, func(u User) int { return u.ID })
	user3 := byID.Get(3)
	fmt.Printf("KeyBy(ID=3): %s\n", user3.Name)

	// CountBy - 计数
	ageCount := collections.CountBy(users, func(u User) int { return u.Age })
	fmt.Printf("CountBy(Age=30): %d人\n", ageCount.Get(30))

	// Partition - 分区
	numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	evens, odds := numbers.Partition(func(n int) bool { return n%2 == 0 })
	fmt.Printf("Partition - 偶数: %v, 奇数: %v\n", evens.All(), odds.All())

	// ChunkInto - 分块
	chunks := numbers.ChunkInto(3)
	fmt.Printf("ChunkInto(3): ")
	for i, chunk := range chunks {
		fmt.Printf("块%d%v ", i+1, chunk.All())
	}
	fmt.Println()

	// Split - 分成指定数量的组
	groups := numbers.Split(3)
	fmt.Printf("Split(3): ")
	for i, group := range groups {
		fmt.Printf("组%d%v ", i+1, group.All())
	}
	fmt.Println()

	// Sliding - 滑动窗口
	windows := numbers.Sliding(3)
	fmt.Printf("Sliding(3): ")
	for _, w := range windows[:3] {
		fmt.Printf("%v ", w.All())
	}
	fmt.Println("...")

	fmt.Println()
}

// ============================================
// 8. 集合操作示例
// ============================================
func setOperationsExamples() {
	fmt.Println("【8. 集合操作】")
	fmt.Println("─────────────────────────────────────────────────")

	a := collections.New([]int{1, 2, 3, 4, 5})
	b := collections.New([]int{4, 5, 6, 7, 8})
	fmt.Printf("集合A: %v\n", a.All())
	fmt.Printf("集合B: %v\n", b.All())

	// Diff - 差集
	diff := collections.Diff(a, b)
	fmt.Printf("Diff(A-B): %v\n", diff.All())

	// Intersect - 交集
	intersect := collections.Intersect(a, b)
	fmt.Printf("Intersect(A∩B): %v\n", intersect.All())

	// Merge - 合并
	merged := a.Merge(b)
	fmt.Printf("Merge(A∪B): %v\n", merged.All())

	// Unique - 去重
	withDups := collections.New([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})
	unique := collections.UniqueComparable(withDups)
	fmt.Printf("Unique: %v\n", unique.All())

	// Duplicates - 找重复
	duplicates := collections.Duplicates(withDups)
	fmt.Printf("Duplicates: %v\n", duplicates.All())

	// Contains 和 Every
	fmt.Printf("Contains(3): %v\n", collections.ContainsComparable(a, 3))
	fmt.Printf("Every(>0): %v\n", a.Every(func(n int) bool { return n > 0 }))

	// Nth - 每N个取一个
	nums := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("Nth(2): %v\n", nums.Nth(2).All())
	fmt.Printf("Nth(2, offset=1): %v\n", nums.Nth(2, 1).All())

	// ForPage - 分页
	page2 := nums.ForPage(2, 3)
	fmt.Printf("ForPage(page=2, perPage=3): %v\n", page2.All())

	// Pad - 填充
	short := collections.New([]int{1, 2, 3})
	fmt.Printf("Pad(6, 0): %v\n", short.Pad(6, 0).All())
	fmt.Printf("Pad(-6, 0): %v\n", short.Pad(-6, 0).All())

	fmt.Println()
}

// ============================================
// 9. MapCollection 示例
// ============================================
func mapCollectionExamples() {
	fmt.Println("【9. MapCollection 用法】")
	fmt.Println("─────────────────────────────────────────────────")

	// 创建 MapCollection
	inventory := collections.NewMap(map[string]int{
		"苹果": 100,
		"香蕉": 150,
		"橙子": 80,
		"葡萄": 200,
	})

	fmt.Printf("库存: %v\n", inventory.All())
	fmt.Printf("键: %v\n", inventory.Keys().All())
	fmt.Printf("值: %v\n", inventory.Values().All())

	// 基本操作
	fmt.Printf("Get('苹果'): %d\n", inventory.Get("苹果"))
	fmt.Printf("Has('苹果'): %v\n", inventory.Has("苹果"))
	fmt.Printf("Has('芒果'): %v\n", inventory.Has("芒果"))

	// Put 和 Forget
	inventory.Put("芒果", 50)
	fmt.Printf("Put('芒果', 50)后: %v\n", inventory.All())

	// Filter
	lowStock := inventory.Filter(func(v int, k string) bool { return v < 100 })
	fmt.Printf("Filter(库存<100): %v\n", lowStock.All())

	// Only 和 Except
	only := inventory.Only("苹果", "香蕉")
	fmt.Printf("Only('苹果', '香蕉'): %v\n", only.All())

	except := inventory.Except("苹果")
	fmt.Printf("Except('苹果'): %v\n", except.All())

	// Merge
	newItems := collections.NewMap(map[string]int{"西瓜": 30})
	merged := inventory.Merge(newItems)
	fmt.Printf("Merge后元素数: %d\n", merged.Count())

	// MapValues - 转换值
	doubled := collections.MapValues(inventory, func(v int, k string) int {
		return v * 2
	})
	fmt.Printf("MapValues(v*2): %v\n", doubled.All())

	fmt.Println()
}

// ============================================
// 10. 实际业务场景示例
// ============================================
func businessScenarioExamples() {
	fmt.Println("【10. 实际业务场景示例】")
	fmt.Println("─────────────────────────────────────────────────")

	// 场景1: 电商订单处理
	fmt.Println("\n>>> 场景1: 电商数据分析")
	orders := collections.New([]Product{
		{ID: 1, Name: "iPhone", Price: 6999, Stock: 50},
		{ID: 2, Name: "iPad", Price: 3999, Stock: 30},
		{ID: 3, Name: "MacBook", Price: 9999, Stock: 20},
		{ID: 4, Name: "AirPods", Price: 999, Stock: 100},
		{ID: 5, Name: "Apple Watch", Price: 2999, Stock: 40},
	})

	// 分析:
	// 1. 找出价格超过3000的高端产品
	highEnd := orders.Filter(func(p Product) bool { return p.Price > 3000 })
	fmt.Printf("高端产品(>3000): ")
	collections.Pluck(highEnd, func(p Product) string { return p.Name }).Each(func(s string, i int) {
		fmt.Printf("%s ", s)
	})
	fmt.Println()

	// 2. 计算总库存价值
	totalValue := collections.SumBy(orders, func(p Product) float64 {
		return p.Price * float64(p.Stock)
	})
	fmt.Printf("总库存价值: ¥%.2f\n", totalValue)

	// 3. 按价格分组
	priceGroups := collections.GroupBy(orders, func(p Product) string {
		if p.Price < 2000 {
			return "低价"
		} else if p.Price < 5000 {
			return "中价"
		}
		return "高价"
	})
	fmt.Printf("价格分布 - 低价:%d个, 中价:%d个, 高价:%d个\n",
		priceGroups.Get("低价").Count(),
		priceGroups.Get("中价").Count(),
		priceGroups.Get("高价").Count())

	// 场景2: 用户数据处理
	fmt.Println("\n>>> 场景2: 用户数据处理")
	users := collections.New([]User{
		{ID: 1, Name: "张三", Age: 28, Gender: "M", Email: "zhang@example.com"},
		{ID: 2, Name: "李四", Age: 35, Gender: "M", Email: "li@example.com"},
		{ID: 3, Name: "王五", Age: 42, Gender: "M", Email: "wang@example.com"},
		{ID: 4, Name: "赵六", Age: 25, Gender: "F", Email: "zhao@example.com"},
		{ID: 5, Name: "钱七", Age: 31, Gender: "F", Email: "qian@example.com"},
	})

	// 1. 找出30岁以上的用户
	over30 := users.Filter(func(u User) bool { return u.Age >= 30 })
	fmt.Printf("30岁以上用户: %d人\n", over30.Count())

	// 2. 按性别统计平均年龄
	maleUsers := users.Filter(func(u User) bool { return u.Gender == "M" })
	femaleUsers := users.Filter(func(u User) bool { return u.Gender == "F" })
	avgMaleAge := collections.AvgBy(maleUsers, func(u User) int { return u.Age })
	avgFemaleAge := collections.AvgBy(femaleUsers, func(u User) int { return u.Age })
	fmt.Printf("平均年龄 - 男: %.1f岁, 女: %.1f岁\n", avgMaleAge, avgFemaleAge)

	// 3. 创建用户快速查找表
	userLookup := collections.KeyBy(users, func(u User) int { return u.ID })
	fmt.Printf("用户ID=3: %s\n", userLookup.Get(3).Name)

	// 4. 获取所有邮箱列表
	emails := collections.Pluck(users, func(u User) string { return u.Email })
	emailList := collections.ImplodeStrings(emails, ", ")
	fmt.Printf("所有邮箱: %s\n", emailList)

	// 场景3: 数据转换和链式操作
	fmt.Println("\n>>> 场景3: 链式操作示例")
	result := users.
		Filter(func(u User) bool { return u.Age >= 25 }). // 筛选25岁以上
		Take(3)                                           // 取前3个

	fmt.Printf("链式操作结果: ")
	result.Each(func(u User, i int) {
		fmt.Printf("%s ", u.Name)
	})
	fmt.Println()

	// 场景4: 条件执行
	fmt.Println("\n>>> 场景4: 条件执行")
	numbers := collections.New([]int{1, 2, 3, 4, 5})
	shouldDouble := true

	processed := numbers.When(shouldDouble, func(c *collections.Collection[int]) *collections.Collection[int] {
		return collections.Map(c, func(n int, _ int) int { return n * 2 })
	})
	fmt.Printf("When(shouldDouble=true): %v\n", processed.All())

	// 场景5: 分页展示
	fmt.Println("\n>>> 场景5: 分页展示")
	allItems := collections.Range(1, 20)
	pageSize := 5
	totalPages := (allItems.Count() + pageSize - 1) / pageSize

	for page := 1; page <= totalPages; page++ {
		pageData := allItems.ForPage(page, pageSize)
		fmt.Printf("第%d页: %v\n", page, pageData.All())
	}

	fmt.Println()
}
