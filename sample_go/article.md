# Go言語のinit関数を完全解説：実行順序から実践的な使い方まで

## はじめに

Go言語を学び始めると必ず出会う`init`関数。この特殊な関数は、プログラムの実行前に自動的に呼び出される重要な仕組みです。本記事では、実際のコード例を交えながら、init関数の特性や実践的な使い方を詳しく解説します。

## init関数とは

`init`関数は、Go言語において特別な意味を持つ関数です。`main`関数と同様に、プログラマが明示的に呼び出すことはありません。パッケージが初期化される際に自動的に実行されます。

```go
func init() {
    // 初期化処理
}
```

## 基本的な動作を確認してみよう

まず、最もシンプルなinit関数の例を見てみましょう。

### Package A: 基本的なinit関数

```go
package main

import "fmt"

func init() {
    fmt.Println("最初に実装されます init")
}

func main() {
    fmt.Println("次に")
}
```

**実行結果：**
```
最初に実装されます init
次に
```

この例から、init関数がmain関数よりも先に実行されることが確認できます。

## パッケージ間でのinit関数の実行順序

複数のパッケージを使った場合、init関数はどのような順序で実行されるのでしょうか？実際に確認してみましょう。

### Package C: 依存される側のパッケージ

```go
package c

import "fmt"

func init() {
    fmt.Println("package c の init 関数が実行されました")
}

func CallFromC() {
    fmt.Println("package c の関数が呼び出されました")
}
```

### Package B: Package Cを利用するパッケージ

```go
package b

import (
    "fmt"
    "github.com/yoshioka0101/sample_go/c"
)

func init() {
    fmt.Println("package b の init 関数が実行されました")
}

func CallCFunction() {
    c.CallFromC()
}
```

### Main Package: エントリーポイント

```go
package main

import (
    "fmt"
    "github.com/yoshioka0101/sample_go/b"
)

func init() {
    fmt.Println("main package の init 関数が実行されました")
}

func main() {
    fmt.Println("main 関数が開始されました")
    b.CallCFunction()
    fmt.Println("main 関数が終了しました")
}
```

**実行結果：**
```
package c の init 関数が実行されました
package b の init 関数が実行されました
main package の init 関数が実行されました
main 関数が開始されました
package c の関数が呼び出されました
main 関数が終了しました
```

## init関数の実行順序のルール

実行結果から、以下のルールが確認できます：

### 1. 依存関係に基づく実行順序
- **最も深い依存から実行**: package c → package b → main package
- **import文の順序**: パッケージがimportされた順番ではなく、依存の深さによって決まる
- **1回のみ実行**: 同じパッケージのinit関数は1回のみ実行される

### 2. main関数より前に実行
- すべてのinit関数はmain関数の実行前に完了する
- init関数でエラーが発生するとプログラム全体が停止する

## init関数の特徴と制約

### 特徴
- **自動実行**: プログラマが明示的に呼び出す必要がない
- **引数・戻り値なし**: `func init() { }` の形式のみ
- **複数定義可能**: 1つのパッケージに複数のinit関数を定義できる
- **変数初期化後**: パッケージレベルの変数初期化後に実行される

### 制約
- **直接呼び出し不可**: プログラムから直接呼び出すことはできない
- **エラーハンドリング**: panicが発生するとプログラム全体が停止する
- **実行順序の保証**: 同一パッケージ内でも定義順に実行される保証はない（ファイル名順）

## 実践的なinit関数の使い方

### 1. グローバル変数の初期化

```go
package config

import "os"

var (
    DatabaseURL string
    APIKey     string
)

func init() {
    DatabaseURL = os.Getenv("DATABASE_URL")
    if DatabaseURL == "" {
        DatabaseURL = "localhost:5432"
    }
    
    APIKey = os.Getenv("API_KEY")
    if APIKey == "" {
        panic("API_KEY environment variable is required")
    }
}
```

### 2. データベース接続の初期化

```go
package database

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
    var err error
    DB, err = sql.Open("postgres", "postgres://user:pass@localhost/db")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    
    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to ping database:", err)
    }
    
    log.Println("Database connection established")
}
```

### 3. ロガーの設定

```go
package logger

import (
    "log"
    "os"
)

func init() {
    log.SetOutput(os.Stdout)
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("Logger initialized")
}
```

### 4. フラグやバリデーションの設定

```go
package validation

import "regexp"

var (
    EmailRegex *regexp.Regexp
    PhoneRegex *regexp.Regexp
)

func init() {
    EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    PhoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
}
```

## init関数使用時の注意点

### 1. 循環依存の回避
```go
// 避けるべき例
// package a が package b をimport
// package b が package a をimport
// -> 実行時エラーが発生
```

### 2. テスタビリティの考慮
```go
// 悪い例：テストが困難
var client *http.Client

func init() {
    client = &http.Client{
        Timeout: 30 * time.Second,
    }
}

// 良い例：テスト可能な設計
var client *http.Client

func InitClient() {
    client = &http.Client{
        Timeout: 30 * time.Second,
    }
}

func init() {
    InitClient()
}
```

### 3. エラーハンドリングの重要性
```go
func init() {
    file, err := os.Open("config.json")
    if err != nil {
        // panicではなく、適切なデフォルト値の設定を検討
        log.Printf("Config file not found, using defaults: %v", err)
        return
    }
    defer file.Close()
    
    // 設定ファイルの処理
}
```

## パフォーマンスへの影響

init関数は起動時に実行されるため、プログラムの起動時間に直接影響します。

```go
// 避けるべき例：重い処理
func init() {
    // 大量のデータの読み込み
    data := loadLargeDataset() // 起動時間が長くなる
    processData(data)
}

// 改善例：遅延初期化
var data []string
var once sync.Once

func GetData() []string {
    once.Do(func() {
        data = loadLargeDataset()
    })
    return data
}
```

## 実際に動かしてみよう

このプロジェクトに含まれるサンプルコードを実行してみましょう：

### Package A の実行
```bash
cd a
go run main.go
```

### メインプログラムの実行（Package B, C を含む）
```bash
go run main.go
```

## まとめ

Go言語のinit関数は、プログラムの初期化において非常に重要な役割を果たします。主要なポイントを振り返ってみましょう：

### 実行順序の理解
- 依存関係に基づいて実行される（深い依存から順番に）
- main関数よりも先に実行される
- 同一パッケージでは定義順（ファイル名順）に実行される

### 適切な使用場面
- グローバル変数の初期化
- 外部リソースへの接続設定
- ログ設定やバリデーション準備
- 環境変数の読み込み

### 注意すべき点
- 循環依存の回避
- テスタビリティの確保
- 起動時間への配慮
- 適切なエラーハンドリング

init関数を適切に活用することで、Goプログラムの初期化処理を効率的かつ安全に行うことができます。ただし、過度に複雑な処理は避け、可読性とメンテナンス性を重視した設計を心がけましょう。

## 参考資料

- [The Go Programming Language Specification](https://golang.org/ref/spec)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

---

*この記事で使用したサンプルコードは、実際に動作確認済みです。ぜひご自身の環境でも試してみてください。*