package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestZkLock(t *testing.T) {
	fmt.Println("TEST zk lock")
	zkLock1 := NewZkLock([]string{"127.0.0.1:2181"})

	if err := zkLock1.Lock("/lock"); err != nil {
		fmt.Printf("lock1 fail lock. err:%s \n", err.Error())
	} else {
		fmt.Println("lock1 lock")
	}

	unlock := zkLock1.Unlock()
	require.NoError(t, unlock)

	fmt.Println("lock1 unLocked")

}
