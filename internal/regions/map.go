package regions

import (
	"fmt"
	"sync"

	"github.com/dpsommer/eventstream/internal/utils"
)

type Node struct {
	Value Location
	// declare priority to implement Prioritizable for the pqueue
	priority int64
}

func (n Node) Priority() int64 { return n.priority }

type Map struct {
	nodes map[Location]*Node
	// XXX: simplify edges; only have distance prop for now
	edges map[Location]map[Location]int

	// common concurrent-access heap
	safeHeap utils.SafeHeap
	sync.Mutex
}

func NewMap() *Map {
	return &Map{
		nodes:    map[Location]*Node{},
		edges:    map[Location]map[Location]int{},
		safeHeap: utils.SafeHeap{},
	}
}

func (m *Map) AddNode(loc Location) {
	m.Lock()
	defer m.Unlock()

	m.nodes[loc] = &Node{
		Value: loc,
	}

	m.edges[loc] = map[Location]int{}
}

func (m *Map) AddEdge(from, to Location, distance int) error {
	m.Lock()
	defer m.Unlock()

	if err := m.hasNode(from, to); err != nil {
		return fmt.Errorf("failed to add edge: %w", err)
	}

	// XXX: assume all edges are bidirectional
	m.edges[from][to] = distance
	m.edges[to][from] = distance

	return nil
}

func (m *Map) hasNode(nodes ...Location) error {
	for _, loc := range nodes {
		if _, ok := m.nodes[loc]; !ok {
			return fmt.Errorf("no node %s in graph", loc)
		}
	}

	return nil
}

func (m *Map) Distance(from, to Location) ([]*Node, int, error) {
	m.Lock()
	defer m.Unlock()

	if err := m.hasNode(from, to); err != nil {
		return nil, -1, fmt.Errorf("failed to find path from %s to %s: %w", from, to, err)
	}

	if d, ok := m.edges[from][to]; ok {
		return []*Node{m.nodes[to], m.nodes[from]}, d, nil
	}

	if from == to {
		return []*Node{m.nodes[from]}, 0, nil
	}

	path, distance, found := m.aStar(from, to)
	if !found {
		return nil, -1, fmt.Errorf("no path found from %s to %s", from, to)
	}
	return path, distance, nil
}

// implement a simple A* search for pathfinding between locations
func (m *Map) aStar(from, to Location) ([]*Node, int, bool) {
	paths := map[*Node]*Node{}
	costs := map[*Node]int{}

	frontier := &utils.PriorityQueue[Node]{}
	m.safeHeap.Init(frontier)

	fromNode := m.nodes[from]
	fromNode.priority = 0
	m.safeHeap.Push(frontier, fromNode)

	for {
		if frontier.Len() == 0 {
			return nil, 0, false
		}

		current := m.safeHeap.Pop(frontier).(*Node)

		if current.Value == to {
			// return the first path found and its cost
			path := []*Node{}
			next := current
			for next != nil {
				path = append(path, next)
				next = paths[next]
			}
			return path, costs[current], true
		}

		for next, dist := range m.edges[current.Value] {
			nextNode := m.nodes[next]
			nextCost := costs[current] + dist
			if _, ok := costs[nextNode]; !ok || nextCost < costs[nextNode] {
				costs[nextNode] = nextCost
				// TODO: should the map be grid-based so we can use a simple
				// straight-line heuristic here?
				nextNode.priority = int64(nextCost) // + heuristic(next, to)
				m.safeHeap.Push(frontier, nextNode)
				paths[nextNode] = current
			}
		}
	}
}
