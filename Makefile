run:
	go run main.go

# 查看 Go 文档
doc:
	godoc -http=:6060

.PHONY: run doc