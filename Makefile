.PHONY: dev
dev:
	wrangler pages dev --local

.PHONY: build
build:
	mkdir -p dist
	tinygo build -o ./functions/api/app.wasm -target wasm ./main.go

.PHONY: publish
publish:
	wrangler publish

.PHONY: generate
generate:
	go generate ./...

.PHONY: create-db
create-db:
	wrangler d1 create zztkm-bbs

.PHONY: create-db-preview
create-db-preview:
	wrangler d1 create zztkm-bbs-preview

.PHONY: init-db
init-db:
	wrangler d1 execute d1-blog-server --file=./schema.sql

.PHONY: init-db-preview
init-db-preview:
	wrangler d1 execute d1-blog-server-preview --file=./schema.sql