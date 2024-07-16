package cluster

import (
	"math/rand"
	"sync"
)

type LoadBalancer struct {
	nodes []string
	mu    sync.RWMutex
}

func NewLoadBalancer(nodes []string) *LoadBalancer {
	return &LoadBalancer{
		nodes: nodes,
	}
}

func (lb *LoadBalancer) GetNode() string {
	lb.mu.RLock()
	defer lb.mu.RUnlock()
	return lb.nodes[rand.Intn(len(lb.nodes))]
}

func (lb *LoadBalancer) AddNode(node string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	lb.nodes = append(lb.nodes, node)
}

func (lb *LoadBalancer) RemoveNode(node string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	for i, n := range lb.nodes {
		if n == node {
			lb.nodes = append(lb.nodes[:i], lb.nodes[i+1:]...)
			break
		}
	}
}
