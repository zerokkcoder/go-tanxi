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

.PHONY: run air doc deploy

REMOTE=117.你的.IP.地址
APPNAME=goblog

deploy:
    @echo "\n--- 开始构建可执行文件 ---"
    GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -o tmp/$(APPNAME)_tmp

    @echo "\n--- 上传可执行文件 ---"
    scp tmp/$(APPNAME)_tmp root@$(REMOTE):/data/www/goblog.com/

    @echo "\n--- 停止服务 ---"
    ssh root@$(REMOTE) "supervisorctl stop $(APPNAME)"

    @echo "\n--- 替换新文件 ---"
    ssh root@$(REMOTE) "cd /data/www/goblog.com/ \
                            && rm $(APPNAME) \
                            && mv $(APPNAME)_tmp $(APPNAME) \
                            && chown www-data:www-data $(APPNAME)"

    @echo "\n--- 开始服务 ---"
    ssh root@$(REMOTE) "supervisorctl start $(APPNAME)"

    @echo "\n--- 部署完毕 ---\n"

