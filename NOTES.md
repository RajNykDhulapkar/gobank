docker run --name gobank_db -e POSTGRES_PASSWORD=password -d postgres

docker run --name gobank_db -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres

```bash
    docker run -d --name gobank_db \
    -e POSTGRES_DB=gobank_db \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=password \
    -p 5432:5432 \
    postgres:latest
```
