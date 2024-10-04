# flo-go
app web golang

Start bdd

```bash
docker compose -f docker-compose.yaml -p flo-go up -d
docker compose -f docker-compose.yaml -p flo-go down -d
```

Start backend

``` bash
cd backend/cmd
go build cmd/main.go
./main
```

Start front

``` bash
cd front
pnpm install -r
cd app
pnpm run dev
```
