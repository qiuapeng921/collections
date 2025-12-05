# Collections 使用示例

本目录包含 Collections 库的完整使用示例，涵盖所有主要功能和方法。

## 📂 目录结构

```
examples/
├── creation/       # 集合创建方法
├── access/         # 元素访问方法
├── filter/         # 过滤操作方法
├── transform/      # 转换操作方法
├── sorting/        # 排序和聚合方法
├── grouping/       # 分组操作方法
├── map_collection/ # MapCollection 键值对集合
├── arr_helpers/    # Arr 帮助类方法
├── helpers/        # Helper 帮助函数
├── errors/         # 错误处理和安全方法
├── conditions/     # 条件操作和迭代方法
└── real_world/     # 实际应用综合示例
```

## 🚀 运行示例

在任意示例目录中运行：

```bash
go run main.go
```

或从根目录运行指定示例：

```bash
go run ./examples/creation/main.go
go run ./examples/access/main.go
go run ./examples/filter/main.go
# ... 以此类推
```

## 📖 示例内容

### creation - 集合创建
- `New()` - 从切片创建
- `Make()` - 从可变参数创建
- `Range()` - 创建数字序列
- `Times()` - 重复生成
- `Empty()` - 创建空集合
- `Collect()` / `CollectSlice()` - 快捷创建

### access - 元素访问
- `All()` / `Count()` - 获取全部/数量
- `First()` / `Last()` - 获取首尾元素
- `Get()` / `GetOr()` - 按索引获取
- `FirstWhere()` / `LastWhere()` - 条件查找
- `Search()` - 搜索元素
- `Random()` / `RandomN()` - 随机获取

### filter - 过滤操作
- `Filter()` / `Reject()` - 过滤/排除
- `Take()` / `Skip()` - 获取/跳过
- `TakeWhile()` / `SkipWhile()` - 条件获取/跳过
- `Slice()` - 切片
- `Partition()` - 分区
- `ChunkInto()` / `Split()` - 分块/分组
- `Nth()` / `ForPage()` - 间隔取/分页
- `Unique()` - 去重

### transform - 转换操作
- `Map()` - 映射转换
- `FlatMap()` - 映射并展平
- `Reduce()` - 归约
- `Pluck()` - 提取属性
- `Transform()` - 原地转换
- `Reverse()` / `Shuffle()` - 反转/打乱
- `Merge()` / `Concat()` / `Clone()` - 合并/连接/克隆
- `Pad()` / `Sliding()` - 填充/滑动窗口

### sorting - 排序和聚合
- `Sort()` / `SortDesc()` - 升序/降序排序
- `SortBy()` / `SortByDesc()` - 按字段排序
- `Sum()` / `Avg()` / `Min()` / `Max()` - 数学运算
- `Median()` / `Mode()` - 中位数/众数
- `Diff()` / `Intersect()` - 差集/交集
- `Duplicates()` / `UniqueComparable()` - 重复/去重
- `IndexOf()` / `LastIndexOf()` - 查找索引

### grouping - 分组操作
- `GroupBy()` - 按条件分组
- `KeyBy()` - 按键索引
- `CountBy()` - 按条件计数
- `Partition()` - 二分区
- `ChunkInto()` / `Split()` - 分块
- `MapWithKeys()` - 创建键值对映射
- `MapToDictionary()` - 映射到字典

### map_collection - MapCollection
- `NewMap()` / `NewMapOrdered()` - 创建
- `Get()` / `GetOr()` / `Has()` - 访问
- `Keys()` / `Values()` - 获取键值
- `Put()` / `Pull()` / `Forget()` - 修改
- `Filter()` / `Reject()` - 过滤
- `Only()` / `Except()` - 选择/排除
- `Merge()` / `Union()` - 合并
- `DiffKeys()` / `IntersectByKeys()` - 集合操作

### arr_helpers - Arr 帮助类
- `Get()` / `Set()` - 点号语法访问
- `Has()` / `HasAny()` - 键存在检查
- `Forget()` - 删除键
- `Dot()` / `Undot()` - 扁平化/还原
- `Only()` / `Except()` - 选择/排除
- `Wrap()` / `First()` / `Last()` - 工具方法
- `CrossJoin()` - 笛卡尔积

### helpers - Helper 函数
- `Head()` / `Tail()` / `Init()` / `LastItem()` - 切片操作
- `Blank()` / `Filled()` - 值检查
- `Optional` - 安全的可空值处理
- `Once()` / `Retry()` / `Rescue()` - 执行控制
- `Tap()` / `Transform()` - 值处理
- `DataGet()` / `DataSet()` - 嵌套数据访问

### errors - 错误处理
- `FirstOrFail()` / `LastOrFail()` - 安全获取首尾
- `GetOrFail()` - 安全按索引获取
- `FirstWhereOrFail()` - 安全条件查找
- `Sole()` / `SoleWhere()` - 获取唯一元素
- `RandomOrFail()` / `PopOrFail()` / `ShiftOrFail()` - 安全操作

### conditions - 条件和迭代
- `Contains()` / `Some()` / `Every()` - 条件检查
- `When()` / `Unless()` - 条件执行
- `WhenEmpty()` / `WhenNotEmpty()` - 空值条件
- `Each()` / `EachSpread()` - 遍历
- `Tap()` / `Pipe()` - 链式调试
- `Push()` / `Pop()` / `Prepend()` / `Shift()` - 栈操作
- `ToJSON()` / `String()` / `Dump()` - 序列化

### real_world - 实际应用
- 用户分析场景
- 产品查询场景
- 订单分析场景
- 关联查询场景
- 报表生成场景
- 数据分页场景
- 配置管理场景

## 📝 数据类型示例

每个示例都包含以下数据类型的使用：
- **基本类型**: `int`, `float64`, `string`, `bool`
- **切片**: `[]int`, `[]string`, `[]User`
- **结构体**: `User`, `Product`, `Order`
- **Map**: `map[string]int`, `map[string]any`

## 💡 最佳实践

1. **链式调用** - 利用方法链简化代码
2. **类型安全** - 充分利用泛型特性
3. **不可变操作** - 大多数方法返回新集合
4. **错误处理** - 使用 `OrFail` 方法处理边界情况
