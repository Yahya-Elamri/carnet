run:
	@TEMPL_EXPERIMENT=rawgo templ generate
	@go run cmd/main.go

up:
	@go run migration/cmd/main.go -direction=up

down:
	@go run migration/cmd/main.go -direction=down
