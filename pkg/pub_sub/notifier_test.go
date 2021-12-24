package pub_sub

import (
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_notifier_Subscribe(t *testing.T) {
	tf := func() {}

	n := &notifier{
		RWMutex: *new(sync.RWMutex),
		s:       make([]func(), 0),
	}
	n.Subscribe(tf)
	assert.Equal(t, reflect.ValueOf(tf).Pointer(), reflect.ValueOf(n.s[0]).Pointer())
}

func Test_notifier_Notify(t *testing.T) {
	sig := make(chan struct{}, 1)
	tf := func() {
		close(sig)
	}

	n := &notifier{
		RWMutex: *new(sync.RWMutex),
		s:       make([]func(), 0),
	}
	n.Subscribe(tf)
	n.Notify()
	<-sig
}
