package etcd

import (
	"context"
	"fmt"
	clientV3 "go.etcd.io/etcd/client/v3"
	"log"
	"net"
)

const schema = "etcd"

var client *clientV3.Client
var key string

func Register(host, port, service, target string, ttl int64) error {
	addr := net.JoinHostPort(host, port)
	key = fmt.Sprintf("%s://%s/%s", schema, service, addr)

	var err error
	client, err = clientV3.New(clientV3.Config{
		Endpoints: []string{target},
	})
	if err != nil {
		return err
	}
	resp, err := client.Grant(context.TODO(), ttl)
	if err != nil {
		return err
	}
	_, err = client.Put(context.TODO(), key, addr, clientV3.WithLease(resp.ID))
	if err != nil {
		return err
	}
	ch, err := client.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		return err
	}
	go func() {
		for {
			c := <-ch
			log.Println(key, c.TTL)
		}
	}()
	log.Println(key, "Register")
	return nil
}

func UnRegister() error {
	_, err := client.Delete(context.TODO(), key)
	if err != nil {
		return err
	}
	log.Println(key, "UnRegister")
	return nil
}
