# GoExts 通用工具库

[![Go 报告卡](https://goreportcard.com/badge/github.com/goexts/generic)](https://goreportcard.com/report/github.com/goexts/generic)
[![GoDoc](https://godoc.org/github.com/goexts/generic?status.svg)](https://pkg.go.dev/github.com/goexts/generic)
[![MIT 许可证](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub 发布](https://img.shields.io/github/release/goexts/generic.svg)](https://github.com/goexts/generic/releases)
[![Go 版本](https://img.shields.io/github/go-mod/go-version/goexts/generic)](go.mod)
[![GitHub 星标](https://img.shields.io/github/stars/goexts/generic?style=social)](https://github.com/goexts/generic/stargazers)

一个现代化的、健壮的、类型安全的 Go 通用工具库，旨在通过优雅且高性能的 API 解决常见问题。本库致力于成为现代 Go 开发的基础工具包。

> **状态**：**稳定版** - 本项目已可用于生产环境，并遵循语义化版本控制。

## 特性

- 🚀 **类型安全**：基于 Go 1.18+ 泛型实现类型安全
- ⚡ **高性能**：零分配或最小化内存分配优化
- 🧩 **模块化**：可独立使用的独立包
- 🛠 **全面测试**：完整的测试覆盖
- 📚 **完善文档**：完整的 API 文档和示例

## 快速开始

### 安装

```bash
go get github.com/goexts/generic@latest
```

### 基本用法

```go
package main

import (
	"fmt"
	"github.com/goexts/generic/maps"
	"github.com/goexts/generic/slices"
)

func main() {
	// 处理切片
	nums := []int{1, 2, 3, 4, 5}
	doubled := slices.Map(nums, func(x int) int { return x * 2 })
	filtered := slices.Filter(doubled, func(x int) bool { return x > 5 })
	
	fmt.Println("原始数据:", nums)    // [1 2 3 4 5]
	fmt.Println("翻倍后:", doubled)  // [2 4 6 8 10]
	fmt.Println("过滤后:", filtered) // [6 8 10]

	// 处理映射
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := maps.Keys(m)     // [a b c] (顺序不保证)
	values := maps.Values(m) // [1 2 3] (顺序不保证)
	
	fmt.Println("键:", keys)    // [a b c]
	fmt.Println("值:", values)  // [1 2 3]
}
```

## 文档

完整文档请访问：

- [API 文档](https://pkg.go.dev/github.com/goexts/generic)
- [示例](docs/examples/)
- [贡献指南](.github/CONTRIBUTING.md)

## 核心包

本项目提供了一系列独立的通用包：

*   **`cast`**: 提供安全的泛型类型转换函数。
*   **`cmp`**: 为复杂类型提供通用的比较函数。
*   **`cond`**: 条件操作函数（类三元运算符）。
*   **`configure`**: 功能选项模式的强大实现，支持高级编译工作流。
*   **`maps`**: 提供常见的映射操作函数（如 `Keys`、`Values`、`Clone`）。
*   **`must`**: 错误处理包装器（`must.Must`），适用于错误被视为致命情况的场景（如初始化）。
*   **`promise`**: 类似 JavaScript 的 `Promise` 实现，用于管理异步操作。
*   **`ptr`**: 从字面值创建指针的辅助函数（`ptr.To`）。
*   **`res`**: 受 Rust 启发的泛型 `Result[T, E]` 类型，用于表达性强的显式错误处理。
*   **`set`**: 提供集合数据结构的通用实现，包含常用集合操作。
*   **`slices`**: 提供全面的切片操作函数，包含 `bytes` 和 `runes` 子包。
*   **`strings`**: 字符串操作和转换的通用工具。

## 安装

```bash
go get github.com/goexts/generic
```

## 特色包：`configure`

为了展示本库的设计理念，这里简要介绍 `configure` 包。它提供了一套最佳实践工具，用于对象创建，实现了配置对象构建和最终产品编译的清晰分离。

以下示例展示了如何从专用的 `ClientConfig` 对象创建完全配置的 `*http.Client`：

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goexts/generic/configure"
)

// 1. 定义配置对象及其选项
type ClientConfig struct {
	Timeout   time.Duration
	Transport http.RoundTripper
}

type Option = configure.Option[ClientConfig]

func WithTimeout(d time.Duration) Option {
	return func(c *ClientConfig) { c.Timeout = d }
}

func WithTransport(rt http.RoundTripper) Option {
	return func(c *ClientConfig) { c.Transport = rt }
}

// 2. 定义工厂函数（"编译器"）
func NewHttpClient(c *ClientConfig) (*http.Client, error) {
	return &http.Client{
		Timeout:   c.Timeout,
		Transport: c.Transport,
	}, nil
}

func main() {
	// 3. 使用 Builder 收集选项，然后使用 Compile 创建最终产品
	configBuilder := configure.NewBuilder[ClientConfig]().
		Add(WithTimeout(20 * time.Second)).
		Add(WithTransport(http.DefaultTransport))

	httpClient, err := configure.Compile(configBuilder, NewHttpClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("成功创建 http.Client，超时时间: %s\n", httpClient.Timeout)
}
```

## 贡献

我们欢迎所有贡献！您可以通过以下方式提供帮助：

1. 通过[创建 issue](https://github.com/goexts/generic/issues/new/choose) 报告问题
2. 建议新功能或改进
3. 提交拉取请求

请阅读我们的[贡献指南](.github/CONTRIBUTING.md)和[行为准则](.github/CODE_OF_CONDUCT.md)，了解提交拉取请求的流程以及我们如何协作。

### 开发设置

1. 分叉仓库
2. 克隆您的分叉：`git clone https://github.com/your-username/goexts.git`
3. 运行测试：`go test ./...`
4. 进行更改并提交拉取请求

## 社区

- **问题**：[GitHub Issues](https://github.com/goexts/generic/issues)
- **讨论**：[GitHub Discussions](https://github.com/goexts/generic/discussions)
- **聊天**：[Gitter](https://gitter.im/goexts/community) (即将推出)

## 相关项目

- [GoExts](https://github.com/goexts) - Go 扩展和工具集合
- [Go 标准库](https://pkg.go.dev/std) - 本项目扩展的 Go 标准库

## 项目星标历史

[![Stargazers over time](https://starchart.cc/goexts/generic.svg)](https://starchart.cc/goexts/generic)

## 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。
