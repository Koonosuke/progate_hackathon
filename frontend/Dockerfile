# --- Build Stage ---
    FROM node:20-alpine AS builder
    WORKDIR /app
    
    # 依存関係のインストール
    COPY package.json package-lock.json ./
    RUN npm ci
    
    # 環境変数ファイルを先にコピー（Next.jsビルドで使用）
    COPY .env.production .env.production
    
    # ソースコードのコピー
    COPY . .
    
    # Next.js のビルド（.env.productionを読み込んで実行）
    RUN npm run build
    
    # --- Run Stage ---
    FROM node:20-alpine AS runner
    WORKDIR /app
    
    # 必要なファイルだけをコピー（軽量化）
    COPY --from=builder /app/package.json ./
    COPY --from=builder /app/.next ./.next
    COPY --from=builder /app/public ./public
    COPY --from=builder /app/node_modules ./node_modules
    
    # 本番実行環境
    ENV NODE_ENV=production
    EXPOSE 3000
    
    # Next.js 本番起動
    CMD ["npx", "next", "start"]
    