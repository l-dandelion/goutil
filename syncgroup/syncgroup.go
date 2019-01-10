package syncgroup

import (
	"sync"
)

/*SyncGroup 并发控制器
 *用于控制并发数量、等待所有goroutine执行结束
 */
type SyncGroup interface {
	Add()
	Done()
	Wait()
	Num() int
}

type mySyncGroup struct {
	Pool chan bool      // 控制数量
	Wg   sync.WaitGroup // 用于等待goroutine执行完毕
}

/*New 新建一个sg
 */
func New(maxNum int) SyncGroup {
	sg := &mySyncGroup{}
	sg.init(maxNum)
	return sg
}

/*Add 增加goroutine
 */
func (sg *mySyncGroup) Add() {
	sg.Pool <- true
	sg.Wg.Add(1)
}

/*Done 完成goroutine
 */
func (sg *mySyncGroup) Done() {
	<-sg.Pool
	sg.Wg.Done()
}

/*Wait 等待所有goroutine执行完毕
 */
func (sg *mySyncGroup) Wait() {
	sg.Wg.Wait()
}

/*init 初始化最大goroutine数量
 */
func (sg *mySyncGroup) init(maxNum int) {
	sg.Pool = make(chan bool, maxNum)
}

/*Num 已分配的goroutine数量
 */
func (sg *mySyncGroup) Num() int {
	return len(sg.Pool)
}
