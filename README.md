# CharV-backend
[![codecov](https://codecov.io/gh/CharVstack/CharV-backend/branch/main/graph/badge.svg?token=GJOQH6VXLA)](https://codecov.io/gh/CharVstack/CharV-backend)
CharV-backend は、CharVstack のバックエンドです。

## 開発の環境構築

必要なディレクトリは作成し、環境変数の設定を行います。
`.env.sample` を参考に設定してください。

```sh
ORIGIN_URI=1.2.3.4
IMAGES_DIR=/var/lib/charv/images/
GUESTS_DIR=/var/lib/charv/guests/
STORAGE_POOLS_DIR=/var/lib/charv/storage_pools/
QMP_DIR=/tmp/charv/qmp/
VNC_DIR=/tmp/charv/vnc/
```

- `/var/lib/charv/guests`
  - 作成した VM 情報が格納されます
- `/var/lib/charv/images`
  - `ubuntu-20.04.5-live-server-amd64.iso` をダウンロードし配置してください
- `/var/lib/charv/storage_pools`
  - ストレージプールの情報を記載した JSON を格納してください
  - `CharV-backend/testdata/resources/storage_pools` を参考に JSON を作成してください
- `/tmp/charv/qmp`
  - 電源操作の情報が格納されます
- `/tmp/charv/vnc`
  - VNCの情報が格納されます

## 開発環境の構築

### 依存ツールのインストール

はじめに、依存しているツールをインストールするため、以下のコマンドを実行してください。

```shell
make tools
```

### 開発サーバーの起動

開発用のサーバーを起動する場合、以下のコマンドを実行してください。

```shell
make dev
```

### テストやフォーマットなど

`make` コマンドを使用してテスト, コードフォーマット, 依存関係の更新や最適化, ビルドなどを行えるようになっております。
`make help` を参照してください。

## License

MIT
