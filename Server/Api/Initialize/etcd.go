package Initialize

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
	"time"
)

type Registry struct {
	client   *clientv3.Client
	Services map[string]string
	mu       sync.RWMutex
}

var ETCD *Registry

func InitETCD() {
	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.195.129:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Panicln("初始化注册中心失败 " + err.Error())
	}
	ETCD = &Registry{
		client:   etcd,
		Services: make(map[string]string),
		mu:       sync.RWMutex{},
	}
}

func (client *Registry) DiscoverService(serviceName string) error {
	resp, err := client.client.Get(context.Background(), serviceName, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, v := range resp.Kvs {
		client.UpdateService(string(v.Key), string(v.Value))
	}
	// 监听服务的变化
	go func() {
		watchchan := client.client.Watch(context.Background(), serviceName, clientv3.WithPrefix())
		for watch := range watchchan {
			for _, event := range watch.Events {
				switch event.Type {
				case mvccpb.PUT:
					client.UpdateService(string(event.Kv.Key), string(event.Kv.Value))
				case mvccpb.DELETE:
					client.DeleteService(string(event.Kv.Key))
				}
			}
		}
	}()
	return nil
}

func (client *Registry) UpdateService(key string, value string) {
	client.mu.Lock()
	defer client.mu.Unlock()
	client.Services[key] = value
}

func (client *Registry) DeleteService(serviceName string) {
	client.mu.Lock()
	defer client.mu.Unlock()
	delete(client.Services, serviceName)
}

func (client *Registry) GetService(serviceName string) string {
	client.mu.RLock()
	defer client.mu.RUnlock()
	value, ok := client.Services[serviceName]
	if !ok {
		return ""
	}
	return value
}
