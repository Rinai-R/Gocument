# 介绍

----

该demo主要实现了：
- 用户
  - 注册
  - 登陆
  - 修改信息
  - 个人主页
- 文档
  - 创建
  - 权限管理
  - 文档删除
  - 文档搜索
  - 文档编辑(采取websocket)
----
涉及到的技术栈：
- Hertz: 采取hertz框架
- MySQL: 加密存储用户信息，存储文档的基本信息(不包括内容)
- Redis: 对用户的个人主页，权限验证等信息进行了缓存，并且保证了数据一致性。
- Elasticsearch: 存储了文档的所有信息，目的是实现文档检索的功能，同时存储了敏感词汇(可自行在yaml文件添加),实现敏感词审查。
- Nacos: 使用nacos作为服务注册中心。
- gRPC: 采取gRPC进行通信
- gorm: 使用gorm来对数据库进行操作
- zap: zap进行日志管理
- viper: viper配置管理
---
## 结构梳理
- App: 实现了api接口以及token中间件。
- Logger: 日志初始化，以及日志的存储位置。
- models: 包括用户表，文档表，一系列关系表(权限关系)的结构体表示，以及websocket的连接情况。
- Registry: 包括注册中心的初始化，相关的服务注册，删除的函数实现。
- Server: 服务端，内有相关的服务实现逻辑以及dao层函数。
- Utils: 错误码/返回json信息/密码加密函数/其他工具函数的实现。
----
## 其他
- 虽然从这次的作业中了解到了很多新东西，但都是蜻蜓点水。
- 读了一下学长的年总结，锐评了去年自己的寒假作业，哈哈，感觉自己也被狠狠锐评了😭😭😭，本来也想用docker-compose来一键启动的，但是nacos的部署还涉及数据库的操作，我也不知道咋解决，就作罢了.
- 这次作业的不足之处还有很多，比如最开始我是把数据库的操作单独出来的，后面发现应该把dao层放进server里面比较好，就copy了一份分开存放了~，然而都还是用的一个数据库🤣，纯纯是为了好看而改动。
- 其他应该没啥想说的了~
- 20250309-更新了Auth服务，使用非对称加密保证安全性