package etcd

import (
	"context"
	"fmt"
	clientV3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var client *clientV3.Client
var key string
var manager endpoints.Manager

func Register(serviceHost, servicePort, serviceName, etcdTarget string, ttl int64) error {
	addr := net.JoinHostPort(serviceHost, servicePort)
	key = fmt.Sprintf("%s/%s", serviceName, addr)

	var err error
	client, err = clientV3.NewFromURL(etcdTarget)
	if err != nil {
		return err
	}
	manager, err = endpoints.NewManager(client, serviceName)
	if err != nil {
		return err
	}
	lease, err := client.Grant(context.TODO(), ttl)
	if err != nil {
		return err
	}
	err = manager.AddEndpoint(context.TODO(), key, endpoints.Endpoint{
		Addr: addr,
	}, clientV3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	ch, err := client.KeepAlive(context.TODO(), lease.ID)
	if err != nil {
		return err
	}
	go func() {
		for {
			<-ch
		}
	}()
	log.Println(key, "Register")
	exitRegister()
	return nil
}

func UnRegister() error {
	err := manager.DeleteEndpoint(context.TODO(), key)
	if err != nil {
		return err
	}
	log.Println(key, "UnRegister")
	return nil
}

func exitRegister() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		err := UnRegister()
		if err != nil {
			return
		}
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()
}
