// Package main demonstrates Arr helper methods
package main

import (
	"fmt"
	"strings"

	"github.com/qiuapeng921/collections"
)

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Arr 帮助类方法示例")
	fmt.Println(strings.Repeat("=", 60))

	// 准备嵌套数据
	data := map[string]any{
		"user": map[string]any{
			"name":  "张三",
			"email": "zhangsan@example.com",
			"profile": map[string]any{
				"age":    25,
				"gender": "male",
				"address": map[string]any{
					"city":    "北京",
					"country": "中国",
				},
			},
		},
		"settings": map[string]any{
			"theme":    "dark",
			"language": "zh-CN",
		},
	}

	// 1. Get - 点号语法获取嵌套值
	fmt.Println("\n【1. Get - 点号语法获取】")
	fmt.Printf("user.name: %v\n", collections.Arr.Get(data, "user.name"))
	fmt.Printf("user.profile.age: %v\n", collections.Arr.Get(data, "user.profile.age"))
	fmt.Printf("user.profile.address.city: %v\n", collections.Arr.Get(data, "user.profile.address.city"))
	fmt.Printf("settings.theme: %v\n", collections.Arr.Get(data, "settings.theme"))

	// 带默认值
	fmt.Printf("不存在的键(默认值): %v\n", collections.Arr.Get(data, "user.phone", "未设置"))

	// 2. Set - 点号语法设置值
	fmt.Println("\n【2. Set - 点号语法设置】")
	testData := make(map[string]any)
	collections.Arr.Set(testData, "user.name", "李四")
	collections.Arr.Set(testData, "user.email", "lisi@example.com")
	collections.Arr.Set(testData, "app.config.debug", true)
	fmt.Printf("设置后: %v\n", testData)

	// 3. Has - 检查键是否存在
	fmt.Println("\n【3. Has - 检查键存在】")
	fmt.Printf("Has(user.name): %v\n", collections.Arr.Has(data, "user.name"))
	fmt.Printf("Has(user.phone): %v\n", collections.Arr.Has(data, "user.phone"))
	fmt.Printf("Has多个: %v\n", collections.Arr.Has(data, "user.name", "settings.theme"))

	// 4. HasAny - 检查任意键存在
	fmt.Println("\n【4. HasAny - 任意键存在】")
	fmt.Printf("HasAny(user.phone,user.name): %v\n", collections.Arr.HasAny(data, "user.phone", "user.name"))
	fmt.Printf("HasAny(不存在的键): %v\n", collections.Arr.HasAny(data, "foo", "bar"))

	// 5. Forget - 删除键
	fmt.Println("\n【5. Forget - 删除键】")
	forgetData := map[string]any{
		"a": 1, "b": 2, "c": map[string]any{"d": 3, "e": 4},
	}
	fmt.Printf("删除前: %v\n", forgetData)
	collections.Arr.Forget(forgetData, "a", "c.d")
	fmt.Printf("删除后: %v\n", forgetData)

	// 6. Dot - 扁平化为点号形式
	fmt.Println("\n【6. Dot - 扁平化】")
	nested := map[string]any{
		"user": map[string]any{
			"name": "张三",
			"address": map[string]any{
				"city": "北京",
			},
		},
	}
	dotted := collections.Arr.Dot(nested)
	fmt.Println("扁平化后:")
	for k, v := range dotted {
		fmt.Printf("  %s: %v\n", k, v)
	}

	// 7. Undot - 还原嵌套
	fmt.Println("\n【7. Undot - 还原嵌套】")
	flat := map[string]any{
		"user.name":         "李四",
		"user.age":          30,
		"user.address.city": "上海",
	}
	undotted := collections.Arr.Undot(flat)
	fmt.Printf("还原后: %v\n", undotted)

	// 8. Only 和 Except
	fmt.Println("\n【8. Only 和 Except】")
	simpleData := map[string]any{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Printf("Only(a,b): %v\n", collections.Arr.Only(simpleData, "a", "b"))
	fmt.Printf("Except(a,b): %v\n", collections.Arr.Except(simpleData, "a", "b"))

	// 9. Add - 仅在键不存在时添加
	fmt.Println("\n【9. Add - 条件添加】")
	addData := map[string]any{"name": "张三"}
	collections.Arr.Add(addData, "name", "李四") // 不会覆盖
	collections.Arr.Add(addData, "age", 25)    // 会添加
	fmt.Printf("Add后: %v\n", addData)

	// 10. Pull - 获取并删除
	fmt.Println("\n【10. Pull - 获取并删除】")
	pullData := map[string]any{"name": "张三", "age": 25}
	pulled := collections.Arr.Pull(pullData, "name")
	fmt.Printf("Pull的值: %v, 剩余: %v\n", pulled, pullData)

	// 11. Exists - 顶层键存在检查
	fmt.Println("\n【11. Exists - 顶层键存在】")
	fmt.Printf("Exists(age): %v\n", collections.Arr.Exists(pullData, "age"))
	fmt.Printf("Exists(name): %v\n", collections.Arr.Exists(pullData, "name"))

	// 12. Wrap - 包装为切片
	fmt.Println("\n【12. Wrap - 包装为切片】")
	fmt.Printf("Wrap(单值): %v\n", collections.Arr.Wrap("hello"))
	fmt.Printf("Wrap(切片): %v\n", collections.Arr.Wrap([]any{1, 2, 3}))
	fmt.Printf("Wrap(nil): %v\n", collections.Arr.Wrap(nil))

	// 13. First 和 Last
	fmt.Println("\n【13. First 和 Last】")
	items := []any{1, 2, 3, 4, 5}
	fmt.Printf("First: %v\n", collections.Arr.First(items))
	fmt.Printf("Last: %v\n", collections.Arr.Last(items))

	// 带条件
	first := collections.Arr.First(items, func(v any) bool { return v.(int) > 3 })
	fmt.Printf("First(>3): %v\n", first)

	// 14. Where 和 WhereNotNull
	fmt.Println("\n【14. Where 和 WhereNotNull】")
	mixedItems := []any{1, nil, 2, nil, 3, 4, nil, 5}
	notNull := collections.Arr.WhereNotNull(mixedItems)
	fmt.Printf("WhereNotNull: %v\n", notNull)

	filtered := collections.Arr.Where(items, func(v any, i any) bool {
		return v.(int)%2 == 0
	})
	fmt.Printf("Where(偶数): %v\n", filtered)

	// 15. Divide - 分离键值
	fmt.Println("\n【15. Divide - 分离键值】")
	divideData := map[string]any{"a": 1, "b": 2, "c": 3}
	keys, values := collections.Arr.Divide(divideData)
	fmt.Printf("Keys: %v\n", keys)
	fmt.Printf("Values: %v\n", values)

	// 16. Query - 构建查询字符串
	fmt.Println("\n【16. Query - 构建查询字符串】")
	params := map[string]string{"page": "1", "limit": "10", "sort": "name"}
	query := collections.Arr.Query(params)
	fmt.Printf("Query: %s\n", query)

	// 17. Collapse - 展平二维数组
	fmt.Println("\n【17. Collapse - 展平二维数组】")
	nested2d := [][]any{{1, 2}, {3, 4}, {5, 6}}
	collapsed := collections.Arr.Collapse(nested2d)
	fmt.Printf("Collapse: %v\n", collapsed)

	// 18. Prepend - 前置添加
	fmt.Println("\n【18. Prepend - 前置添加】")
	arr := []any{2, 3, 4}
	prepended := collections.Arr.Prepend(arr, 1)
	fmt.Printf("Prepend: %v\n", prepended)

	// 19. CrossJoin - 笛卡尔积
	fmt.Println("\n【19. CrossJoin - 笛卡尔积】")
	a := []any{"a", "b"}
	b := []any{1, 2}
	cross := collections.Arr.CrossJoin(a, b)
	fmt.Println("CrossJoin:")
	for _, row := range cross {
		fmt.Printf("  %v\n", row)
	}

	// 20. IsList 和 Accessible
	fmt.Println("\n【20. IsList 和 Accessible】")
	fmt.Printf("IsList: %v\n", collections.Arr.IsList(items))
	fmt.Printf("Accessible: %v\n", collections.Arr.Accessible(items))
}
