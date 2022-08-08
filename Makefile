run:
	go run main.go

# 自动热重载，前提需要安装 air `go install github.com/cosmtrek/air@latest`
air:
	air

# 查看 Go 文档
doc:
	godoc -http=:6060

# -v 显示详情
# -count=1 清楚缓存
test:
	go test ./tests -v -count=1

.PHONY: run air doc