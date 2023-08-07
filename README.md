# MCDR2Kook-golang



> 一个Kook机器人的事件代理服务

本项目创建的初衷是为了更方便地将kook的机器人的事件转发到下游连接的clients中(类似于[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)
的事件上报服务)，采用websocket方式进行网络通信。

### 使用场景

- 作为机器人的代理服务，对接下游的处理服务进行处理

### 使用姿势

1. 下载对应平台的可执行程序，执行

   ```shell
   ./kook_bot
   ```

1. 如果在kook_bot同级目录已经存在配置文件，将会读取对应配置文件，否则将会创建新的配置文件，请修改配置文件中的`token`为自己的机器人的token值

2. 重新执行步骤1，出现以下日志表示已经连接成功，此时去Kook中查看机器人状态应已经处于`online`状态

   ```
   INFO[0001] state change    from=ws_connected to=connected
   ```

### 配置

```yaml
# Websocket服务启动的Host
host: 127.0.0.1
# Http API的监听端口
http_api_port: 9001
# Websocket服务监听的端口
port: 9000
# Kook机器人的Token，可在“Kook开发者平台->应用”中找到对应的机器人Token
token: ""
```

### 其他

本项目使用[golang-bot](https://github.com/idodo/golang-bot)的sdk进行开发

> 本项目目前还在开发中，如果感觉对自己有帮助还请点个star⭐️，有任何建议或者意见欢迎提issue~**