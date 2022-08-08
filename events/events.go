package events

import (
	"sync"
)

var lock sync.Mutex

type UpdaterEventHandler func(event *UpdaterEvent, arg interface{})

type UpdaterEvent struct {
	handlers []UpdaterEventHandler
}

func (e *UpdaterEvent) Register(h UpdaterEventHandler) {
	lock.Lock()
	e.handlers = append(e.handlers, h)
	lock.Unlock()
}

func (e *UpdaterEvent) Fire(arg interface{}) {
	for _, h := range e.handlers {
		h(e, arg)
	}
}

var instance = &UpdaterEvent{}

func UpdaterEventInstance() *UpdaterEvent {
	return instance
}
