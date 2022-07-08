package main

import "sync"

//使用结构体代替类
type Tool struct {
	values int
}
//建立私有变量
var instance *Tool


var once sync.Once

func GetInstance() *Tool {
	once.Do(func() {
		instance = new(Tool)
	})
	return instance
}

func main()  {
	GetInstance()
}