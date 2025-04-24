package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	mu        sync.RWMutex
	CreatedAt time.Time
	UpdateAt  time.Time
}

func genereteAPIKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func NewAccount(name, email string) *Account {
	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     uuid.New().String(),
		APIKey:    genereteAPIKey(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
	return account
}

func (a *Account) AddBalance(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	a.UpdateAt = time.Now()
}
