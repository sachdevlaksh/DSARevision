package LC632

import (
	"container/heap"
	"math"
)

//type Element struct {
//    val   int
//    row   int
//    index int
//}
//
//type MinHeap []Element
//
//func (h MinHeap) Len() int            { return len(h) }
//func (h MinHeap) Less(i, j int) bool  { return h[i].val < h[j].val }
//func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
//func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Element)) }
//func (h *MinHeap) Pop() interface{} {
//    old := *h
//    n := len(old)
//    x := old[n-1]
//    *h = old[:n-1]
//    return x
//}
//
//func smallestRange(nums [][]int) []int {
//    h := &MinHeap{}
//    heap.Init(h)
//    currentMax := math.MinInt32
//
//    // Initialize heap with the first element from each list
//    for i := 0; i < len(nums); i++ {
//        heap.Push(h, Element{nums[i][0], i, 0})
//        if nums[i][0] > currentMax {
//            currentMax = nums[i][0]
//        }
//    }
//
//    rangeStart, rangeEnd := math.MinInt32, math.MaxInt32
//
//    for {
//        minElem := heap.Pop(h).(Element)
//
//        // Update range if smaller
//        if currentMax-minElem.val < rangeEnd-rangeStart {
//            rangeStart, rangeEnd = minElem.val, currentMax
//        }
//
//        // Move next in same list
//        if minElem.index+1 == len(nums[minElem.row]) {
//            break // list exhausted
//        }
//
//        nextVal := nums[minElem.row][minElem.index+1]
//        heap.Push(h, Element{nextVal, minElem.row, minElem.index + 1})
//        if nextVal > currentMax {
//            currentMax = nextVal
//        }
//    }
//
//    return []int{rangeStart, rangeEnd}
//}
//
//func main() {
//    nums := [][]int{
//        {4, 10, 15, 24, 26},
//        {0, 9, 12, 20},
//        {5, 18, 22, 30},
//    }
//    fmt.Println(smallestRange(nums)) // Output: [20,24]
//}

type Ele struct {
	Value int
	Row   int
	Index int
}

type MinHeap []Ele

func (m MinHeap) Len() int {
	//TODO implement me
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	//TODO implement me
	return m[i].Value < m[j].Value
}

func (m MinHeap) Swap(i, j int) {
	//TODO implement me
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	//TODO implement me
	(*m) = append((*m), x.(Ele))
}

func (m *MinHeap) Pop() any {
	//TODO implement me
	out := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return out
}

func SmallestRange(nums [][]int) []int {
	pq := &MinHeap{}
	heap.Init(pq)
	n := len(nums)
	currentMax := math.MinInt
	for i := 0; i < n; i++ {
		heap.Push(pq, Ele{Value: nums[i][0], Index: 0, Row: i})
		if currentMax < nums[i][0] {
			currentMax = nums[i][0]
		}
	}

	minRange, maxRange := 0, math.MaxInt

	for {
		currentEleMin := heap.Pop(pq).(Ele)

		if currentMax-currentEleMin.Value < maxRange-minRange {
			minRange, maxRange = currentEleMin.Value, currentMax
		}

		if currentEleMin.Index+1 == len(nums[currentEleMin.Row]) {
			break
		}

		nextEleVal := nums[currentEleMin.Row][currentEleMin.Index+1]
		heap.Push(pq, Ele{Value: nextEleVal, Row: currentEleMin.Row, Index: currentEleMin.Index + 1})

		if nextEleVal > currentMax {
			currentMax = nextEleVal
		}
	}

	return []int{minRange, maxRange}
}
