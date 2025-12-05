// Package main demonstrates collection transform methods
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
	Email string
	Tags  []string
}

// UserDTO 用户DTO
type UserDTO struct {
	ID       int
	FullName string
	IsAdult  bool
}

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Collections 转换方法示例")
	fmt.Println(strings.Repeat("=", 60))

	numbers := collections.New([]int{1, 2, 3, 4, 5})
	users := collections.New([]User{
		{ID: 1, Name: "张三", Age: 25, Tags: []string{"vip", "new"}},
		{ID: 2, Name: "李四", Age: 17, Tags: []string{"regular"}},
		{ID: 3, Name: "王五", Age: 35, Tags: []string{"vip"}},
	})

	// 1. Map - 映射转换
	fmt.Println("\n【1. Map - 映射转换】")
	doubled := collections.Map(numbers, func(n int, i int) int { return n * 2 })
	fmt.Printf("乘以2: %v\n", doubled.All())

	userDTOs := collections.Map(users, func(u User, i int) UserDTO {
		return UserDTO{ID: u.ID, FullName: "用户:" + u.Name, IsAdult: u.Age >= 18}
	})
	userDTOs.Each(func(dto UserDTO, i int) {
		fmt.Printf("  DTO: %+v\n", dto)
	})

	// 2. FlatMap - 映射并展平
	fmt.Println("\n【2. FlatMap - 映射并展平】")
	expanded := collections.FlatMap(numbers, func(n int, i int) []int {
		return []int{n, n * 10}
	})
	fmt.Printf("展开: %v\n", expanded.All())

	allTags := collections.FlatMap(users, func(u User, i int) []string { return u.Tags })
	fmt.Printf("所有标签: %v\n", allTags.All())

	// 3. Reduce - 归约
	fmt.Println("\n【3. Reduce - 归约】")
	sum := collections.Reduce(numbers, func(acc, n, i int) int { return acc + n }, 0)
	fmt.Printf("求和: %d\n", sum)

	// 4. Pluck - 提取属性
	fmt.Println("\n【4. Pluck - 提取属性】")
	names := collections.Pluck(users, func(u User) string { return u.Name })
	fmt.Printf("名字: %v\n", names.All())

	// 5. Transform - 原地转换
	fmt.Println("\n【5. Transform - 原地转换】")
	nums := collections.Make(1, 2, 3)
	nums.Transform(func(n, i int) int { return n * 10 })
	fmt.Printf("原地转换: %v\n", nums.All())

	// 6. Reverse - 反转
	fmt.Println("\n【6. Reverse - 反转】")
	fmt.Printf("反转: %v\n", numbers.Reverse().All())

	// 7. Merge - 合并
	fmt.Println("\n【7. Merge - 合并】")
	merged := collections.Make(1, 2).Merge(collections.Make(3, 4))
	fmt.Printf("合并: %v\n", merged.All())

	// 8. Clone - 克隆
	fmt.Println("\n【8. Clone - 克隆】")
	original := collections.Make(1, 2, 3)
	cloned := original.Clone()
	cloned.Push(4)
	fmt.Printf("原始: %v, 克隆: %v\n", original.All(), cloned.All())

	// 9. Pad - 填充
	fmt.Println("\n【9. Pad - 填充】")
	fmt.Printf("右填充: %v\n", collections.Make(1, 2).Pad(5, 0).All())
	fmt.Printf("左填充: %v\n", collections.Make(1, 2).Pad(-5, 0).All())

	// 10. Sliding - 滑动窗口
	fmt.Println("\n【10. Sliding - 滑动窗口】")
	for i, w := range numbers.Sliding(3) {
		fmt.Printf("  窗口%d: %v\n", i+1, w.All())
	}
}
