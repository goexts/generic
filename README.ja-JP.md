# GoExts ジェネリックユーティリティ

<div align="right">
  <a href="./README.md">English</a> | 
  <a href="./README.zh-CN.md">中文</a> | 
  <a href="./README.ja-JP.md">日本語</a>
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/goexts/generic)](https://goreportcard.com/report/github.com/goexts/generic)
[![GoDoc](https://godoc.org/github.com/goexts/generic?status.svg)](https://pkg.go.dev/github.com/goexts/generic)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub release](https://img.shields.io/badge/release-MIT-blue.svg)](https://github.com/goexts/generic/releases)
[![Go version](https://img.shields.io/github/go-mod/go-version/goexts/generic)](go.mod)
[![GitHub stars](https://img.shields.io/github/stars/goexts/generic?style=social)](https://github.com/goexts/generic/stargazers)

Goのためのモダンで堅牢、かつ型安全なジェネリックユーティリティのコレクション。エレガントで高性能なAPIで一般的な問題を解決するように設計されています。

> **ステータス**: **安定版** - このプロジェクトは本番環境に対応しており、セマンティックバージョニングに従っています。

## インストール

```bash
go get github.com/goexts/generic@latest
```

## ドキュメント

すべてのパッケージの完全なガイド、APIリファレンス、および使用例については、公式のGoドキュメントをご覧ください。

**[pkg.go.dev/github.com/goexts/generic](https://pkg.go.dev/github.com/goexts/generic)**

## パッケージ

このライブラリは、独立したジェネリックパッケージの豊富なセットを提供します。

*   **`cast`**: 安全でジェネリックな型キャスト関数。
*   **`cmp`**: ソートと順序付けのためのジェネリックな比較関数。
*   **`cond`**: 三項演算子のような条件関数。
*   **`configure`**: 関数型オプションパターンの強力な実装。
*   **`maps`**: 一般的なマップ操作のためのジェネリック関数のスイート（`x/exp/maps`のアダプター）。
*   **`must`**: よりクリーンな初期化コードのためのパニックオンエラーラッパー。
*   **`promise`**: 非同期操作を管理するためのジェネリックなJavaScriptライクな`Promise`実装。
*   **`ptr`**: リテラル値からポインタを作成するためのヘルパー関数。
*   **`res`**: 表現力豊かなエラーハンドリングのためのRustにインスパイアされたジェネリックな`Result[T]`型。
*   **`set`**: ステートレスなスライスベースのセット操作。
*   **`slices`**: スライス操作のための包括的な関数のスイート。
*   **`strings`**: 文字列操作のための関数のコレクション（標準の`strings`パッケージのアダプター）。

## 注目の例: `configure`

`configure`パッケージは、堅牢で型安全な**関数型オプションパターン**の実装を提供します。このパターンは、複数のオプションパラメータを持つ複雑なオブジェクトをクリーンで読みやすい方法で作成するのに理想的です。

### 1. 基本的な関数型オプション (`configure.Apply`)

これはコアパターンを示しています。オプションを定義し、それらをデフォルト設定に適用します。`configure.Apply`はオプションのスライスを期待することに注意してください。

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goexts/generic/configure"
)

// ClientConfig holds the configuration for our http.Client.
type ClientConfig struct {
	Timeout   time.Duration
	Transport http.RoundTripper
}

// Option defines the functional option type for our configuration.
type Option = configure.Option[ClientConfig]

// WithTimeout returns an option to set the client timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *ClientConfig) {
		c.Timeout = d
	}
}

// WithTransport returns an option to set the client's transport.
func WithTransport(rt http.RoundTripper) Option {
	return func(c *ClientConfig) {
		c.Transport = rt
	}
}

// NewClient creates a new http.Client with default settings,
// then applies the provided options.
func NewClient(opts ...Option) *http.Client {
	// 1. Start with a default configuration.
	config := &ClientConfig{
		Timeout:   10 * time.Second,
		Transport: http.DefaultTransport,
	}

	// 2. Apply any user-provided options over the defaults.
	// configure.Apply modifies the config in-place and returns it.
	configure.Apply(config, opts)

	// 3. The final, configured object is now ready to be used.
	return &http.Client{
		Timeout:   config.Timeout,
		Transport: config.Transport,
	}
}

func main() {
	// Create a client with default settings (no options).
	defaultClient := NewClient()
	fmt.Printf("Default client timeout: %s\n", defaultClient.Timeout)

	// Create a client with a custom timeout, overriding the default.
	// Pass options as variadic arguments.
	customClient := NewClient(WithTimeout(30 * time.Second))
	fmt.Printf("Custom client timeout: %s\n", customClient.Timeout)

	// If you have individual options, you can pass them directly:
	// anotherClient := NewClient(WithTimeout(20 * time.Second), WithTransport(&http.Transport{}))
}
```

### 2. `Builder` を使用した高度な設定

`Builder`は、オプションを段階的に収集するための流れるようなインターフェースを提供します。これは、オプションがさまざまなソースから、または複数の段階で収集される場合に役立ちます。また、ベース設定の設定もサポートしています。

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goexts/generic/configure"
)

// ClientConfig holds the configuration for our http.Client.
type ClientConfig struct {
	Timeout       time.Duration
	Transport     http.RoundTripper
	EnableTracing bool // Added to demonstrate AddWhen's optIfFalse
}

// Option defines the functional option type for our configuration.
type Option = configure.Option[ClientConfig]

// WithTimeout returns an option to set the client timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *ClientConfig) {
		c.Timeout = d
	}
}

// WithTransport returns an option to set the client's transport.
func WithTransport(rt http.RoundTripper) Option {
	return func(c *ClientConfig) {
		c.Transport = rt
	}
}

// WithTracing enables or disables tracing.
func WithTracing(enable bool) Option {
	return func(c *ClientConfig) {
		c.EnableTracing = enable
	}
}

// NewClient is the factory function for Compile.
// It takes the final, fully-built configuration and creates the product (*http.Client).
func NewClient(config *ClientConfig) (*http.Client, error) {
	// In a real scenario, you might use config.EnableTracing here
	// to configure the http.Client or a wrapper around it.
	return &http.Client{
		Timeout:   config.Timeout,
		Transport: config.Transport,
	}, nil
}

func main() {
	// Define a base configuration that can be reused.
	baseConfig := &ClientConfig{
		Timeout:       5 * time.Second,
		Transport:     http.DefaultTransport,
		EnableTracing: false,
	}

	// Create a builder, passing the base configuration directly to NewBuilder.
	builder := configure.NewBuilder[ClientConfig](baseConfig).
		Add(WithTimeout(15 * time.Second)). // Overrides baseConfig.Timeout
		// Use AddWhen with optIfTrue and optIfFalse
		AddWhen(true, WithTracing(true), WithTracing(false))

	// Compile the final product using the builder and a factory function.
	// Note: configure.Compile now expects factory first, then builder.
	client1, err := configure.Compile(NewClient, builder)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client 1 (Builder) timeout: %s\n", client1.Timeout)

	// Demonstrate using Chain to group options before adding to builder
	commonOptions := configure.Chain(
		WithTransport(&http.Transport{}), // Custom transport
		WithTimeout(20 * time.Second),
	)

	client2, err := configure.Compile(
		NewClient,
		configure.NewBuilder[ClientConfig](baseConfig).
			Add(commonOptions). // Add chained options
			// AddWhen with false condition, so optIfFalse (WithTracing(false)) will be applied
			AddWhen(false, WithTracing(true), WithTracing(false)),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client 2 (Builder) transport type: %T, timeout: %s\n", client2.Transport, client2.Timeout)

	// Example of using Builder directly as an option (implements ApplierE)
	// This is useful if you want to apply a set of options defined by a builder
	// to an existing config object or within another ApplyAny call.
	existingConfig := &ClientConfig{Timeout: 1 * time.Second}
	err = configure.NewBuilder[ClientConfig]().
		Add(WithTimeout(30 * time.Second)).
		Apply(existingConfig) // Apply builder's options to an existing config
	if err != nil {
		panic(err)
	}
	fmt.Printf("Existing config after builder.Apply: %s\n", existingConfig.Timeout)
}
```

## 貢献

貢献を歓迎します！詳細については、[貢献ガイド](.github/CONTRIBUTING.md)をご覧ください。

## ライセンス

このプロジェクトはMITライセンスの下でライセンスされています。詳細については、[LICENSE](LICENSE)ファイルをご覧ください。
