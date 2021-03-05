package volatile_test

import (
	volatile "github.com/elgohr/go-map-ttl"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	m := volatile.Map(time.Millisecond)
	m.Set("key", "value")
	if v := m.Get("key"); v != "value" {
		t.Errorf("%v didn't equal value", v)
	}
	time.Sleep(2 * time.Millisecond)
	if m.Get("key") != nil {
		t.Error("the map didn't forget the value")
	}
	if m.Get("key") != nil {
		t.Error("the map didn't forget the value")
	}
}


