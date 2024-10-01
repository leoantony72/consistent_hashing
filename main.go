package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Node struct {
	hash   int
	server string
	data   []string
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

// k,l   f,g   a,b
// 45	 55	   70
// s1 -> s2 -> s3

// k,l  f     g,a   b
// s1 -> s2 -> s4 -> s3
func (h *HashRing) AddServer(server string) {
	hash := crc32.ChecksumIEEE([]byte(server))
	node := Node{hash: int(hash), server: server, data: []string{}}
	*h = append(*h, node)
	sort.Sort(h)

	idx := sort.Search(len(*h), func(i int) bool {
		return (*h)[i].hash == int(hash)
	})
	idx=idx+1
	if idx >= len(*h) {
		idx = 0
	}
	sd := []string{}
	sd = append(sd, (*h)[idx].data...)
	(*h)[idx].data = (*h)[idx].data[:0]
	for _, v := range sd {
		h.AddData(v)
	}
}
func (h *HashRing) RemoveServer(server string) {
	hash := crc32.ChecksumIEEE([]byte(server))
	for i, node := range *h {
		if node.hash == int(hash) {
			next := i + 1
			if next >= len(*h) {
				next = 0
			}
			(*h)[next].data = append((*h)[next].data, node.data...)
			*h = append((*h)[:i], (*h)[i+1:]...)
			break
		}
	}
	sort.Sort(*h)
}

func (h *HashRing) AddData(data string) {
	hash := crc32.ChecksumIEEE([]byte(data))
	fmt.Println("DATA HASH: ", int(hash))

	idx := sort.Search(len(*h), func(i int) bool {
		return (*h)[i].hash >= int(hash)
	})

	if idx == len(*h) {
		idx = 0
	}
	(*h)[idx].data = append((*h)[idx].data, data)
}

func main() {
	ring := HashRing{}
	for {
		var num int
		fmt.Scanf("%d", &num)
		switch num {
		case 1:
			{
				fmt.Println("ADD SERVER:")
				name := ""
				fmt.Scanf("%s", &name)
				ring.AddServer(name)
			}
		case 2:
			{
				fmt.Println("ADD DATA:")
				data := ""
				fmt.Scanf("%s", &data)
				ring.AddData(data)
			}
		case 3:
			{
				fmt.Println("Print Ring:")
				for _, n := range ring {
					fmt.Println(n)
				}
			}
		case 4:
			{
				fmt.Println("REMOVE SERVER:")
				data := ""
				fmt.Scanf("%s", &data)
				ring.RemoveServer(data)
			}
		}

	}
}
