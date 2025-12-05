// Package main demonstrates collection sorting and aggregation methods
package main

import (
	"fmt"
	"strings"

	"github.com/qiuapeng921/collections"
)

// User 用户结构体
type User struct {
	ID    int
	Name  string
	Age   int
	Score float64
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
	fmt.Println("Collections 排序和聚合方法示例")
	fmt.Println(strings.Repeat("=", 60))

	numbers := collections.New([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5})
	floats := collections.New([]float64{1.5, 2.3, 3.7, 4.1, 5.9})
	users := collections.New([]User{
		{ID: 1, Name: "张三", Age: 25, Score: 85.5},
		{ID: 2, Name: "李四", Age: 30, Score: 92.0},
		{ID: 3, Name: "王五", Age: 35, Score: 78.5},
		{ID: 4, Name: "赵六", Age: 28, Score: 88.0},
	})
	products := collections.New([]Product{
		{ID: 1, Name: "iPhone", Price: 999.99, Category: "手机", Stock: 50},
		{ID: 2, Name: "MacBook", Price: 1999.99, Category: "电脑", Stock: 30},
		{ID: 3, Name: "iPad", Price: 799.99, Category: "平板", Stock: 100},
	})

	// ============== 排序 ==============
	fmt.Println("\n【1. Sort - 升序排序】")
	sorted := collections.Sort(numbers)
	fmt.Printf("升序: %v\n", sorted.All())

	fmt.Println("\n【2. SortDesc - 降序排序】")
	sortedDesc := collections.SortDesc(numbers)
	fmt.Printf("降序: %v\n", sortedDesc.All())

	fmt.Println("\n【3. SortBy - 按字段排序】")
	byAge := collections.SortBy(users, func(u User) int { return u.Age })
	byAge.Each(func(u User, i int) { fmt.Printf("  %s: %d岁\n", u.Name, u.Age) })

	fmt.Println("\n【4. SortByDesc - 按字段降序】")
	byScoreDesc := collections.SortByDesc(users, func(u User) float64 { return u.Score })
	byScoreDesc.Each(func(u User, i int) { fmt.Printf("  %s: %.1f分\n", u.Name, u.Score) })

	fmt.Println("\n【5. SortFunc - 自定义排序】")
	custom := numbers.SortFunc(func(a, b int) int { return b - a })
	fmt.Printf("自定义降序: %v\n", custom.All())

	// ============== 聚合 ==============
	fmt.Println("\n【6. Sum - 求和】")
	fmt.Printf("整数和: %d\n", collections.Sum(numbers))
	fmt.Printf("浮点和: %.2f\n", collections.Sum(floats))

	fmt.Println("\n【7. SumBy - 按字段求和】")
	totalStock := collections.SumBy(products, func(p Product) int { return p.Stock })
	fmt.Printf("总库存: %d\n", totalStock)

	fmt.Println("\n【8. Avg - 平均值】")
	fmt.Printf("平均值: %.2f\n", collections.Avg(numbers))
	fmt.Printf("浮点平均: %.2f\n", collections.Avg(floats))

	fmt.Println("\n【9. AvgBy - 按字段平均】")
	avgAge := collections.AvgBy(users, func(u User) int { return u.Age })
	fmt.Printf("平均年龄: %.1f\n", avgAge)

	fmt.Println("\n【10. Min/Max - 最小/最大值】")
	fmt.Printf("最小: %d, 最大: %d\n", collections.Min(numbers), collections.Max(numbers))

	fmt.Println("\n【11. MinBy/MaxBy - 按字段最小/最大】")
	youngest := collections.MinBy(users, func(u User) int { return u.Age })
	oldest := collections.MaxBy(users, func(u User) int { return u.Age })
	fmt.Printf("最年轻: %s(%d), 最年长: %s(%d)\n", youngest.Name, youngest.Age, oldest.Name, oldest.Age)

	fmt.Println("\n【12. Median - 中位数】")
	medianNums := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("中位数: %.1f\n", collections.Median(medianNums))

	fmt.Println("\n【13. Mode - 众数】")
	mode := collections.Mode(numbers)
	fmt.Printf("众数: %v\n", mode)

	// ============== 集合操作 ==============
	fmt.Println("\n【14. Diff - 差集】")
	a := collections.Make(1, 2, 3, 4, 5)
	b := collections.Make(3, 4, 5, 6, 7)
	fmt.Printf("A-B差集: %v\n", collections.Diff(a, b).All())

	fmt.Println("\n【15. Intersect - 交集】")
	fmt.Printf("A∩B交集: %v\n", collections.Intersect(a, b).All())

	fmt.Println("\n【16. Duplicates - 重复元素】")
	dups := collections.Duplicates(numbers)
	fmt.Printf("重复元素: %v\n", dups.All())

	fmt.Println("\n【17. UniqueComparable - 去重】")
	unique := collections.UniqueComparable(numbers)
	fmt.Printf("去重后: %v\n", unique.All())

	fmt.Println("\n【18. IndexOf/LastIndexOf - 查找索引】")
	fmt.Printf("5的第一次出现: %d\n", collections.IndexOf(numbers, 5))
	fmt.Printf("5的最后出现: %d\n", collections.LastIndexOf(numbers, 5))

	fmt.Println("\n【19. ContainsComparable - 包含检查】")
	fmt.Printf("包含5: %v\n", collections.ContainsComparable(numbers, 5))
	fmt.Printf("包含100: %v\n", collections.ContainsComparable(numbers, 100))
}
