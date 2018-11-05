# zhihuimage

> 既然产品经理允许钓鱼，我们就不能浪费这些鱼

安装go,配置好环境后执行
```bash
go get github.com/jinxZz/zhihuimage
```
切到$GOPATH/bin/, 执行
```bash
./zhihu-image -h
```
查看帮助,比如<https://www.zhihu.com/question/28997505/answer/515804330> 这条链接，questionId是**28997505**
```bash
./zhihu-iamge -d '/root/zhihu' -i 28997505
```
采取分页下载，每页最多(size)为5条回答，最多无数条回答，可以通过"-l"限定，10条回答估计有100m图片(高赞回答)