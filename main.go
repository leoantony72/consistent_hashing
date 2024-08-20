package main

type Node struct {
	Name string
	Host string
}

func (n *Node) put_file(path string)string{
	return path
}
func (n *Node) fetch(path string)string{
	return path
}

var storageNode []Node = []Node{
	{"A", "192.8.12.1"},
	{"B", "192.8.12.2"},
	{"C", "192.8.12.3"},
	{"D", "192.8.12.4"},
	{"E", "192.8.12.5"},
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

func upload(path string) string{

	index := hash(path)

	node := storageNode[index]

	return node.put_file(path)
}

func fetch(path string) string{
	index := hash(path)

	node := storageNode[index]

	return node.fetch(path)
}
