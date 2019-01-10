package syncgroup

import (
	"fmt"
	"testing"
	"time"
)

func TestSyncGroup(t *testing.T) {
	sg := New(2)
	for i := 0; i < 4; i++ {
		// 申请资源
		sg.Add()
		go func(x int) {
			// 释放资源
			defer sg.Done()
			fmt.Println(x)
			time.Sleep(5 * time.Second)
		}(i)
	}
	// 等待所有协程结束
	sg.Wait()
	fmt.Println("Finish!!")
}
