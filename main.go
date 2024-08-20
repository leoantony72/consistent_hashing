package main

type Node struct {
	Name string
	Host string
}

var storageNode []Node = []Node{
	Node{"A", "192.8.12.1"},
	Node{"B", "192.8.12.2"},
	Node{"C", "192.8.12.3"},
	Node{"D", "192.8.12.4"},
	Node{"E", "192.8.12.5"},
}

func main() {

}

func hash(path string) int {
	sum := 0
	for _, b := range []byte(path) {
		sum += int(b)
	}
	return sum % 5
}

func upload(path string) {

	index := hash(path)

	node := storageNode[index]

	return node.put_file(path)
}

