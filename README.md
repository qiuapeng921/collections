# Go Collections

[![Go Version](https://img.shields.io/badge/Go-1.22+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

一个功能强大、类型安全的 Go 集合处理库，灵感来自 Laravel 的 Collection 类。使用 Go 1.22+ 泛型特性实现，提供流畅的链式 API。

## ✨ 特性

- 🎯 **泛型支持** - 使用 Go 1.22+ 泛型，完整的类型安全
- 🔗 **链式调用** - 流畅的 API，支持方法链式调用
- 🛡️ **不可变操作** - 大多数操作返回新集合，不修改原集合
- 📦 **丰富的 API** - 70+ 方法，覆盖过滤、映射、排序、聚合等
- 🗺️ **Map 集合** - 支持键值对集合，类似 PHP 关联数组
- 📄 **JSON 支持** - 内置 JSON 序列化/反序列化

## 📦 安装

```bash
go get github.com/qiuapeng921/collections
```

## 🚀 快速开始

```go
package main

import (
    "fmt"
    "github.com/qiuapeng921/collections"
)

func main() {
    // 创建集合
    numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

    // 过滤偶数并取前3个
    result := numbers.
        Filter(func(n int) bool { return n%2 == 0 }).
        Take(3)

    fmt.Println(result.All()) // [2, 4, 6]

    // 计算总和
    sum := collections.Sum(numbers)
    fmt.Println("Sum:", sum) // Sum: 55
}
```

---

## 📖 完整用法指南

### 1. 创建集合

#### 从切片创建
```go
// 从切片创建
numbers := collections.New([]int{1, 2, 3, 4, 5})

// 从可变参数创建
names := collections.Make("Alice", "Bob", "Charlie")
```

#### 使用 Range 创建数字序列
```go
// 创建 1 到 10 的序列
nums := collections.Range(1, 10)
fmt.Println(nums.All()) // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

// 创建倒序序列
reverse := collections.Range(5, 1)
fmt.Println(reverse.All()) // [5, 4, 3, 2, 1]
```

#### 使用 Times 重复生成
```go
// 生成 5 个平方数
squares := collections.Times(5, func(i int) int {
    return i * i
})
fmt.Println(squares.All()) // [1, 4, 9, 16, 25]
```

#### 创建空集合
```go
empty := collections.Empty[string]()
fmt.Println(empty.IsEmpty()) // true
```

---

### 2. 访问元素

#### 获取第一个/最后一个元素
```go
names := collections.Make("Alice", "Bob", "Charlie")

fmt.Println(names.First())  // "Alice"
fmt.Println(names.Last())   // "Charlie"

// 使用默认值
empty := collections.Empty[string]()
fmt.Println(empty.FirstOr("默认值"))  // "默认值"
fmt.Println(empty.LastOr("默认值"))   // "默认值"
```

#### 按索引获取
```go
numbers := collections.New([]int{10, 20, 30, 40, 50})

fmt.Println(numbers.Get(2))       // 30
fmt.Println(numbers.GetOr(10, 0)) // 0 (索引超出范围，返回默认值)
```

#### 条件获取
```go
users := collections.Make("Alice", "Bob", "Charlie")

// 获取第一个以 "B" 开头的名字
name, found := users.FirstWhere(func(s string) bool {
    return len(s) > 0 && s[0] == 'B'
})
if found {
    fmt.Println("Found:", name) // Found: Bob
}
```

---

### 3. 集合信息

```go
numbers := collections.New([]int{1, 2, 3, 4, 5})

fmt.Println(numbers.Count())       // 5
fmt.Println(numbers.IsEmpty())     // false
fmt.Println(numbers.IsNotEmpty())  // true
fmt.Println(numbers.ContainsOneItem()) // false

empty := collections.Empty[int]()
fmt.Println(empty.IsEmpty())       // true
```

---

### 4. 过滤操作

#### Filter - 保留满足条件的元素
```go
numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

// 过滤出偶数
evens := numbers.Filter(func(n int) bool {
    return n%2 == 0
})
fmt.Println(evens.All()) // [2, 4, 6, 8, 10]
```

#### Reject - 排除满足条件的元素
```go
// 排除小于 5 的数字
result := numbers.Reject(func(n int) bool {
    return n < 5
})
fmt.Println(result.All()) // [5, 6, 7, 8, 9, 10]
```

#### Take 和 Skip
```go
numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

fmt.Println(numbers.Take(3).All())  // [1, 2, 3] - 取前 3 个
fmt.Println(numbers.Take(-3).All()) // [8, 9, 10] - 取后 3 个
fmt.Println(numbers.Skip(7).All())  // [8, 9, 10] - 跳过前 7 个
```

#### TakeWhile 和 SkipWhile
```go
numbers := collections.New([]int{1, 2, 3, 4, 5, 4, 3, 2, 1})

// 取元素直到条件不满足
takeWhile := numbers.TakeWhile(func(n int) bool { return n < 4 })
fmt.Println(takeWhile.All()) // [1, 2, 3]

// 跳过元素直到条件不满足
skipWhile := numbers.SkipWhile(func(n int) bool { return n < 4 })
fmt.Println(skipWhile.All()) // [4, 5, 4, 3, 2, 1]
```

#### TakeUntil 和 SkipUntil
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})

