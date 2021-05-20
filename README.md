从apollo拉取配置，并以固定格式追加到目标文件尾部。

使用方法
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
例如：
```
go run main.go -configServerUrl=http://prd-conf.apollo.cc.dm-ai.cn -appId=service-adp-env -clusterName=11111 -namespaceName=application -destFilePath=a.txt
```