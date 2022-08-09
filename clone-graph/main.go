package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	created := make([]*Node, 101)
	toVisit := make([]*Node, 1)
	toVisit[0] = node
	created[node.Val] = &Node{Val: node.Val}
	result := created[node.Val]
	for len(toVisit) > 0 {
		cur := toVisit[len(toVisit)-1]
		toVisit = toVisit[:len(toVisit)-1]

		upd := created[cur.Val]
		upd.Neighbors = make([]*Node, 0, len(cur.Neighbors))
		for _, n := range cur.Neighbors {
			if created[n.Val] == nil {
				created[n.Val] = &Node{Val: n.Val}
				toVisit = append(toVisit, n) //nolint:makezero
			}
			upd.Neighbors = append(upd.Neighbors, created[n.Val])
		}
	}
	return result
}

//nolint:deadcode,unused
func cloneGraphPointers(node *Node) *Node {
	if node == nil {
		return nil
	}

	created := make([]Node, 101)
	toVisit := make([]*Node, 1)
	toVisit[0] = node
	result := &created[node.Val]
	result.Val = node.Val
	for len(toVisit) > 0 {
		cur := toVisit[len(toVisit)-1]
		toVisit = toVisit[:len(toVisit)-1]

		upd := &created[cur.Val]
		upd.Neighbors = make([]*Node, 0, len(cur.Neighbors))
		for _, n := range cur.Neighbors {
			upd.Neighbors = append(upd.Neighbors, &created[n.Val])
			if created[n.Val].Val == 0 {
				created[n.Val].Val = n.Val
				toVisit = append(toVisit, n) //nolint:makezero
			}
		}
	}
	return result
}

func main() {
	nodes := make([]Node, 5)

	nodes[1].Val = 1
	nodes[1].Neighbors = append(nodes[1].Neighbors, &nodes[2])

	nodes[2].Val = 2
	nodes[2].Neighbors = append(nodes[2].Neighbors, &nodes[3])

	nodes[3].Val = 3
	nodes[3].Neighbors = append(nodes[3].Neighbors, &nodes[4])

	nodes[4].Val = 4
	nodes[4].Neighbors = append(nodes[4].Neighbors, &nodes[1])

	cloneGraph(&nodes[1])
}