// 取元素直到条件满足
takeUntil := numbers.TakeUntil(func(n int) bool { return n > 3 })
fmt.Println(takeUntil.All()) // [1, 2, 3]

// 跳过元素直到条件满足
skipUntil := numbers.SkipUntil(func(n int) bool { return n > 3 })
fmt.Println(skipUntil.All()) // [4, 5]
```

#### Slice - 切片操作
```go
numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

// 从索引 2 开始取 4 个
result := numbers.Slice(2, 4)
fmt.Println(result.All()) // [3, 4, 5, 6]

// 负索引（从末尾开始）
result2 := numbers.Slice(-3)
fmt.Println(result2.All()) // [8, 9, 10]
```

---

### 5. 转换操作

#### Map - 映射转换
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})

// 将每个数字乘以 2
doubled := collections.Map(numbers, func(n int, index int) int {
    return n * 2
})
fmt.Println(doubled.All()) // [2, 4, 6, 8, 10]

// 转换为不同类型
strings := collections.Map(numbers, func(n int, index int) string {
    return fmt.Sprintf("Number: %d", n)
})
fmt.Println(strings.All()) // [Number: 1, Number: 2, ...]
```

#### FlatMap - 映射并展平
```go
numbers := collections.New([]int{1, 2, 3})

// 每个元素生成多个值
expanded := collections.FlatMap(numbers, func(n int, index int) []int {
    return []int{n, n * 10, n * 100}
})
fmt.Println(expanded.All()) // [1, 10, 100, 2, 20, 200, 3, 30, 300]
```

#### Reduce - 归约
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})

// 计算总和
sum := collections.Reduce(numbers, func(acc int, n int, index int) int {
    return acc + n
}, 0)
fmt.Println(sum) // 15

// 计算乘积
product := collections.Reduce(numbers, func(acc int, n int, index int) int {
    return acc * n
}, 1)
fmt.Println(product) // 120
```

#### Pluck - 提取属性
```go
type User struct {
    ID   int
    Name string
    Age  int
}

users := collections.New([]User{
    {ID: 1, Name: "Alice", Age: 30},
    {ID: 2, Name: "Bob", Age: 25},
    {ID: 3, Name: "Charlie", Age: 35},
})

// 提取所有名字
names := collections.Pluck(users, func(u User) string {
    return u.Name
})
fmt.Println(names.All()) // [Alice, Bob, Charlie]

// 提取所有年龄
ages := collections.Pluck(users, func(u User) int {
    return u.Age
})
fmt.Println(ages.All()) // [30, 25, 35]
```

---

### 6. 排序操作

#### Sort - 升序排序
```go
numbers := collections.New([]int{3, 1, 4, 1, 5, 9, 2, 6})

sorted := collections.Sort(numbers)
fmt.Println(sorted.All()) // [1, 1, 2, 3, 4, 5, 6, 9]
```

#### SortDesc - 降序排序
```go
sorted := collections.SortDesc(numbers)
fmt.Println(sorted.All()) // [9, 6, 5, 4, 3, 2, 1, 1]
```

#### SortBy - 按字段排序
```go
type User struct {
    Name string
    Age  int
}

