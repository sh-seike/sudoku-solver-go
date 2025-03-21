# sudoku-solver-go

これはGo言語で書かれた数独を解くためのプログラムです。

## 使い方

### ビルド

以下のコマンドを実行してプロジェクトをビルドします。

```sh
go build -o sudoku-solver
```

### 実行

ビルドが完了したら、以下のコマンドでプログラムを実行できます。

```sh
./sudoku-solver <input-file>
```

`<input-file>`には、解きたい数独の問題が含まれたファイルを指定してください。  
ファイルを指定しなかった場合は`question1.csv`を使用します。

### オプション

プログラム実行時に使用できるオプションは以下の通りです。

- `-d` : デバッグモードを有効にします。
- `-l <level>` : 難易度レベルを指定します。有効な値は1と2で2以上を指定すると仮定を繰り返しながら解くようになります。

例:

```sh
./sudoku-solver -d -l 2 question2.csv
```

## 問題について

問題は1マスを1つの数字で表したCSVファイルで表現されることを想定しています。  
答えが決まっていないマスは0としてください。
