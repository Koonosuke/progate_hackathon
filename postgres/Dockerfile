FROM postgres:15

# 必要なパッケージをインストールして pgvector をビルド
RUN apt-get update && \
    apt-get install -y git build-essential postgresql-server-dev-15 && \
    git clone --branch v0.5.0 https://github.com/pgvector/pgvector.git && \
    cd pgvector && \
    make && \
    make install && \
    cd .. && rm -rf pgvector && \
    apt-get remove -y git build-essential postgresql-server-dev-15 && \
    apt-get autoremove -y && \
    apt-get clean

CMD ["postgres"]
