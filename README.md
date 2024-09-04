#基于go语言的高并发分布式游戏服务器

这是一个高性能、可扩展的游戏服务器实现，使用Go1.22开发.

## 项目概述

本项目实现了一个分布式的游戏服务器架构，支持高并发，易于扩展。 主要特点包括：

-分离的登录网关(LoginGate)和游戏服务器(GameServer)
-使用Docker容器化部署
-服务器集群支持，实现负载均衡
-使用Kafka消息队列缓存用户消息
-使用gorroutine池优化并发性能
-Redis作为LoginGate的缓存，减少数据库查询
-使用插座实现客户端与服务端通信
-游戏逻辑与服务器分离，支持独立更新
-跨平台客户端支持(Windows)

## 系统架构

1. **LoginGate**：处理用户注册和登录，负责将用户分发到负载较低的游戏服务器。
2. **游戏服务器**: 处理游戏逻辑，支持多实例部署。
3. **客户端**：跨平台(Windows)支持，通过插座与服务器通信。

## 技术栈

-走1.22
-Docker
-卡夫卡
-Redis
-插座

## 项目结构
```YAML
蛇游戏/
│
├--cmd/
│└--客户端/
│└--main.go
├--客户端/
│├--资产/
││├--图像/
││└--声音/
│├--网络/
││└--客户端.go
│└--ui/
│├--游戏.go
│├--游戏窗口.go
│└--菜单.go
├--驳接器/
│├--logingate/
││└--Dockerfile
│└--服务器/
│└--Dockerfile
├--内部/
│├--博弈/
││├--游戏.go
││├--播放器.go
││└--snake.go
│└--网络/
│├--连接.go
│└--server.go
├--pkg/
│└--config/
│└--config.go
├--脚本/
│├--start_client.sh
│└--start_Servers.sh
├--go.mod
├-go.sum
```


## 运行说明

### 服务器端

1.构建并启动LoginGate：

CD docker/logingate
Docker build-t logingate。
Docker run-d--name logingate logingate



2.构建并启动游戏服务器：

CD docker/服务器
Docker build-t gameserver。
   docker run -d --name gameserver1 gameserver
3. 根据需求启动多个 GameServer 实例。

### 客户端

在 Windows 设备上运行：
go run cmd/client/main.go
## 开发说明

- 游戏逻辑位于 `internal/game/` 目录，可以独立更新而不影响服务器核心功能。
- 客户端代码位于 `client/` 目录。
- 服务器配置可以在 `config.yaml` 文件中修改。

## 注意事项

- 确保 Go 版本为 1.22 或以上。
- 我尽可能使用的是 Go 标准库，以减少第三方依赖。
- 当前游戏实现为简单的贪吃蛇，可以通过修改 `internal/game/` 目录下的文件来替换或扩展游戏逻辑。
- server、kafka和logingate的容器可能会因为启动速度产生冲突，可以手动单独启动一下（暂时还没搞清楚具体因为什么，所以还没做延迟启动）。

## 贡献

欢迎提交 Issue 或 Pull Request 来改进这个项目。

## 游戏运行界面：

![fa50f80febca75b7d0d24032bc669f7](readme_image/fa50f80febca75b7d0d24032bc669f7.png)

![0f1bc639972b97587efe4f7680d7ce8](readme_image/0f1bc639972b97587efe4f7680d7ce8.png)

![f84581ee0bd1c33d372730ab007772f](readme_image/f84581ee0bd1c33d372730ab007772f.png)

![989953293548255a19aacc9201e4684](readme_image/989953293548255a19aacc9201e4684.png)
