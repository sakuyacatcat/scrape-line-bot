# 使用するベースイメージ
FROM python:3.9-slim

# 作業ディレクトリの指定
WORKDIR /app

# 依存関係のファイルをコンテナにコピー
COPY requirements.txt .

# 依存関係のインストール
RUN pip install --no-cache-dir -r requirements.txt

# ソースコードをコンテナにコピー
COPY . .

# ポート番号の指定
EXPOSE 8888

# コンテナ起動時のコマンド
CMD ["python", "line_bot.py"]