users := collections.New([]User{
    {Name: "Charlie", Age: 35},
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})

// 按年龄排序
byAge := collections.SortBy(users, func(u User) int {
    return u.Age
})
// 输出: [{Bob 25} {Alice 30} {Charlie 35}]

// 按年龄降序排序
byAgeDesc := collections.SortByDesc(users, func(u User) int {
    return u.Age
})
// 输出: [{Charlie 35} {Alice 30} {Bob 25}]
```

#### Reverse - 反转
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})
reversed := numbers.Reverse()
fmt.Println(reversed.All()) // [5, 4, 3, 2, 1]
```

#### Shuffle - 随机打乱
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})
shuffled := numbers.Shuffle()
fmt.Println(shuffled.All()) // 随机顺序
```

---

### 7. 聚合操作

#### 数学计算
```go
numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

fmt.Println(collections.Sum(numbers))    // 55
fmt.Println(collections.Avg(numbers))    // 5.5
fmt.Println(collections.Min(numbers))    // 1
fmt.Println(collections.Max(numbers))    // 10
fmt.Println(collections.Median(numbers)) // 5.5
```

#### Mode - 众数
```go
numbers := collections.New([]int{1, 2, 2, 3, 3, 3, 4})
mode := collections.Mode(numbers)
fmt.Println(mode) // [3] - 出现最多的值
```

#### 按字段聚合
```go
type Product struct {
    Name  string
    Price float64
}

products := collections.New([]Product{
    {Name: "Apple", Price: 1.50},
    {Name: "Banana", Price: 0.75},
    {Name: "Orange", Price: 2.00},
})

// 计算总价
total := collections.SumBy(products, func(p Product) float64 {
    return p.Price
})
fmt.Println(total) // 4.25

// 计算平均价格
avgPrice := collections.AvgBy(products, func(p Product) float64 {
    return p.Price
})
fmt.Println(avgPrice) // 1.4166...

// 找最贵的产品
mostExpensive := collections.MaxBy(products, func(p Product) float64 {
    return p.Price
})
fmt.Println(mostExpensive.Name) // Orange

// 找最便宜的产品
cheapest := collections.MinBy(products, func(p Product) float64 {
    return p.Price
})
fmt.Println(cheapest.Name) // Banana
```

---

### 8. 分组操作

#### GroupBy - 分组
```go
type User struct {
    Name   string
    Age    int
    Gender string
}

users := collections.New([]User{
    {Name: "Alice", Age: 30, Gender: "F"},
    {Name: "Bob", Age: 25, Gender: "M"},
    {Name: "Charlie", Age: 35, Gender: "M"},
    {Name: "Diana", Age: 28, Gender: "F"},
})

// 按性别分组
byGender := collections.GroupBy(users, func(u User) string {
    return u.Gender
})

fmt.Println("Women:", byGender.Get("F").Count()) // 2
fmt.Println("Men:", byGender.Get("M").Count())   // 2

// 按年龄段分组
byAgeGroup := collections.GroupBy(users, func(u User) string {
    if u.Age < 30 {
        return "young"
    }
    return "senior"
})
```

#### KeyBy - 按键索引
```go
// 按 ID 创建查找表
byID := collections.KeyBy(users, func(u User) string {
    return u.Name
})

alice := byID.Get("Alice")
fmt.Printf("Alice is %d years old\n", alice.Age) // Alice is 30 years old
```

#### CountBy - 计数
```go
// 统计每个性别的人数
genderCount := collections.CountBy(users, func(u User) string {
    return u.Gender
})

fmt.Println("Female count:", genderCount.Get("F")) // 2
fmt.Println("Male count:", genderCount.Get("M"))   // 2
```

#### Partition - 分区
```go
numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

// 分成偶数和奇数两组
evens, odds := numbers.Partition(func(n int) bool {
    return n%2 == 0
})

fmt.Println("Evens:", evens.All()) // [2, 4, 6, 8, 10]
fmt.Println("Odds:", odds.All())   // [1, 3, 5, 7, 9]
```

#### ChunkInto - 分块
```go
numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

// 每 3 个一组
chunks := numbers.ChunkInto(3)
for i, chunk := range chunks {
    fmt.Printf("Chunk %d: %v\n", i, chunk.All())
}
// Chunk 0: [1 2 3]
// Chunk 1: [4 5 6]
// Chunk 2: [7 8 9]
// Chunk 3: [10]
```

#### Split - 分成指定数量的组
```go
numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

