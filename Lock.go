package util

import (
	"sync"
)

//UsingRead 读锁
func UsingRead(lk *sync.RWMutex, f func()) {
	lk.RLock()
	defer lk.RUnlock()
	f()
}

//UsingWiter 写锁
func UsingWiter(lk *sync.RWMutex, f func()) {
	lk.Lock()
	defer lk.Unlock()
	f()
}

//释放资源接口
type IDisposable interface {

	//释放资源
	Dispose()
}

//离开作用域时，会释放资源
func Using(obj IDisposable, f func()) {
	defer obj.Dispose()
	f()
}
