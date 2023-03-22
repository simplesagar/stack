package lock

import (
	"sync"
)

// TODO(gfyrag): In a future release, we should be able to clear old lockers from memory
type Manager struct {
	mu      sync.Mutex
	ledgers map[string]Locker
}

func (m *Manager) ForLedger(ledger string) Locker {
	m.mu.Lock()
	defer m.mu.Unlock()

	locker, ok := m.ledgers[ledger]
	if !ok {
		locker = NewDefaultLocker(ledger)
		m.ledgers[ledger] = locker
	}
	return locker
}

func NewManager() *Manager {
	return &Manager{
		ledgers: map[string]Locker{},
	}
}
