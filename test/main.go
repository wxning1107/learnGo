package main

import (
	"fmt"
	"time"
)

type In interface {
	start() error
}

var shell = `

top - 17:08:49 up 55 days, 21:50,  3 users,  load average: 0.03, 0.03, 0.05
Tasks: 171 total,   1 running, 170 sleeping,   0 stopped,   0 zombie
%Cpu(s):  0.5 us,  0.2 sy,  0.0 ni, 99.2 id,  0.0 wa,  0.0 hi,  0.2 si,  0.0 st
KiB Mem :  3880792 total,   982784 free,  1258588 used,  1639420 buff/cache
KiB Swap:   511996 total,   491160 free,    20836 used.  2185504 avail Mem 

  PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND                                            
16086 influxdb  20   0  501288  84036   7884 S   0.7  2.2 122:15.45 influxd                                            
 5581 root      20   0  300936   3240   2836 S   0.3  0.1  72:01.95 vmtoolsd                                           
16393 root      20   0 5000052 996.5m  15024 S   0.3 26.3 151:43.61 java                                               
    1 root      20   0  128124   4540   2692 S   0.0  0.1   1:37.41 systemd                                            
    2 root      20   0       0      0      0 S   0.0  0.0   0:01.05 kthreadd                                           
    3 root      20   0       0      0      0 S   0.0  0.0   0:30.41 ksoftirqd/0                                        
    5 root       0 -20       0      0      0 S   0.0  0.0   0:00.00 kworker/0:0H                                       
    7 root      rt   0       0      0      0 S   0.0  0.0   0:00.50 migration/0                                        
    8 root      20   0       0      0      0 S   0.0  0.0   0:00.00 rcu_bh                                             
    9 root      20   0       0      0      0 S   0.0  0.0  19:05.18 rcu_sched                                          
   10 root       0 -20       0      0      0 S   0.0  0.0   0:00.00 lru-add-drain                                      
   11 root      rt   0       0      0      0 S   0.0  0.0   0:14.90 watchdog/0                                         
   12 root      rt   0       0      0      0 S   0.0  0.0   0:24.47 watchdog/1                                         
   13 root      rt   0       0      0      0 S   0.0  0.0   0:00.49 migration/1                                        
   14 root      20   0       0      0      0 S   0.0  0.0   0:20.60 ksoftirqd/1                                        
   16 root       0 -20       0      0      0 S   0.0  0.0   0:00.00 kworker/1:0H                                       
   18 root      20   0       0      0      0 S   0.0  0.0   0:00.00 kdevtmpfs 

`

func main() {
	//fmt.Println(shell)
	//part := strings.Split(strings.Trim(shell, "\n"), "\n")[2]
	//fmt.Printf("%q\n", part)
	//fmt.Printf("%q\n", strings.Split(part, " "))
	//fmt.Println(strings.Split(strings.Trim(shell, "\n"), " ")[2])
	//fmt.Println(strings.Replace(shell, " ", "", -1))
	//fmt.Println(strings.Trim(shell, "\n"), "\n")
	//data := strings.Split(strings.Trim(shell, "\n"), "\n")
	//for _, v := range data {
	//	fmt.Println(v)
	//}
	//c := make(chan int, 5)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		c <- i
	//	}
	//	close(c)
	//}()
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//}()
	//
	//time.Sleep(time.Millisecond)

	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
		close(ch)
	}()
	time.Sleep(time.Second * 10)
	v := <-ch
	fmt.Println(v)
}
