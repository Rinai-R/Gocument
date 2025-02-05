# Gocument

**Gocument** 是一个**在线文档协作平台**，目前正在完善中...
运行环境需要nacos/redis/mysql/elasticsearch，地址/用户名/密码需自行配置。

## 实现功能

- [x] 用户注册登陆
- [x] 登陆状态创建/编辑文档
- [x] 文档存储于数据库/存储优化
- [x] 文档分享/文档的私有和公开
- [x] 个人主页(查看我的文档及其基本信息)
- [x] 删改文档

----

- [x] 密码加密
- [ ] 自动保存(高性能)
- [ ] 验证码
- [x] 敏感词审查
- [ ] 多人实时协作
- [x] 文档权限管理
- [ ] 团队/群组管理
- [ ] 评论功能(带有消息通知)
- [ ] 版本管理
- [ ] 文档目录
- [ ] 富文本扩展(支持md/图片/公式等格式)
- [x] 搜索功能
- [ ] 部署
- [x] 缓存
- [ ] 安全(XSS/SQL注入/CSRF等)
- [ ] 高阶挑战

-----

## 接口文档

[点我查看接口文档](https://yzgun2n454.apifox.cn/)

[自定义错误码详解](https://github.com/Rinai-R/Gocument/blob/main/Utils/Error/ErrCode/ErrorCode.go)

----

## 项目结构

```bat
Gocument
│
├───App
│   │   main.go
│   │
│   ├───Func
│   │   ├───Document
│   │   │   ├───Api
│   │   │   │       Create.go
│   │   │   │       DeleteDocument.go
│   │   │   │       Enter.go
│   │   │   │       Grant.go
│   │   │   │       Search.go
│   │   │   │
│   │   │   └───Client
│   │   │       │   DocumentClient.go
│   │   │       │
│   │   │       └───rpc
│   │   │               code
│   │   │               document.pb.go
│   │   │               document.proto
│   │   │               document_grpc.pb.go
│   │   │
│   │   └───User
│   │       ├───Api
│   │       │       Alter.go
│   │       │       Login.go
│   │       │       PersonalPage.go
│   │       │       Register.go
│   │       │
│   │       └───Client
│   │           │   UserClient.go
│   │           │
│   │           └───rpc
│   │                   code
│   │                   user.pb.go
│   │                   user.proto
│   │                   user_grpc.pb.go
│   │
│   ├───Middleware
│   │       jwt.go
│   │       Token.go
│   │
│   └───Router
│           router.go
│
├───Logger
│   │   Logger.go
│   │
│   └───log
│           logger.log
│
├───models
│       Document.go
│       Permission.go
│       User.go
│       websocket.go
│
├───Registry
│   │   init.go
│   │
│   ├───conf
│   │   └───Registry
│   │           init.go
│   │           registry.yaml
│   │
│   └───Nacos
│           Func.go
│           nacos.go
│
├───Server
│   ├───Document
│   │   │   main.go
│   │   │
│   │   ├───DataBase
│   │   │   ├───conf
│   │   │   │   └───DB
│   │   │   │           db.yaml
│   │   │   │           init.go
│   │   │   │
│   │   │   ├───dao
│   │   │   │       Check.go
│   │   │   │       Create.go
│   │   │   │       Delete.go
│   │   │   │       Edit.go
│   │   │   │       Get.go
│   │   │   │       Grant.go
│   │   │   │       Search.go
│   │   │   │       utils.go
│   │   │   │
│   │   │   └───DB
│   │   │       │   init.go
│   │   │       │
│   │   │       ├───ElasticSearch
│   │   │       │       elasticsearch.go
│   │   │       │       Sensitive.go
│   │   │       │
│   │   │       ├───MySQL
│   │   │       │       mysql.go
│   │   │       │
│   │   │       └───Redis
│   │   │               redis.go
│   │   │
│   │   ├───rpc
│   │   │       code
│   │   │       document.pb.go
│   │   │       document.proto
│   │   │       document_grpc.pb.go
│   │   │
│   │   └───service
│   │           Check.go
│   │           Create.go
│   │           Delete.go
│   │           Edit.go
│   │           Get.go
│   │           Grant.go
│   │           Search.go
│   │
│   └───User
│       │   main.go
│       │
│       ├───DataBase
│       │   ├───conf
│       │   │   └───DB
│       │   │           db.yaml
│       │   │           init.go
│       │   │
│       │   ├───dao
│       │   │       Alter.go
│       │   │       LoginAndRegister.go
│       │   │       PersonalPage.go
│       │   │
│       │   └───DB
│       │       │   init.go
│       │       │
│       │       ├───ElasticSearch
│       │       │       elasticsearch.go
│       │       │       Sensitive.go
│       │       │
│       │       ├───MySQL
│       │       │       mysql.go
│       │       │
│       │       └───Redis
│       │               redis.go
│       │
│       ├───rpc
│       │       code
│       │       user.pb.go
│       │       user.proto
│       │       user_grpc.pb.go
│       │
│       └───service
│               Alter.go
│               LoginAndRegister.go
│               PersonalPage.go
│
└───Utils
    ├───encrypt
    │       encrypt.go
    │
    ├───Error
    │   │   Error.go
    │   │
    │   └───ErrCode
    │           ErrorCode.go
    │
    ├───Rsp
    │       model.go
    │       Rsp.go
    │
    └───utils
            Time.go
```

