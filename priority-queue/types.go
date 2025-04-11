package priority_queue

import (
	"dijkstras-alg-go/model"
)

type PriorityQueue []*Item

type Item struct {
	Value    model.NodeID
	Priority int
	Index    int
}
