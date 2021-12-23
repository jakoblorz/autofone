package pkg

import (
	"sync"
)

type SubscriptionReceiverNotifier interface {
	SubscriptionNotifier
	SubscriptionReceiver
}

type SubscriptionReceiver interface {
	Subscribe(func())
}

type SubscriptionNotifier interface {
	Notify()
}

type notifier struct {
	sync.RWMutex
	s []func()
}

func (n *notifier) Subscribe(handlerFunc func()) {
	n.Lock()
	defer n.Unlock()

	n.s = append(n.s, handlerFunc)
}

func (n *notifier) Notify() {
	n.RLock()
	defer n.RUnlock()
	for _, fn := range n.s {
		go fn()
	}
}