// 分成 3 组
groups := numbers.Split(3)
for i, group := range groups {
    fmt.Printf("Group %d: %v\n", i, group.All())
}
// Group 0: [1 2 3 4]
// Group 1: [5 6 7]
// Group 2: [8 9 10]
```

---

### 9. 集合操作

#### Diff - 差集
```go
a := collections.New([]int{1, 2, 3, 4, 5})
b := collections.New([]int{3, 4, 5, 6, 7})

// a 中有但 b 中没有的元素
diff := collections.Diff(a, b)
fmt.Println(diff.All()) // [1, 2]
```

#### Intersect - 交集
```go
intersect := collections.Intersect(a, b)
fmt.Println(intersect.All()) // [3, 4, 5]
```

#### Merge - 合并
```go
c := collections.New([]int{8, 9, 10})
merged := a.Merge(b, c)
fmt.Println(merged.All()) // [1, 2, 3, 4, 5, 3, 4, 5, 6, 7, 8, 9, 10]
```

#### Unique - 去重
```go
numbers := collections.New([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})

unique := collections.UniqueComparable(numbers)
fmt.Println(unique.All()) // [1, 2, 3, 4]

// 使用自定义键去重
type User struct {
    ID   int
    Name string
}

users := collections.New([]User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 1, Name: "Alice Duplicate"},
})

uniqueUsers := users.Unique(func(u User) string {
    return fmt.Sprintf("%d", u.ID)
})
fmt.Println(uniqueUsers.Count()) // 2
```

#### Duplicates - 找重复
```go
numbers := collections.New([]int{1, 2, 2, 3, 3, 3, 4})

duplicates := collections.Duplicates(numbers)
fmt.Println(duplicates.All()) // [2, 3]
```

---

### 10. 条件操作

#### Contains - 检查是否存在
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})

hasEven := numbers.Contains(func(n int) bool {
    return n%2 == 0
})
fmt.Println(hasEven) // true
```

#### Every - 检查所有元素是否满足条件
```go
allPositive := numbers.Every(func(n int) bool {
    return n > 0
})
fmt.Println(allPositive) // true
```

#### When 和 Unless - 条件执行
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})
shouldDouble := true

result := numbers.When(shouldDouble, func(c *collections.Collection[int]) *collections.Collection[int] {
    return collections.Map(c, func(n int, _ int) int { return n * 2 })
})
fmt.Println(result.All()) // [2, 4, 6, 8, 10]

// Unless - 条件为 false 时执行
result2 := numbers.Unless(shouldDouble, func(c *collections.Collection[int]) *collections.Collection[int] {
    return collections.Map(c, func(n int, _ int) int { return n * 3 })
})
fmt.Println(result2.All()) // [1, 2, 3, 4, 5] - shouldDouble 为 true，所以不执行
```

---

### 11. 修改操作

#### Push 和 Pop
```go
numbers := collections.New([]int{1, 2, 3})

numbers.Push(4, 5)
fmt.Println(numbers.All()) // [1, 2, 3, 4, 5]

last := numbers.Pop()
fmt.Println(last)          // 5
fmt.Println(numbers.All()) // [1, 2, 3, 4]
```

#### Prepend 和 Shift
```go
numbers := collections.New([]int{2, 3, 4})

numbers.Prepend(0, 1)
fmt.Println(numbers.All()) // [0, 1, 2, 3, 4]

first := numbers.Shift()
fmt.Println(first)         // 0
fmt.Println(numbers.All()) // [1, 2, 3, 4]
```

#### Transform - 原地转换
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})

numbers.Transform(func(n int, index int) int {
    return n * 10
})
fmt.Println(numbers.All()) // [10, 20, 30, 40, 50]
```

#### Put 和 Forget
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})

// 设置指定索引的值
numbers.Put(2, 100)
fmt.Println(numbers.All()) // [1, 2, 100, 4, 5]

// 删除指定索引
numbers.Forget(2)
fmt.Println(numbers.All()) // [1, 2, 4, 5]
```

---

### 12. 迭代操作

#### Each - 遍历
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})

numbers.Each(func(n int, index int) {
    fmt.Printf("Index %d: %d\n", index, n)
})
```

