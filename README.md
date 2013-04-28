git commit で自動的に Pivotal Tracker にノートを追加
==============================================

大まかな流れ
----------

1. git に設定を追加
2. git2pivotal をパスが通った場所に置く
3. post-commit フックに登録する

git に必要な設定を追加
------------------

1. Git のリポジトリ上で以下の設定を追加する
```sh
# PivotalTracker の API トークン
# https://www.pivotaltracker.com/profile の一番下で作成する
$ git config git2pivotal.token "0f1e2d3c4b5a69788796a5b4c3d2e1f0"
```
```sh
# プロジェクトの ID
$ git config git2pivotal.project "123456"
```
```sh
# git2pivotal で投稿されるときに文頭に追記されるフレーズ
# git log --pretty=format で使用可能なものが使える
# この例ではコミットタイトルの行の後に改行を設定している
$ git config git2pivotal.prescript "%s%n"
```
```sh
# git2pivotal で投稿されるときに文末に追記されるフレーズ
# git log --pretty=format で使用可能なものが使える
# この例では空行に続いて github 上で参照可能になる予定の URL　の加えている
$ git config git2pivotal.postscript "%nhttps://github.com/MyAccount/myrepos/commit/%H"
```
もし対象になるプロジェクトが違う場合は自前に config を書き換えておかないといけないので注意。
2. 好きな方法で git2pivotal を呼び出す
  * git2pivotal を手動で起動する
  * post-commit フックで呼ぶ
  * など
