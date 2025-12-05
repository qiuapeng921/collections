// Package main demonstrates error handling and safe methods
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
	Active bool
}

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("错误处理和安全方法示例")
	fmt.Println(strings.Repeat("=", 60))

	// 准备测试数据
	users := collections.New([]User{
		{ID: 1, Name: "张三", Age: 25, Active: true},
		{ID: 2, Name: "李四", Age: 30, Active: true},
		{ID: 3, Name: "王五", Age: 35, Active: false},
	})
	empty := collections.Empty[User]()
	numbers := collections.Make(1, 2, 3, 4, 5)

	// ============== 错误类型 ==============
	fmt.Println("\n【1. 错误类型】")

	// ItemNotFoundException
	err1 := &collections.ItemNotFoundException{}
	fmt.Printf("ItemNotFoundException: %v\n", err1.Error())

	err2 := &collections.ItemNotFoundException{Message: "用户未找到"}
	fmt.Printf("自定义消息: %v\n", err2.Error())

	// MultipleItemsFoundException
	err3 := &collections.MultipleItemsFoundException{}
	fmt.Printf("MultipleItemsFoundException: %v\n", err3.Error())

	// InvalidArgumentException
	err4 := &collections.InvalidArgumentException{}
	fmt.Printf("InvalidArgumentException: %v\n", err4.Error())

	// ============== FirstOrFail ==============
	fmt.Println("\n【2. FirstOrFail - 获取第一个或返回错误】")

	// 成功情况
	first, err := users.FirstOrFail()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("第一个用户: %s\n", first.Name)
	}

	// 失败情况
	_, err = empty.FirstOrFail()
	if err != nil {
		fmt.Printf("空集合错误: %v\n", err)
	}

	// ============== LastOrFail ==============
	fmt.Println("\n【3. LastOrFail - 获取最后一个或返回错误】")

	last, err := users.LastOrFail()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("最后一个用户: %s\n", last.Name)
	}

	_, err = empty.LastOrFail()
	if err != nil {
		fmt.Printf("空集合错误: %v\n", err)
	}

	// ============== GetOrFail ==============
	fmt.Println("\n【4. GetOrFail - 按索引获取或返回错误】")

	// 成功情况
	user, err := users.GetOrFail(1)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("索引1的用户: %s\n", user.Name)
	}

	// 失败情况 - 索引超出范围
	_, err = users.GetOrFail(10)
	if err != nil {
		fmt.Printf("索引超出范围: %v\n", err)
	}

	_, err = users.GetOrFail(-1)
	if err != nil {
		fmt.Printf("负索引: %v\n", err)
	}

	// ============== FirstWhereOrFail ==============
	fmt.Println("\n【5. FirstWhereOrFail - 条件查找或返回错误】")

	// 成功情况
	found, err := users.FirstWhereOrFail(func(u User) bool {
		return u.Age > 30
	})
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("年龄>30的第一个用户: %s\n", found.Name)
	}

	// 失败情况 - 没有匹配
	_, err = users.FirstWhereOrFail(func(u User) bool {
		return u.Age > 50
	})
	if err != nil {
		fmt.Printf("没有年龄>50的用户: %v\n", err)
	}

	// ============== RandomOrFail ==============
	fmt.Println("\n【6. RandomOrFail - 随机获取或返回错误】")

	rand, err := numbers.RandomOrFail()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("随机数字: %d\n", rand)
	}

	_, err = collections.Empty[int]().RandomOrFail()
	if err != nil {
		fmt.Printf("空集合随机: %v\n", err)
	}

	// ============== PopOrFail ==============
	fmt.Println("\n【7. PopOrFail - 弹出最后一个或返回错误】")

	nums := collections.Make(1, 2, 3)
	popped, err := nums.PopOrFail()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("弹出: %d, 剩余: %v\n", popped, nums.All())
	}

	emptyNums := collections.Empty[int]()
	_, err = emptyNums.PopOrFail()
	if err != nil {
		fmt.Printf("空集合Pop: %v\n", err)
	}

	// ============== ShiftOrFail ==============
	fmt.Println("\n【8. ShiftOrFail - 弹出第一个或返回错误】")

	nums2 := collections.Make(1, 2, 3)
	shifted, err := nums2.ShiftOrFail()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("Shift: %d, 剩余: %v\n", shifted, nums2.All())
	}

	_, err = collections.Empty[int]().ShiftOrFail()
	if err != nil {
		fmt.Printf("空集合Shift: %v\n", err)
	}

	// ============== Sole ==============
	fmt.Println("\n【9. Sole - 获取唯一元素】")

	// 只有一个元素
	single := collections.Make(42)
	sole, err := single.Sole()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("唯一元素: %d\n", sole)
	}

	// 多个元素
	_, err = numbers.Sole()
	if err != nil {
		fmt.Printf("多个元素: %v\n", err)
	}

	// 空集合
	_, err = collections.Empty[int]().Sole()
	if err != nil {
		fmt.Printf("空集合: %v\n", err)
	}

	// ============== SoleWhere ==============
	fmt.Println("\n【10. SoleWhere - 获取唯一匹配元素】")

	// 唯一匹配
	soleUser, err := users.SoleWhere(func(u User) bool {
		return u.Name == "张三"
	})
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("唯一匹配: %s\n", soleUser.Name)
	}

	// 多个匹配
	_, err = users.SoleWhere(func(u User) bool {
		return u.Active == true
	})
	if err != nil {
		fmt.Printf("多个匹配: %v\n", err)
	}

	// ============== 实际使用场景 ==============
	fmt.Println("\n【11. 实际使用场景】")

	// 场景1：查找用户
	findUserByID := func(id int) (User, error) {
		return users.FirstWhereOrFail(func(u User) bool {
			return u.ID == id
		})
	}

	user, err = findUserByID(2)
	if err != nil {
		fmt.Printf("用户未找到\n")
	} else {
		fmt.Printf("找到用户: %s\n", user.Name)
	}

	// 场景2：确保唯一管理员
	ensureSingleAdmin := func() (User, error) {
		admins := users.Filter(func(u User) bool {
			return u.Name == "张三" // 假设张三是管理员
		})
		return admins.Sole()
	}

	admin, err := ensureSingleAdmin()
	if err != nil {
		fmt.Printf("管理员配置错误: %v\n", err)
	} else {
		fmt.Printf("唯一管理员: %s\n", admin.Name)
	}

	// 场景3：安全的队列操作
	queue := collections.Make("任务1", "任务2", "任务3")
	fmt.Println("处理队列:")
	for {
		task, err := queue.ShiftOrFail()
		if err != nil {
			fmt.Println("队列为空")
			break
		}
		fmt.Printf("  处理: %s\n", task)
	}
}
