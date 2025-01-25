package Registry

import (
	"github.com/Rinai-R/Gocument/Logger"
	conf "github.com/Rinai-R/Gocument/Registry/conf/Registry"
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
			conf.Registry.Addr,                 // Nacos服务器的IP地址
			conf.Registry.Port,                 // Nacos服务器的端口号
			constant.WithContextPath("/nacos"), // 上下文路径
		),
	}

	// 创建ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(conf.Registry.Namespace), // 命名空间ID
		constant.WithTimeoutMs(5000),                      // 超时时间（毫秒）
		constant.WithNotLoadCacheAtStart(true),            // 启动时不加载缓存
		constant.WithLogDir("/tmp/nacos/log"),             // 日志目录
		constant.WithCacheDir("/tmp/nacos/cache"),         // 缓存目录
		constant.WithLogLevel("debug"),                    // 日志级别
		constant.WithUsername(conf.Registry.Username),     // 用户名
		constant.WithPassword(conf.Registry.Password),     // 密码
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
