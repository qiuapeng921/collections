// Package main demonstrates comprehensive real-world examples
package main

import (
	"fmt"
	"strings"

	"github.com/qiuapeng921/collections"
)

// ============== 数据模型 ==============

// User 用户
type User struct {
	ID       int
	Name     string
	Email    string
	Age      int
	Role     string
	Active   bool
	Balance  float64
	JoinYear int
}

// Product 产品
type Product struct {
	ID       int
	Name     string
	Price    float64
	Category string
	Stock    int
	Tags     []string
}

// Order 订单
type Order struct {
	ID        int
	UserID    int
	ProductID int
	Quantity  int
	Amount    float64
	Status    string
}

// OrderDetail 订单详情
type OrderDetail struct {
	OrderID     int
	UserName    string
	ProductName string
	Quantity    int
	Amount      float64
	Status      string
}

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("实际应用综合示例")
	fmt.Println(strings.Repeat("=", 60))

	// 准备测试数据
	users := collections.New([]User{
		{ID: 1, Name: "张三", Email: "zhang@test.com", Age: 25, Role: "admin", Active: true, Balance: 1000.0, JoinYear: 2020},
		{ID: 2, Name: "李四", Email: "li@test.com", Age: 30, Role: "user", Active: true, Balance: 500.0, JoinYear: 2019},
		{ID: 3, Name: "王五", Email: "wang@test.com", Age: 35, Role: "user", Active: false, Balance: 2000.0, JoinYear: 2021},
		{ID: 4, Name: "赵六", Email: "zhao@test.com", Age: 28, Role: "moderator", Active: true, Balance: 800.0, JoinYear: 2020},
		{ID: 5, Name: "钱七", Email: "qian@test.com", Age: 22, Role: "user", Active: true, Balance: 300.0, JoinYear: 2022},
	})

	products := collections.New([]Product{
		{ID: 1, Name: "iPhone 15", Price: 999.0, Category: "手机", Stock: 50, Tags: []string{"热卖", "新品"}},
		{ID: 2, Name: "MacBook Pro", Price: 1999.0, Category: "电脑", Stock: 30, Tags: []string{"新品"}},
		{ID: 3, Name: "iPad Air", Price: 599.0, Category: "平板", Stock: 100, Tags: []string{"热卖"}},
		{ID: 4, Name: "AirPods Pro", Price: 249.0, Category: "配件", Stock: 200, Tags: []string{"热卖", "促销"}},
		{ID: 5, Name: "Apple Watch", Price: 399.0, Category: "穿戴", Stock: 80, Tags: []string{"新品"}},
	})

	orders := collections.New([]Order{
		{ID: 1001, UserID: 1, ProductID: 1, Quantity: 1, Amount: 999.0, Status: "completed"},
		{ID: 1002, UserID: 2, ProductID: 2, Quantity: 1, Amount: 1999.0, Status: "pending"},
		{ID: 1003, UserID: 1, ProductID: 4, Quantity: 2, Amount: 498.0, Status: "completed"},
		{ID: 1004, UserID: 3, ProductID: 3, Quantity: 1, Amount: 599.0, Status: "cancelled"},
		{ID: 1005, UserID: 4, ProductID: 5, Quantity: 1, Amount: 399.0, Status: "completed"},
		{ID: 1006, UserID: 5, ProductID: 1, Quantity: 1, Amount: 999.0, Status: "pending"},
	})

	// ============== 场景1：用户分析 ==============
	fmt.Println("\n【场景1：用户分析】")

	// 活跃用户统计
	activeUsers := users.Filter(func(u User) bool { return u.Active })
	fmt.Printf("活跃用户数: %d/%d\n", activeUsers.Count(), users.Count())

	// 按角色分组统计
	byRole := collections.GroupBy(users, func(u User) string { return u.Role })
	fmt.Println("角色分布:")
	byRole.Each(func(role string, us *collections.Collection[User]) {
		fmt.Printf("  %s: %d人\n", role, us.Count())
	})

	// 用户余额统计
	totalBalance := collections.SumBy(users, func(u User) float64 { return u.Balance })
	avgBalance := collections.AvgBy(users, func(u User) float64 { return u.Balance })
	richest := collections.MaxBy(users, func(u User) float64 { return u.Balance })
	fmt.Printf("余额统计: 总计=%.2f, 平均=%.2f, 最高=%s(%.2f)\n",
		totalBalance, avgBalance, richest.Name, richest.Balance)

	// 年龄分析
	youngest := collections.MinBy(users, func(u User) int { return u.Age })
	oldest := collections.MaxBy(users, func(u User) int { return u.Age })
	fmt.Printf("年龄范围: %d~%d岁\n", youngest.Age, oldest.Age)

	// ============== 场景2：产品查询 ==============
	fmt.Println("\n【场景2：产品查询】")

	// 按类别分组
	byCategory := collections.GroupBy(products, func(p Product) string { return p.Category })
	fmt.Println("类别分组:")
	byCategory.Each(func(cat string, ps *collections.Collection[Product]) {
		total := collections.SumBy(ps, func(p Product) float64 { return p.Price * float64(p.Stock) })
		fmt.Printf("  %s: %d件产品, 总价值=%.2f\n", cat, ps.Count(), total)
	})

	// 低库存产品预警
	lowStock := products.Filter(func(p Product) bool { return p.Stock < 50 })
	fmt.Println("低库存预警(<50):")
	lowStock.Each(func(p Product, i int) {
		fmt.Printf("  - %s: 仅剩%d件\n", p.Name, p.Stock)
	})

	// 热卖产品
	hotProducts := products.Filter(func(p Product) bool {
		tags := collections.New(p.Tags)
		return tags.Contains(func(t string) bool { return t == "热卖" })
	})
	fmt.Println("热卖产品:")
	hotProducts.Each(func(p Product, i int) { fmt.Printf("  - %s\n", p.Name) })

	// ============== 场景3：订单分析 ==============
	fmt.Println("\n【场景3：订单分析】")

	// 订单状态统计
	statusCount := collections.CountBy(orders, func(o Order) string { return o.Status })
	fmt.Println("订单状态:")
	statusCount.Each(func(s string, c int) { fmt.Printf("  %s: %d单\n", s, c) })

	// 已完成订单总额
	completed := orders.Filter(func(o Order) bool { return o.Status == "completed" })
	totalRevenue := collections.SumBy(completed, func(o Order) float64 { return o.Amount })
	fmt.Printf("已完成订单总额: %.2f\n", totalRevenue)

	// 用户消费排行
	userOrders := collections.GroupBy(orders.Filter(func(o Order) bool {
		return o.Status == "completed"
	}), func(o Order) int { return o.UserID })

	type UserSpend struct {
		UserID int
		Name   string
		Total  float64
	}
	spends := collections.Map(userOrders.Keys(), func(uid int, i int) UserSpend {
		total := collections.SumBy(userOrders.Get(uid), func(o Order) float64 { return o.Amount })
		user, _ := users.FirstWhere(func(u User) bool { return u.ID == uid })
		return UserSpend{UserID: uid, Name: user.Name, Total: total}
	})
	sorted := collections.SortByDesc(spends, func(s UserSpend) float64 { return s.Total })
	fmt.Println("消费排行:")
	sorted.Each(func(s UserSpend, i int) {
		fmt.Printf("  %d. %s: %.2f\n", i+1, s.Name, s.Total)
	})

	// ============== 场景4：订单详情关联查询 ==============
	fmt.Println("\n【场景4：订单详情关联查询】")

	// 创建用户和产品索引
	userIndex := collections.KeyBy(users, func(u User) int { return u.ID })
	productIndex := collections.KeyBy(products, func(p Product) int { return p.ID })

	// 关联生成订单详情
	details := collections.Map(orders, func(o Order, i int) OrderDetail {
		user := userIndex.Get(o.UserID)
		product := productIndex.Get(o.ProductID)
		return OrderDetail{
			OrderID:     o.ID,
			UserName:    user.Name,
			ProductName: product.Name,
			Quantity:    o.Quantity,
			Amount:      o.Amount,
			Status:      o.Status,
		}
	})

	fmt.Println("订单详情:")
	details.Each(func(d OrderDetail, i int) {
		fmt.Printf("  #%d: %s购买%s x%d = %.2f (%s)\n",
			d.OrderID, d.UserName, d.ProductName, d.Quantity, d.Amount, d.Status)
	})

	// ============== 场景5：报表生成 ==============
	fmt.Println("\n【场景5：报表生成】")

	// 按产品类别的销售统计
	fmt.Println("按类别销售统计:")
	ordersByProduct := collections.GroupBy(completed, func(o Order) int { return o.ProductID })
	ordersByProduct.Each(func(pid int, os *collections.Collection[Order]) {
		product := productIndex.Get(pid)
		qty := collections.SumBy(os, func(o Order) int { return o.Quantity })
		amt := collections.SumBy(os, func(o Order) float64 { return o.Amount })
		fmt.Printf("  %s: %d件, %.2f元\n", product.Name, qty, amt)
	})

	// ============== 场景6：数据分页 ==============
	fmt.Println("\n【场景6：数据分页】")
	pageSize := 2
	totalPages := (products.Count() + pageSize - 1) / pageSize

	for page := 1; page <= totalPages; page++ {
		pageData := products.ForPage(page, pageSize)
		fmt.Printf("第%d页(共%d页):\n", page, totalPages)
		pageData.Each(func(p Product, i int) {
			fmt.Printf("  - %s: %.2f\n", p.Name, p.Price)
		})
	}

	// ============== 场景7：配置管理 ==============
	fmt.Println("\n【场景7：配置管理】")

	config := map[string]any{
		"app": map[string]any{
			"name":  "MyApp",
			"debug": true,
		},
		"database": map[string]any{
			"host": "localhost",
			"port": 3306,
			"name": "mydb",
		},
	}

	fmt.Printf("app.name: %v\n", collections.Arr.Get(config, "app.name"))
	fmt.Printf("database.host: %v\n", collections.Arr.Get(config, "database.host"))
	fmt.Printf("不存在的配置: %v\n", collections.Arr.Get(config, "cache.ttl", 3600))
}
