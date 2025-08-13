package cache

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type cacheItem struct {
	Value      string    `json:"value"`
	Expiration time.Time `json:"expiration"`
}

type memory struct {
	data  map[string]cacheItem
	mutex sync.RWMutex
}

func New() Cache {
	return &memory{
		data: make(map[string]cacheItem),
	}
}

func (m *memory) Get(key string, result any) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	item, exists := m.data[key]
	if !exists {
		return fmt.Errorf("key not found: %s", key)
	}

	if !item.Expiration.IsZero() && time.Now().After(item.Expiration) {
		m.mutex.RUnlock()
		m.mutex.Lock()
		delete(m.data, key)
		m.mutex.Unlock()
		m.mutex.RLock()

		return fmt.Errorf("key not found: %s", key)
	}

	return json.Unmarshal([]byte(item.Value), &result)
}

func (m *memory) Set(key string, value any, ttl time.Duration) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var expTime time.Time
	if ttl > 0 {
		expTime = time.Now().Add(ttl)
	}

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	m.data[key] = cacheItem{
		Value:      string(jsonValue),
		Expiration: expTime,
	}

	return nil
}

func (m *memory) Delete(key string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.data, key)
	return nil
}
