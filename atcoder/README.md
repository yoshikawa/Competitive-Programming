# AtCoder

AtCoderを解くためのtipsなどをここに入れていく

### 開発環境構築

[Tatamo/atcoder-cli](https://github.com/Tatamo/atcoder-cli) と [online-judge-tools/oj](https://github.com/online-judge-tools/oj) のインストールを前提に書いていきます。

```sh
pip3 install online-judge-tools
npm install -g atcoder-cli
```

```sh
# Login AtCoder with acc
$ acc login

# Login AtCoder with oj
$ oj login https://beta.atcoder.jp/
```

```sh
# コンテストの問題のダウンロード
$ acc new (コンテスト名) # Ex) acc new abc001

# テストをするためのコマンド
$ oj t -d tests -c "go run main.go" # ディレクトリ階層とファイル名に注意
```
