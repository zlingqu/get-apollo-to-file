从apollo拉取配置，并以固定格式追加到目标文件尾部。

##### 使用方法
```bash
-appId string
        Apollo应用的appId
  -clusterName string
        Apollo的集群名
  -configServerUrl string
        Apollo的http地址
  -destFilePath string
        目的文件路径
  -namespaceName string
        Apollo的namespaceName
```
##### 例如：
```
go run main.go \
-configServerUrl=http://prd-conf.apollo.cc.**.com \
-appId=service-adp-env \
-clusterName=11111 \
-namespaceName=application \
-destFilePath=a.txt
```


##### 用途：
当我使用Dockerfile制作镜像时，不同的环境可能需要注入不同的环境变量，此时可以将环境变量存到apollo中，然后使用此工具将数据get下来，按照Dockerfile格式追加到文件尾部。