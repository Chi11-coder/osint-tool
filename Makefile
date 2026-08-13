# Makeコマンドを実行する場合はルートディレクトリから実行

# Goコマンドは backend ディレクトリ配下で実行されます
# Go
.PHONY: front-check goup go-check gen up prod-up down ps build
goup:
	cd backend && go run -tags dev ./cmd
go-check:
	cd backend && go mod tidy && go fmt ./... && go fix ./...
	cd backend && go test ./... -cover -race
front-check:
	cd frontend && npx prettier --write "src/**/*" && npm run lint && npm test -- --run
gen:
	cd backend && go generate ./...

# docker-compose
up:
	docker compose build --no-cache
	docker compose up --watch

# docker prod
prod-up:
	docker compose -f docker-compose.prod.yaml build --no-cache
	docker compose -f docker-compose.prod.yaml up -d

# 共通
down:
	docker compose down --remove-orphans
	docker compose -f docker-compose.prod.yaml down --remove-orphans

ps:
	docker compose ps -a
	docker container ls -a

# ビルドファイル
build:
	mkdir -p buildfile
	rm -rf backend/cmd/dist
	cd frontend && npm run build && cp -rf dist ../backend/cmd
	cd backend && CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -tags=prod -v -o ../buildfile/rasp-arm64 ./cmd
	cd backend && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -tags=prod -v -o ../buildfile/win-amd64.exe ./cmd
