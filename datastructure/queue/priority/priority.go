package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item // PriorityQueue 인터페이스를 구현
// golang 은 java와 달리 implements 키워드를 명시하지 않으며, 이 타입이 인터페이스의 메서드를 모두 구현하면 알아서 판단한다.

func (pq PriorityQueue) Len() int { // 길이를 리턴
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool { // 오름차순 or 내림차순 을 결정하는 함수
	return pq[i].priority > pq[j].priority // 내림차순 정렬 (Max heap)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]

	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]

	return item
}

func main() {
	pq := new(PriorityQueue) // 힙 생성

	heap.Init(pq)

	heap.Push(pq, 5)
	heap.Push(pq, 2)
	heap.Push(pq, 7)
	heap.Push(pq, 3)

	fmt.Println(pq, "최댓값 : ", (*pq)[0])
}
