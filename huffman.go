package determine_frames_types

import (
	"container/heap"
)

type HuffmanTree struct {
	frequency int
	value     rune
	code      string
	children  []*HuffmanTree
}

type PriorityQueue []*HuffmanTree

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].frequency < pq[j].frequency }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*HuffmanTree)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func BuildHuffmanTree(freqMap map[rune]int) *HuffmanTree {
	pq := make(PriorityQueue, len(freqMap))
	i := 0
	for value, freq := range freqMap {
		pq[i] = &HuffmanTree{frequency: freq, value: value}
		i++
	}
	heap.Init(&pq)

	for pq.Len() > 1 {
		lo := heap.Pop(&pq).(*HuffmanTree)
		hi := heap.Pop(&pq).(*HuffmanTree)

		merged := &HuffmanTree{
			frequency: lo.frequency + hi.frequency,
			children:  []*HuffmanTree{lo, hi},
		}
		heap.Push(&pq, merged)
	}

	return heap.Pop(&pq).(*HuffmanTree)
}

func AssignCodes(node *HuffmanTree, prefix string) map[rune]string {
	if node == nil {
		return nil
	}
	codes := make(map[rune]string)
	if node.value > 0 { // Leaf node
		codes[node.value] = prefix
	} else {
		leftCodes := AssignCodes(node.children[0], prefix+"0")
		rightCodes := AssignCodes(node.children[1], prefix+"1")
		for k, v := range leftCodes {
			codes[k] = v
		}
		for k, v := range rightCodes {
			codes[k] = v
		}
	}
	return codes
}

func HuffmanCoding(data []rune) map[rune]string {
	freqMap := make(map[rune]int)
	for _, v := range data {
		freqMap[v]++
	}
	tree := BuildHuffmanTree(freqMap)
	return AssignCodes(tree, "")
}