#### Tap - 调试链
```go
result := numbers.
    Filter(func(n int) bool { return n > 2 }).
    Tap(func(c *collections.Collection[int]) {
        fmt.Println("After filter:", c.All())
    }).
    Take(2)
```

---

### 13. MapCollection - 键值对集合

```go
// 创建 MapCollection
data := collections.NewMap(map[string]int{
    "apples":  5,
    "bananas": 3,
    "oranges": 8,
})

// 基本操作
fmt.Println(data.Get("apples"))     // 5
fmt.Println(data.Has("apples"))     // true
fmt.Println(data.Has("grapes"))     // false
fmt.Println(data.Keys().All())      // [apples, bananas, oranges]
fmt.Println(data.Values().All())    // [5, 3, 8]

// 添加和删除
data.Put("grapes", 10)
data.Forget("bananas")

// 过滤
expensive := data.Filter(func(v int, k string) bool {
    return v > 4
})

// Only 和 Except
only := data.Only("apples", "oranges")
except := data.Except("apples")

// 合并
other := collections.NewMap(map[string]int{"mangoes": 6})
merged := data.Merge(other)
```

---

### 14. 其他实用操作

#### Nth - 获取每第 N 个元素
```go
numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

// 每隔 2 个取一个
nth := numbers.Nth(2)
fmt.Println(nth.All()) // [1, 3, 5, 7, 9]

// 从索引 1 开始，每隔 2 个取一个
nth2 := numbers.Nth(2, 1)
fmt.Println(nth2.All()) // [2, 4, 6, 8, 10]
```

#### ForPage - 分页
```go
numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

// 每页 3 个，获取第 2 页
page := numbers.ForPage(2, 3)
fmt.Println(page.All()) // [4, 5, 6]
```

#### Sliding - 滑动窗口
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})

windows := numbers.Sliding(3)
for _, window := range windows {
    fmt.Println(window.All())
}
// [1, 2, 3]
// [2, 3, 4]
// [3, 4, 5]

// 带步长的滑动窗口
windows2 := numbers.Sliding(2, 2)
for _, window := range windows2 {
    fmt.Println(window.All())
}
// [1, 2]
// [3, 4]
```

#### Pad - 填充
```go
numbers := collections.New([]int{1, 2, 3})

// 右填充
padded := numbers.Pad(5, 0)
fmt.Println(padded.All()) // [1, 2, 3, 0, 0]

// 左填充
paddedLeft := numbers.Pad(-5, 0)
fmt.Println(paddedLeft.All()) // [0, 0, 1, 2, 3]
```

#### RandomN - 随机获取多个
```go
numbers := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

random := numbers.RandomN(3)
fmt.Println(random.All()) // 随机 3 个元素
```

#### Clone - 克隆
```go
original := collections.New([]int{1, 2, 3})
cloned := original.Clone()

cloned.Push(4)
fmt.Println(original.All()) // [1, 2, 3] - 原集合不受影响
fmt.Println(cloned.All())   // [1, 2, 3, 4]
```

#### ToJSON - 转为 JSON
```go
numbers := collections.New([]int{1, 2, 3, 4, 5})

jsonStr := numbers.ToJSONString()
fmt.Println(jsonStr) // [1,2,3,4,5]
```

---

### 15. Zip 和 CrossJoin

#### Zip - 合并多个集合
```go
a := collections.New([]int{1, 2, 3})
b := collections.New([]int{4, 5, 6})
c := collections.New([]int{7, 8, 9})

zipped := collections.Zip(a, b, c)
for _, row := range zipped.All() {
    fmt.Println(row)
}
// [1, 4, 7]
// [2, 5, 8]
// [3, 6, 9]
```

#### CrossJoin - 笛卡尔积
```go
colors := collections.New([]string{"red", "blue"})
sizes := collections.New([]string{"S", "M", "L"})

product := collections.CrossJoin(colors, sizes)
for _, combo := range product.All() {
    fmt.Println(combo)
}
// [red, S]
// [red, M]
// [red, L]
// [blue, S]
// [blue, M]
// [blue, L]
```

---

### 16. 组合键值对

#### Combine - 组合
```go
keys := collections.New([]string{"name", "age", "city"})
values := collections.New([]string{"Alice", "30", "Beijing"})

