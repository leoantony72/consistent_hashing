package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Node struct {
	hash   int
	server string
}

type HashRing []Node

func (h HashRing) Len() int {
	return len(h)
}
func (h HashRing) Less(i, j int) bool {
	return h[i].hash < h[j].hash
}

func (h HashRing) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HashRing) AddServer(server string) {
	hash := crc32.ChecksumIEEE([]byte(server))
	node := Node{hash: int(hash), server: server}
	*h = append(*h, node)
	sort.Sort(h)
}
func (h *HashRing) RemoveServer(server string) {
	hash := crc32.ChecksumIEEE([]byte(server))
	for i, node := range *h {
		if node.hash == int(hash) {
			*h = append((*h)[:i], (*h)[i+1:]...)
			break
		}

	}
	sort.Sort(*h)
}

func main() {

	ring := HashRing{}
	ring.AddServer("s1")
	ring.AddServer("s2")
	ring.AddServer("s34")
	ring.AddServer("s78")
	fmt.Println(ring)
}
