// Package main demonstrates MapCollection methods
package main

import (
	"fmt"
	"strings"

	"github.com/qiuapeng921/collections"
)

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("MapCollection 键值对集合示例")
	fmt.Println(strings.Repeat("=", 60))

	// 1. 创建 MapCollection
	fmt.Println("\n【1. 创建 MapCollection】")

	// 从map创建
	data := collections.NewMap(map[string]int{
		"apples":  5,
		"bananas": 3,
		"oranges": 8,
		"grapes":  12,
	})
	fmt.Printf("水果库存: %v\n", data.All())

	// 有序创建
	orderedData := collections.NewMapOrdered(
		map[string]string{"a": "Alice", "b": "Bob", "c": "Charlie"},
		[]string{"a", "b", "c"},
	)
	fmt.Println("有序Map:")
	orderedData.Keys().Each(func(k string, i int) { fmt.Printf("  %s: %s\n", k, orderedData.Get(k)) })

	// CollectMap
	settings := collections.CollectMap(map[string]any{
		"debug":   true,
		"timeout": 30,
		"host":    "localhost",
	})
	fmt.Printf("设置: %v\n", settings.All())

	// 2. 基本访问
	fmt.Println("\n【2. 基本访问】")
	fmt.Printf("Get(apples): %d\n", data.Get("apples"))
	fmt.Printf("GetOr(mangoes, 0): %d\n", data.GetOr("mangoes", 0))
	fmt.Printf("Has(bananas): %v\n", data.Has("bananas"))
	fmt.Printf("Has(mangoes): %v\n", data.Has("mangoes"))
	fmt.Printf("HasAny(mangoes,apples): %v\n", data.HasAny("mangoes", "apples"))

	// 3. Keys 和 Values
	fmt.Println("\n【3. Keys 和 Values】")
	fmt.Printf("Keys: %v\n", data.Keys().All())
	fmt.Printf("Values: %v\n", data.Values().All())
	fmt.Printf("Count: %d\n", data.Count())
	fmt.Printf("IsEmpty: %v\n", data.IsEmpty())

	// 4. 修改操作
	fmt.Println("\n【4. 修改操作】")
	data.Put("mangoes", 10)
	fmt.Printf("Put后: %v\n", data.All())

	pulled := data.Pull("mangoes")
	fmt.Printf("Pull(mangoes): %d, 剩余: %v\n", pulled, data.All())

	data.Forget("grapes")
	fmt.Printf("Forget后: %v\n", data.All())

	// 5. 过滤操作
	fmt.Println("\n【5. 过滤操作】")
	inventory := collections.NewMap(map[string]int{
		"iphone": 50, "macbook": 30, "ipad": 100, "airpods": 200,
	})

	// Filter
	highStock := inventory.Filter(func(v int, k string) bool { return v > 40 })
	fmt.Printf("库存>40: %v\n", highStock.All())

	// Reject
	lowStock := inventory.Reject(func(v int, k string) bool { return v > 50 })
	fmt.Printf("库存<=50: %v\n", lowStock.All())

	// 6. Only 和 Except
	fmt.Println("\n【6. Only 和 Except】")
	only := inventory.Only("iphone", "macbook")
	fmt.Printf("Only(iphone,macbook): %v\n", only.All())

	except := inventory.Except("airpods")
	fmt.Printf("Except(airpods): %v\n", except.All())

	// 7. Merge 和 Union
	fmt.Println("\n【7. Merge 和 Union】")
	m1 := collections.NewMap(map[string]int{"a": 1, "b": 2})
	m2 := collections.NewMap(map[string]int{"b": 20, "c": 3})

	merged := m1.Merge(m2)
	fmt.Printf("Merge(覆盖): %v\n", merged.All())

	union := m1.Union(m2)
	fmt.Printf("Union(保留): %v\n", union.All())

	// 8. DiffKeys 和 IntersectByKeys
	fmt.Println("\n【8. DiffKeys 和 IntersectByKeys】")
	diff := m1.DiffKeys(m2)
	fmt.Printf("DiffKeys: %v\n", diff.All())

	intersect := m1.IntersectByKeys(m2)
	fmt.Printf("IntersectByKeys: %v\n", intersect.All())

	// 9. First/Last
	fmt.Println("\n【9. First 和 Last】")
	prices := collections.NewMapOrdered(
		map[string]float64{"a": 10.5, "b": 20.5, "c": 30.5},
		[]string{"a", "b", "c"},
	)
	fmt.Printf("First: %.1f, Last: %.1f\n", prices.First(), prices.Last())
	fmt.Printf("FirstKey: %s, LastKey: %s\n", prices.FirstKey(), prices.LastKey())

	// 10. Each 遍历
	fmt.Println("\n【10. Each 遍历】")
	inventory.Each(func(k string, v int) {
		fmt.Printf("  %s: %d\n", k, v)
	})

	// 11. MapValues
	fmt.Println("\n【11. MapValues - 值转换】")
	doubled := collections.MapValues(inventory, func(v int, k string) int { return v * 2 })
	fmt.Printf("库存翻倍: %v\n", doubled.All())

	// 12. ToJSON
	fmt.Println("\n【12. ToJSON - 序列化】")
	jsonStr := inventory.ToJSONString()
	fmt.Printf("JSON: %s\n", jsonStr)

	// 13. Flip (使用 Collection[string])
	fmt.Println("\n【13. Flip - 键值翻转】")
	strCol := collections.Make("apple", "banana", "cherry")
	flipped := collections.Flip(strCol)
	fmt.Println("字符串到索引映射:")
	flipped.Each(func(k string, v int) {
		fmt.Printf("  %s -> %d\n", k, v)
	})

	// 14. 实际应用示例
	fmt.Println("\n【14. 实际应用示例】")

	// 用户配置
	config := collections.NewMap(map[string]any{
		"database.host": "localhost",
		"database.port": 3306,
		"database.name": "mydb",
		"cache.enabled": true,
		"cache.ttl":     3600,
		"app.debug":     false,
		"app.log_level": "info",
	})

	// 只获取数据库配置
	dbConfig := config.Filter(func(v any, k string) bool {
		return strings.HasPrefix(k, "database.")
	})
	fmt.Println("数据库配置:")
	dbConfig.Each(func(k string, v any) { fmt.Printf("  %s: %v\n", k, v) })

	// 获取缓存配置
	cacheConfig := config.Filter(func(v any, k string) bool {
		return strings.HasPrefix(k, "cache.")
	})
	fmt.Println("缓存配置:")
	cacheConfig.Each(func(k string, v any) { fmt.Printf("  %s: %v\n", k, v) })
}
