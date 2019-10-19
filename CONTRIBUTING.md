# CONTRIBUTING.md

## 貢献する手順

1. forkする
2. branchを切る
3. commitを縮める
4. Pull Requestを作成
5. コードレビュー
6. Merge

## Indent

`Python`なら、タブスペースの幅は`4` (スペースで統一)

`go`でも同様

## commit

commitをきれいに保つために、Pull Requestを作成する前にブランチを押しつぶしてください。

### squashのやり方

1. git log (commit logの確認)
2. git rebase -i "消したいcommit logの一つ先を見る"
3. テキストエディタで、squashしたい場所に`s`、残したい場所は`pick`
4. git push -f bura bura

## fork先からfork元のリポジトリを追従する方法

1. `git remote add upstream`
2. `git fetch upstream`
3. `git merge upstream branch_name`

**CAUTION**

`git pull` を使わないこと

## Reference

> https://qiita.com/xtetsuji/items/555a1ef19ed21ee42873
