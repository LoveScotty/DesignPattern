package singleton

import "sync"

// 单例模式
type singleton struct{}

var ins1 = &singleton{}

// GetIns01 饿汉式: 提前准备
func GetIns01() *singleton {
	return ins1
}

var (
	mu   sync.Mutex
	ins2 *singleton
)

// GetIns02 懒汉式: 需要时加载, 会有并发问题, 需要加锁
func GetIns02() *singleton {
	if ins2 == nil {
		mu.Lock()
		defer mu.Unlock()
		// double check
		if ins2 == nil {
			ins2 = &singleton{}
		}
	}

	return ins2
}

var (
	once sync.Once
	ins3 *singleton
)

// GetIns03 懒汉式: 需要时加载, 会有并发问题, 直接使用sync包, 内部实现也是加锁
func GetIns03() *singleton {
	once.Do(func() {
		ins3 = &singleton{}
	})

	return ins3
}
