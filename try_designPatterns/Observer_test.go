package try_designPatterns

import (
	"fmt"
	"sync"
	"testing"
)

// Subject 主题接口，定义了添加、删除和通知观察者的方法
type Subject interface {
	AddObserver(observer ...Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(message string)
}

// Observer 观察者接口，定义了更新方法
type Observer interface {
	Update(message string)
}

type observer struct {
}

func (o observer) Update(message string) {
	fmt.Println(message)
}

// ConcreteSubject 具体的主题实现
type ConcreteSubject struct {
	observers []Observer
	mutex     sync.Mutex
}

func (c *ConcreteSubject) AddObserver(observer ...Observer) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.observers = append(c.observers, observer...)
}

func (c *ConcreteSubject) RemoveObserver(observer Observer) {
	for i := 0; i < len(c.observers); i++ {
		if c.observers[i] == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			return
		}
	}
}

func (c *ConcreteSubject) NotifyObservers(message string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for _, i := range c.observers {
		i.Update(message)
	}
}

func TestObserver(t *testing.T) {
	var subject Subject = &ConcreteSubject{}
	o1 := observer{}
	o2 := observer{}
	o3 := observer{}

	subject.AddObserver(o1, o2, o3)
	subject.NotifyObservers("hello world")
	subject.RemoveObserver(o1)
	subject.NotifyObservers("hello world222")
}
