package volatile

import (
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	m := Map(time.Millisecond)
	m.Set("key", "value")
	if v := m.Get("key"); v != "value" {
		t.Errorf("%v didn't equal value", v)
	}
	temporaryCleanUpPeriod = 2 * time.Millisecond
	time.Sleep(3 * time.Millisecond)
	if len(m.lastSeen) != 0 {
		t.Error("didn't clean up lastSeen automatically")
	}
	if len(m.Elements) != 0 {
		t.Error("didn't clean up Elements automatically")
	}
}
