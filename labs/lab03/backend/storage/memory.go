package storage

import (
	"errors"
	"lab03-backend/models"
	"sync"
)

// MemoryStorage implements in-memory storage for messages
type MemoryStorage struct {
	// TODO: Add mutex field for thread safety (sync.RWMutex)
	// TODO: Add messages field as map[int]*models.Message
	// TODO: Add nextID field of type int for auto-incrementing IDs
	mu       sync.RWMutex
	messages map[int]*models.Message
	nextID   int
}

// NewMemoryStorage creates a new in-memory storage instance
func NewMemoryStorage() *MemoryStorage {
	// TODO: Return a new MemoryStorage instance with initialized fields
	// Initialize messages as empty map
	// Set nextID to 1
	return &MemoryStorage{
		messages: make(map[int]*models.Message),
		nextID:   1,
	}
}

// GetAll returns all messages
func (ms *MemoryStorage) GetAll() []*models.Message {
	// TODO: Implement GetAll method
	// Use read lock for thread safety
	// Convert map values to slice
	// Return slice of all messages
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	result := make([]*models.Message, 0, len(ms.messages))
	for _, msg := range ms.messages {
		result = append(result, msg)
	}
	return result
}

// GetByID returns a message by its ID
func (ms *MemoryStorage) GetByID(id int) (*models.Message, error) {
	// TODO: Implement GetByID method
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	msg, exists := ms.messages[id]
	if !exists {
		return nil, ErrMessageNotFound
	}
	return msg, nil
}

// Create adds a new message to storage
func (ms *MemoryStorage) Create(username, content string) (*models.Message, error) {
	// TODO: Implement Create method
	ms.mu.Lock()
	defer ms.mu.Unlock()

	msg := models.NewMessage(ms.nextID, username, content)
	ms.messages[ms.nextID] = msg
	ms.nextID++
	return msg, nil
}

// Update modifies an existing message
func (ms *MemoryStorage) Update(id int, content string) (*models.Message, error) {
	// TODO: Implement Update method
	ms.mu.Lock()
	defer ms.mu.Unlock()

	msg, exists := ms.messages[id]
	if !exists {
		return nil, ErrMessageNotFound
	}
	msg.Content = content
	return msg, nil
}

// Delete removes a message from storage
func (ms *MemoryStorage) Delete(id int) error {
	// TODO: Implement Delete method
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, exists := ms.messages[id]; !exists {
		return ErrMessageNotFound
	}
	delete(ms.messages, id)
	return nil
}

// Count returns the total number of messages
func (ms *MemoryStorage) Count() int {
	// TODO: Implement Count method
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	return len(ms.messages)
}

// Common errors
var (
	ErrMessageNotFound = errors.New("message not found")
	ErrInvalidID       = errors.New("invalid message ID")
)
