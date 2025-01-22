package Registry

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func RegisterServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	// 调用 RegisterInstance 方法注册服务实例
	success, err := client.RegisterInstance(param)
	if !success || err != nil {
		// 如果注册失败，抛出 panic 并打印错误信息
		panic("RegisterServiceInstance failed!" + err.Error())
	}
	// 打印注册参数和结果
	fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

func DeRegisterServiceInstance(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {
	// 调用 DeregisterInstance 方法取消注册服务实例
	success, err := client.DeregisterInstance(param)
	if !success || err != nil {
		// 如果取消注册失败，抛出 panic 并打印错误信息
		panic("DeRegisterServiceInstance failed!" + err.Error())
	}
	// 打印取消注册参数和结果
	fmt.Printf("DeRegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}