combined := collections.Combine(keys, values)
fmt.Println(combined.Get("name")) // Alice
fmt.Println(combined.Get("age"))  // 30
```

#### MapWithKeys - 映射为键值对
```go
type User struct {
    ID   int
    Name string
}

users := collections.New([]User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
})

result := collections.MapWithKeys(users, func(u User, _ int) (int, string) {
    return u.ID, u.Name
})
fmt.Println(result.Get(1)) // Alice
```

---

### 17. 字符串操作

#### Implode 和 Join
```go
names := collections.New([]string{"Alice", "Bob", "Charlie"})

// 简单连接
result := collections.ImplodeStrings(names, ", ")
fmt.Println(result) // Alice, Bob, Charlie

// 带最后一个分隔符
result2 := collections.JoinStrings(names, ", ", " and ")
fmt.Println(result2) // Alice, Bob and Charlie
```

---

## 🔧 Arr 辅助函数

类似 Laravel 的 Arr 类，提供嵌套数据操作：

```go
data := map[string]any{
    "user": map[string]any{
        "name": "Alice",
        "profile": map[string]any{
            "age":  30,
            "city": "Beijing",
        },
    },
}

// 使用点符号获取嵌套值
name := collections.Arr.Get(data, "user.name")
fmt.Println(name) // Alice

city := collections.Arr.Get(data, "user.profile.city")
fmt.Println(city) // Beijing

// 设置嵌套值
collections.Arr.Set(data, "user.profile.country", "China")

// 检查键是否存在
exists := collections.Arr.Has(data, "user.name")
fmt.Println(exists) // true

// 删除键
collections.Arr.Forget(data, "user.profile.age")
```

---

## 📋 API 速查表

### Collection[T] 方法

| 分类 | 方法 | 描述 |
|------|------|------|
| **创建** | `New(items)` | 从切片创建 |
| | `Make(items...)` | 从可变参数创建 |
| | `Range(from, to)` | 创建数字序列 |
| | `Times(n, fn)` | 重复调用函数创建 |
| | `Empty()` | 创建空集合 |
| **访问** | `All()` | 获取所有元素 |
| | `Count()` | 获取元素数量 |
| | `First()` / `Last()` | 获取首/尾元素 |
| | `Get(index)` | 按索引获取 |
| | `FirstWhere(fn)` | 第一个满足条件的 |
| **过滤** | `Filter(fn)` | 保留满足条件的 |
| | `Reject(fn)` | 排除满足条件的 |
| | `Take(n)` / `Skip(n)` | 取/跳过 n 个 |
| | `TakeWhile(fn)` / `SkipWhile(fn)` | 条件取/跳过 |
| | `Slice(offset, length)` | 切片 |
| **转换** | `Map(c, fn)` | 映射 |
| | `FlatMap(c, fn)` | 映射并展平 |
| | `Reduce(c, fn, init)` | 归约 |
| | `Pluck(c, fn)` | 提取字段 |
| **排序** | `Sort(c)` / `SortDesc(c)` | 排序 |
| | `SortBy(c, fn)` | 按字段排序 |
| | `Reverse()` | 反转 |
| | `Shuffle()` | 随机打乱 |
| **聚合** | `Sum(c)` / `Avg(c)` | 求和/平均 |
| | `Min(c)` / `Max(c)` | 最小/最大 |
| | `Median(c)` / `Mode(c)` | 中位数/众数 |
| **分组** | `GroupBy(c, fn)` | 分组 |
| | `KeyBy(c, fn)` | 按键索引 |
| | `Partition(fn)` | 分区 |
| | `ChunkInto(size)` | 分块 |
| **集合** | `Diff(c, other)` | 差集 |
| | `Intersect(c, other)` | 交集 |
| | `Merge(others...)` | 合并 |
| | `UniqueComparable(c)` | 去重 |
| **修改** | `Push(items...)` | 追加 |
| | `Pop()` | 弹出尾部 |
| | `Prepend(items...)` | 前置 |
| | `Shift()` | 弹出头部 |

---

## 📝 License

MIT License

---

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！
