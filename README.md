postgres ログイン
docker exec -it progate_hackathon psql -U progate -d progate

削除起動方法
docker-compose down
docker-compose build
docker-compose up -d

RAG（ベクトル）
既存の PostgreSQL コンテナに pgvector を手動インストール
もし postgres:15 のまま ankane/pgvector を使いたくない場合、手動で pgvector をインストールできます。

1. コンテナにログイン

docker exec -it progate_hackathon bash 2. 必要なパッケージをインストール

apt update && apt install -y git build-essential postgresql-server-dev-15

3. pgvector をダウンロード＆インストール

git clone --branch v0.5.0 https://github.com/pgvector/pgvector.git
cd pgvector
make && make install

4. PostgreSQL に pgvector を追加
   psql -U progate -d progate -c "CREATE EXTENSION vector;"

⑤ テーブルに embedding カラムを追加

psql -U progate -d progate

ALTER TABLE stickies ADD COLUMN embedding VECTOR(1536);
