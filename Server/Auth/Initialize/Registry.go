package Initialize

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type Registry struct {
	client *clientv3.Client
}

var EtcdRegistry *Registry

func InitEtcd() {
	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Panic("初始化注册中心失败 " + err.Error())
	}
	EtcdRegistry = &Registry{client: etcd}
}

func (e *Registry) ServiceRegister(name string, addr string) {
	Lease, err := e.client.Grant(context.Background(), 5)
	if err != nil {
		log.Panic(err)
	}
	_, err = e.client.Put(context.Background(), name, addr, clientv3.WithLease(Lease.ID))
	if err != nil {
		log.Panic(err)
	}
	go e.keepAlive(Lease.ID)
	log.Println(name + " service ok")
}

func (e Registry) keepAlive(leaseID clientv3.LeaseID) {
	// 获取心跳通道
	keepAliveCh, err := e.client.KeepAlive(context.Background(), leaseID)
	if err != nil {
		log.Fatalf("failed to start keep-alive: %v", err)
	}

	// 接收心跳续约响应
	for {
		select {
		case kaResp := <-keepAliveCh:
			if kaResp == nil {
				fmt.Println("Lease expired")
				return
			}
		case <-time.After(6 * time.Second): // 如果没有心跳续约，检查连接是否断开
			fmt.Println("No heartbeats received.")
		}
	}
}
