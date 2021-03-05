package volatile

import "time"

var temporaryCleanUpPeriod = 1 * time.Minute

type TtlMap struct {
	Elements  map[interface{}]interface{}
	lastSeen  map[interface{}]time.Time
	expiresIn time.Duration
}

func Map(expiresIn time.Duration) *TtlMap {
	m := &TtlMap{
		Elements:  map[interface{}]interface{}{},
		lastSeen:  map[interface{}]time.Time{},
		expiresIn: expiresIn,
	}

	go func(m *TtlMap) {
		for range time.Tick(temporaryCleanUpPeriod) {
			for key := range m.Elements {
				m.cleanup(key)
			}
		}
	}(m)

	return m
}

func (t *TtlMap) Set(key interface{}, value interface{}) {
	t.Elements[key] = value
	t.lastSeen[key] = time.Now()
}

func (t *TtlMap) Get(key interface{}) interface{} {
	t.cleanup(key)
	return t.Elements[key]
}

func (t *TtlMap) cleanup(key interface{}) {
	if t.lastSeen[key].Add(t.expiresIn).Before(time.Now()) {
		delete(t.Elements, key)
		delete(t.lastSeen, key)
	}
}
