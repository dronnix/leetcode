package main

// The task: https://leetcode.com/problems/implement-queue-using-stacks/

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{in: make([]int, 0, 25), out: make([]int, 0, 25)}
}

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) Pop() int {
	r := q.Peek()
	q.out = q.out[:len(q.out)-1]
	return r
}

func (q *MyQueue) Peek() int {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			q.out = append(q.out, q.in[len(q.in)-1])
			q.in = q.in[:len(q.in)-1]
		}
	}
	return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.in)+len(q.out) == 0
}

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	for !q.Empty() {
		println(q.Pop())
	}
}
