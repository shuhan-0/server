package cluster

import (
	"sync"
	"time"
)

type Manager struct {
	nodes             map[string]*Node
	nodesMutex        sync.RWMutex
	heartbeatInterval time.Duration
}

type Node struct {
	Address  string
	LastSeen time.Time
}

func NewManager(heartbeatInterval time.Duration) *Manager {
	return &Manager{
		nodes:             make(map[string]*Node),
		heartbeatInterval: heartbeatInterval,
	}
}

func (m *Manager) AddNode(address string) {
	m.nodesMutex.Lock()
	defer m.nodesMutex.Unlock()

	m.nodes[address] = &Node{
		Address:  address,
		LastSeen: time.Now(),
	}
}

func (m *Manager) RemoveNode(address string) {
	m.nodesMutex.Lock()
	defer m.nodesMutex.Unlock()

	delete(m.nodes, address)
}

func (m *Manager) UpdateNodeHeartbeat(address string) {
	m.nodesMutex.Lock()
	defer m.nodesMutex.Unlock()

	if node, exists := m.nodes[address]; exists {
		node.LastSeen = time.Now()
	}
}

func (m *Manager) GetActiveNodes() []string {
	m.nodesMutex.RLock()
	defer m.nodesMutex.RUnlock()

	var activeNodes []string
	now := time.Now()

	for address, node := range m.nodes {
		if now.Sub(node.LastSeen) <= m.heartbeatInterval*2 {
			activeNodes = append(activeNodes, address)
		}
	}

	return activeNodes
}

func (m *Manager) StartHeartbeatChecker() {
	ticker := time.NewTicker(m.heartbeatInterval)
	defer ticker.Stop()

	for range ticker.C {
		m.checkHeartbeats()
	}
}

func (m *Manager) checkHeartbeats() {
	m.nodesMutex.Lock()
	defer m.nodesMutex.Unlock()

	now := time.Now()

	for address, node := range m.nodes {
		if now.Sub(node.LastSeen) > m.heartbeatInterval*3 {
			delete(m.nodes, address)
		}
	}
}
