### 使用swagger


#### 安装包

```shell
go get -u github.com/swaggo/swag  
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files

```


#### 生成文件
```shell
swag init

```


linux系统下出现
> bash: swag: command not found

执行以下命令
```
export PATH=$(go env GOPATH)/bin:$PATH
```