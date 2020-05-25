package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

const (
	AUTH_USER = "cc"
	AUTH_PWD  = "3.0#bkcc"
)

type ZKLock struct {
	zkHost []string
	zkConn *zk.Conn
	zkAcl  []zk.ACL
	zkLock *zk.Lock
}

type ZKLocker interface {
	Lock(path string) error
	Unlock() error
}

func NewZkLock(host []string) ZKLocker {
	return &ZKLock{
		zkHost: host[:],
		zkConn: nil,
		zkAcl:  zk.DigestACL(zk.PermAll, AUTH_USER, AUTH_PWD),
		zkLock: nil,
	}
}

func (zkLock *ZKLock) Lock(path string) error {
	return zkLock.LockSlow(path, time.Second*5)
}

func (zkLock *ZKLock) LockSlow(path string, sessionTimeOut time.Duration) error {
	if zkLock.zkConn == nil {
		conn, _, connErr := zk.Connect(zkLock.zkHost, sessionTimeOut)
		if connErr != nil {
			return connErr
		}

		// auth
		auth := AUTH_USER + ":" + AUTH_PWD
		if connErr := conn.AddAuth("digest", []byte(auth)); connErr != nil {
			conn.Close()
			return connErr
		}

		zkLock.zkConn = conn
	}

	lock := zk.NewLock(zkLock.zkConn, path, zkLock.zkAcl)
	if lock == nil {
		return fmt.Errorf("fail to new lock for path(%s)", path)
	}

	zkLock.zkLock = lock

	return zkLock.zkLock.Lock()
}

func (zkLock *ZKLock) Unlock() error {
	if zkLock.zkLock != nil {
		if unlock := zkLock.zkLock.Unlock(); unlock != nil {
			zkLock.zkConn.Close()
			return unlock
		}
	}

	if zkLock.zkConn != nil {
		zkLock.zkConn.Close()
	}

	return nil
}

var counter int

func main() {
	lock := NewZkLock([]string{"127.0.0.1"})
	lock.Lock("qqq")
	lock.Unlock()
}
