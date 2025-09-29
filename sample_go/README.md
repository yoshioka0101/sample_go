# Go init 関数 サンプルプロジェクト

このプロジェクトは、Go言語のinit関数の動作を実証するためのサンプルコードです。

## プロジェクト構成

```
sample_go/
├── go.mod              # Go モジュール定義
├── main.go             # メインプログラム（package b をimport）
├── README.md           # このファイル
├── docs.md             # init関数の特性まとめ
├── article.md          # init関数の詳細解説記事
├── a/
│   └── main.go         # 基本的なinit関数の例
├── b/
│   └── b.go            # package c を利用するパッケージ
└── c/
    └── c.go            # 依存される側のパッケージ
```

## 実行方法

### 1. 基本的なinit関数の動作確認
```bash
cd a
go run main.go
```

**期待される出力:**
```
最初に実装されます init
次に
```

### 2. パッケージ間でのinit関数の実行順序確認
```bash
go run main.go
```

**期待される出力:**
```
package c の init 関数が実行されました
package b の init 関数が実行されました
main package の init 関数が実行されました
main 関数が開始されました
package c の関数が呼び出されました
main 関数が終了しました
```

## 学習ポイント

1. **実行順序**: 依存関係に基づいて c → b → main の順で実行される
2. **自動実行**: init関数は明示的な呼び出しなしに自動実行される
3. **main関数より優先**: すべてのinit関数がmain関数より先に実行される

## 詳細な解説

- `docs.md`: init関数の特性を簡潔にまとめたドキュメント
- `article.md`: 実践的な使い方まで含めた詳細解説記事

## 要件

- Go 1.21 以上