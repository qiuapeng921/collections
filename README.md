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

    // 链式操作：过滤偶数并取前3个
    result := numbers.
        Filter(func(n int) bool { return n%2 == 0 }).
        Take(3)
    fmt.Println(result.All()) // [2, 4, 6]

    // 聚合计算
    fmt.Println(collections.Sum(numbers))    // 55
    fmt.Println(collections.Avg(numbers))    // 5.5
    fmt.Println(collections.Max(numbers))    // 10
}
```

## 📁 示例

完整使用示例位于 `examples/` 目录：

```bash
go run ./examples/creation/main.go      # 集合创建
go run ./examples/filter/main.go        # 过滤操作
go run ./examples/transform/main.go     # 转换操作
go run ./examples/sorting/main.go       # 排序聚合
go run ./examples/grouping/main.go      # 分组操作
go run ./examples/map_collection/main.go # MapCollection
go run ./examples/real_world/main.go    # 综合示例
```

## 📋 API 速查表

### 创建
| 方法 | 描述 |
|------|------|
| `New(items)` | 从切片创建 |
| `Make(items...)` | 从可变参数创建 |
| `Range(from, to)` | 创建数字序列 |
| `Times(n, fn)` | 重复调用函数创建 |
| `Empty()` | 创建空集合 |

### 访问
| 方法 | 描述 |
|------|------|
| `All()` | 获取所有元素 |
| `Count()` | 获取元素数量 |
| `First()` / `Last()` | 获取首/尾元素 |
| `Get(index)` | 按索引获取 |
| `FirstWhere(fn)` | 条件查找 |
| `Random()` / `RandomN(n)` | 随机获取 |

### 过滤
| 方法 | 描述 |
|------|------|
| `Filter(fn)` | 保留满足条件的 |
| `Reject(fn)` | 排除满足条件的 |
| `Take(n)` / `Skip(n)` | 取/跳过 n 个 |
| `Unique(fn)` | 去重 |
| `Slice(offset, length)` | 切片 |

### 转换
| 方法 | 描述 |
|------|------|
| `Map(c, fn)` | 映射转换 |
| `FlatMap(c, fn)` | 映射并展平 |
| `Reduce(c, fn, init)` | 归约 |
| `Pluck(c, fn)` | 提取字段 |
| `Reverse()` | 反转 |

### 排序
| 方法 | 描述 |
|------|------|
| `Sort(c)` / `SortDesc(c)` | 升序/降序排序 |
| `SortBy(c, fn)` | 按字段排序 |
| `Shuffle()` | 随机打乱 |

### 聚合
| 方法 | 描述 |
|------|------|
| `Sum(c)` / `Avg(c)` | 求和/平均 |
| `Min(c)` / `Max(c)` | 最小/最大 |
| `Median(c)` / `Mode(c)` | 中位数/众数 |

### 分组
| 方法 | 描述 |
|------|------|
| `GroupBy(c, fn)` | 分组 |
| `KeyBy(c, fn)` | 按键索引 |
| `CountBy(c, fn)` | 按条件计数 |
| `Partition(fn)` | 分区 |
| `ChunkInto(size)` | 分块 |

### 集合操作
| 方法 | 描述 |
|------|------|
| `Diff(c, other)` | 差集 |
| `Intersect(c, other)` | 交集 |
| `Merge(others...)` | 合并 |
| `Duplicates(c)` | 获取重复元素 |

### 修改
| 方法 | 描述 |
|------|------|
| `Push(items...)` | 追加到末尾 |
| `Pop()` | 弹出末尾 |
| `Prepend(items...)` | 添加到开头 |
| `Shift()` | 弹出开头 |

### 条件
| 方法 | 描述 |
|------|------|
| `Contains(fn)` | 是否存在 |
| `Every(fn)` | 是否全部满足 |
| `When(cond, fn)` | 条件执行 |

### MapCollection
| 方法 | 描述 |
|------|------|
| `NewMap(map)` | 创建键值对集合 |
| `Get(key)` / `Has(key)` | 获取/检查 |
| `Put(key, value)` | 设置 |
| `Keys()` / `Values()` | 获取键/值 |
| `Filter(fn)` / `Only(keys...)` | 过滤 |
| `Merge(others...)` | 合并 |

### Arr 帮助类
```go
// 点号语法访问嵌套数据
collections.Arr.Get(data, "user.profile.name")
collections.Arr.Set(data, "user.email", "test@example.com")
collections.Arr.Has(data, "user.name")
collections.Arr.Forget(data, "user.temp")
```

## 📝 License

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！
