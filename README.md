# YTElkProducer

elasticsearch操作库

## 配置

```go
elkConf := elasticsearch.Config{
    Addresses: []string{"elk服务器url"},
    Username:  "用户名",
    Password:  "口令",
}
ytESConfig := conf.YTESConfig{
    ESConf:      elkConf,
    DebugMode:   true,
    IndexPrefix: "index前缀",
    IndexType:   "log",
}
client := NewClient(ytESConfig)
```

## 使用

参见Client接口
