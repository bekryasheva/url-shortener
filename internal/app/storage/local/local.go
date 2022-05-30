package local

import (
	"github.com/bekryasheva/url-shortener/pkg"
	"sync"
)

type LocalStorage struct {
	counter int64
	keyURL  map[string]int64
	keyID   map[int64]string
	mu      sync.RWMutex
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		counter: 0,
		keyURL:  make(map[string]int64),
		keyID:   make(map[int64]string),
		mu:      sync.RWMutex{},
	}
}

func (l *LocalStorage) Save(originalURL string) (int64, error) {
	var id int64

	l.mu.Lock()
	if _, ok := l.keyURL[originalURL]; ok {
		id = l.keyURL[originalURL]
	} else {
		l.counter++
		id = l.counter
		l.keyURL[originalURL] = id
		l.keyID[id] = originalURL
	}
	l.mu.Unlock()

	return id, nil
}

func (l *LocalStorage) Get(id int64) (string, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if _, ok := l.keyID[id]; !ok {
		return "", pkg.ErrNotFound
	}

	return l.keyID[id], nil
}
