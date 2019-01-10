# 阿里云物联网的golang版sdk

参考aliyun官方的js版`aliyun-iot-device-sdk`

可以用于开发私人网关，架设在路由器中运行，支持更新和发布属性。

## Sample

``` golang
    dev := iot.NewDevice(cnf)
    if err := dev.Connect(); err != nil {
        logger.Log("error on connecting:", err.Error())
        return
    }
    logger.Log("msg", "connected.")
    if err := dev.SubscribePropertyMessage(onReceiveMessage); err != nil {
        logger.Log("error on subscribe property message:", err.Error())
        return
    }
    logger.Log("msg", "subscrib property message.")
```

更多的用例参考sample文件夹

## Todo

- [x] 和云端建立连接
- [x] 通过云端验证
- [x] 上报物模型属性
- [x] 监听物模型属性
- [] 上报物模型事件
- [x] 上报自定义topic
- [x] 订阅自定义topic
