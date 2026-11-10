package main

import (
	"crypto/rand"
	"errors"
	"sync"
	"time"
)

var ErrNotFound = errors.New("link not found")

type Link struct {
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	Visits    int       `json:"visits"`
	CreatedAt time.Time `json:"created_at"`
}

type Store interface {
	Create(url string) (Link, error)
	Get(code string) (Link, error)
	Hit(code string) (Link, error)
	Delete(code string) error
}

type MemoryStore struct {
	mu    sync.Mutex
	links map[string]*Link
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{links: make(map[string]*Link)}
}

func (s *MemoryStore) Get(code string) (Link, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	l, ok := s.links[code]
	if !ok {
		return Link{}, ErrNotFound
	}
	return *l, nil
}

func (s *MemoryStore) Delete(code string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.links[code]; !ok {
		return ErrNotFound
	}
	delete(s.links, code)
	return nil
}

func (s *MemoryStore) Hit(code string) (Link, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	l, ok := s.links[code]
	if !ok {
		return Link{}, ErrNotFound
	}
	l.Visits++
	return *l, nil
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func newCode(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	for i := range b {
		b[i] = alphabet[int(b[i])%len(alphabet)]
	}
	return string(b)
}

func (s *MemoryStore) Create(url string) (Link, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	code := newCode(6)
	for s.links[code] != nil {
		code = newCode(6)
	}
	l := &Link{Code: code, URL: url, CreatedAt: time.Now().UTC()}
	s.links[code] = l
	return *l, nil
}

var _ Store = (*MemoryStore)(nil)