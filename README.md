# zhihuimage

> 既然产品经理允许钓鱼，我们就不能浪费这些鱼

安装go,配置好环境后执行
```bash
go get github.com/fennecs/zhihuimage
```
切到$GOPATH/bin/, 执行
```bash
./zhihuimage -h
```
查看帮助,比如<https://www.zhihu.com/question/28997505/answer/515804330> 这条链接，questionId是**263952082**
```bash
# unix 系统
./zhihuimage -d '/root/zhihu' -i 263952082
```
采取分页下载，每页最多(size)为5条回答;最多下载无数页，可以通过"-l"限定页数