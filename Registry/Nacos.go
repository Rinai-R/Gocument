package Registry

import (
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var Client naming_client.INamingClient

func init() {
	// 创建ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			"192.168.195.129",                  // Nacos服务器的IP地址
			8848,                               // Nacos服务器的端口号
			constant.WithContextPath("/nacos"), // 上下文路径
		),
	}

	// 创建ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("013f0003-2f55-485a-ba2d-c14e6b888e6e"), // 命名空间ID
		constant.WithTimeoutMs(5000),                                     // 超时时间（毫秒）
		constant.WithNotLoadCacheAtStart(true),                           // 启动时不加载缓存
		constant.WithLogDir("/tmp/nacos/log"),                            // 日志目录
		constant.WithCacheDir("/tmp/nacos/cache"),                        // 缓存目录
		constant.WithLogLevel("debug"),                                   // 日志级别
		constant.WithUsername("nacos"),                                   // 用户名
		constant.WithPassword("nacos"),                               // 密码
	)
	var err error
	// 创建配置客户端
	Client, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		Logger.Logger.Panic(err.Error()) // 如果客户端创建失败，则抛出异常
	}
	Logger.Logger.Debug("Nacos init success")
}
