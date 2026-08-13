# VirusTotal API可視化ツール（実験版）
* VirusTotal API実行結果をReactで可視化させる

> [!NOTE]
> **フィードバック＆機能リクエスト募集中**  
> 本ツールは、VirusTotal APIを活用した可視化の需要や動向を調査するために開発した**実験的ツール**です。
> 現在はまだ機能開発の途上ですが、「こんな表示モードがほしい」「このデータをグラフ化してほしい」などのアイデアや改善案があれば Issue に起票お願いします。

## API 項目
* IP調査：`Get an IP address report`
* ドメイン調査：`Get a domain report`
* ファイル調査：`Get a summary of all behavior reports for a file`

# ツール利用時の注意
* **APIプランと利用条件**: 本ツールはVirusTotal APIを利用します。APIの無料プランには、利用目的、商用利用、リクエスト数、提供機能に関する制限があります。利用者は、自身のAPIキーに適用される利用規約、API規約、料金プラン、レート制限を各サービスの公式サイトで確認し、遵守してください。本ツールは無料利用枠またはAPIの継続提供を保証しません。
* **調査対象データの送信**: IPアドレス、ドメイン、ファイルハッシュなどの調査対象データは、外部APIへ送信される場合があります。送信前に、対象データを外部サービスへ送信する権限、および契約・組織規程上の取扱条件を確認してください。
* **APIキーの管理**: APIキーは`.env`などの追跡されない設定ファイルに保存してください。APIキーをリポジトリ、Issue、Pull Request、ログ、画面共有へ記載しないでください。
* **法令と利用条件の遵守**: 本ツールは、法令、契約、組織規程、および対象システムと外部APIの利用条件に従って使用してください。権限のない調査、アクセス、データ送信に使用しないでください。

# 技術スタック
<img src="https://img.shields.io/badge/DockerEngine-29.5.3-2496ED.svg?logo=docker">
<img src="https://img.shields.io/badge/Ubuntu-wsl2-FF570A.svg?logo=ubuntu">

<p>
<img src="https://img.shields.io/badge/React-19.2.6-61DAFB.svg?logo=react">
<img src="https://img.shields.io/badge/ReactIcons-5.6.0-CB3837.svg?logo=react">
<img src="https://img.shields.io/badge/xyflow（react flow）-12.11.0-61DAFB.svg?logo=xyflow">
<img src="https://img.shields.io/badge/TypeScript-6.0.3-3178C6.svg?logo=typescript">
<img src="https://img.shields.io/badge/Node.js-24.14.1-5FA04E.svg?logo=nodedotjs">
</p>

<p>
<img src="https://img.shields.io/badge/Go-1.26.3-00ADD8.svg?logo=go">
</p>

<p>
<img src="https://img.shields.io/badge/Git-2.43.0-F03C2E.svg?logo=git">
<img src="https://img.shields.io/badge/GitHub-181717.svg?logo=github">
</p>

<img src="https://img.shields.io/badge/CodeRabbit-CodeReview AI Agent-FF570A.svg?logo=coderabbit">

# 初回設定  
`.env.example`に記載されているAPIキーを設定
## APIの利用用途

| KEY | IP調査 | ドメイン調査 | ファイル調査 |
| --- | --- | --- | --- |
| VIRUS_TOTAL | 〇 | 〇 | 〇 |  
  
※ 将来的に複数のレピュテーションサイトが提供するAPIを実装予定

# 1. 開発方法
* `.env`ファイルを`.env.example`と同じパスに用意しキーを用意する
# 2. 開発環境の起動
* `make up` コマンドを実行することでフロントエンド環境とバックエンド環境が起動する  
`localhost:5173` でOSINTホーム画面にアクセス可能
# 3. コンテナを終了する
* `make down` コマンドを実行することで終了

# IP調査・ドメイン調査
```
# IPアドレス検索処理
localhost:4000/host/report?host=1.1.1.1

# ドメイン検索処理
localhost:4000/host/report?host=example.com

# ファイル検索処理
localhost:4000/file/report?hash=sha-256
```  

# Dockerで可視化ツールを利用する場合
`make prod-up` を実行することでビルドも合わせて行われる  
React, Go 1つのファイルに集約されるため起動時は `localhost:4000` でアクセス可能

# ラズパイ初回展開手順
ラズパイに展開したい場合は以下の手順で対応
1. `make build`コマンドを実行
2. `buildfile`ディレクトリに生成された`rasp-arm64`を`C:`など展開しやすいフォルダに移動
3. `raspberry pi`に SSHコマンド を実行しアクセス
4. コマンドプロンプトを起動し`scp <ビルドファイルパス> <raspbeerypi user>@<ip address>:/home/<user>`コマンドでビルドファイルを展開
5. `.env`ファイルを`home/<user>`に作成し各APIキーを設定
6. `APP_ENV=prod ./rasp-arm64`を実行
7. `http://<raspberrypi ip address>:4000`でアクセス

## ラズパイ展開後実行エラーとなる場合
1. `sudo ufw status`コマンドを実行しポート4000の許可設定がされていない場合は許可設定を行う
2. ポート開放は`sudo ufw allow 4000/tcp`を実行
3. 設定後`sudo ufw reload`を実行しアクセスできるか確認

# コーディング規約
* Go: `docs/gogideline.md`を参照
* React: 未作成

# ブランチ名
* 開発・新規機能: `feat/ticket-<number>`
* 修正: `ref/ticket-<number>`
* バグ: `bug/ticket-<number>`

# 免責事項
* **API仕様・提供形態の変更**: 各APIの仕様変更、無料枠の制限変更、またはサービスの突然の停止・中断により本ツールが利用不能になった場合でも、開発者は修正や動作保証の義務を負いません。
* **アカウント等の停止**: 本ツールの利用によりユーザーのAPIキー、アカウント、あるいはIPアドレスが各サービスから制限・停止（BAN）された場合、開発者は一切の責任を負いません。
* **データの正確性と判定結果**: 各APIから取得されるデータおよび判定結果（危険性の有無など）の正確性・完全性は保証いたしません。データの誤判定（誤検知・見落とし）によって生じた不利益や損害について、開発者は一切の責任を負いません。

# サードパーティ
* [React Flow](https://reactflow.dev/) - A customizable React component for building node-based UI.
* [VirusTotal API v3](https://docs.virustotal.com/reference/overview) - VirusTotal API v3 Overview