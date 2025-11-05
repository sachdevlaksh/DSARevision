package LC295

import "container/heap"


type MinHeap []int

func (m MinHeap) Len() int {
	//TODO implement me
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	//TODO implement me
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	//TODO implement me
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	//TODO implement me
	(*m) = append((*m), x.(int))
}

func (m *MinHeap) Pop() any {
	//TODO implement me
	out := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return out
}

func (m *MinHeap) Top() int {
	//TODO implement me
	return (*m)[0]
}

type MaxHeap []int

func (m MaxHeap) Len() int {
	//TODO implement me
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	//TODO implement me
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	//TODO implement me
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	//TODO implement me
	(*m) = append((*m), x.(int))
}

func (m *MaxHeap) Pop() any {
	//TODO implement me
	out := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return out
}

func (m *MaxHeap) Top() int {
	//TODO implement me
	return (*m)[0]
}

type MedianFinder struct {
    small *MaxHeap
	large *MinHeap
}

func Constructor() MedianFinder {

	mh := MedianFinder{
		small: &MaxHeap{},
		large: &MinHeap{},
	}
	heap.Init(mh.small)
	heap.Init(mh.large)

	return mh

}

func (this *MedianFinder) AddNum(num int)  {

	heap.Push(this.small, num)

	heap.Push(this.large, heap.Pop(this.small))

	if this.large.Len()  > 0 && (*this.small)[0] > (*this.large)[0]{
		heap.Push(this.large, heap.Pop(this.small))
	}

	if this.small.Len() > this.large.Len() +1{
		heap.Push(this.large, heap.Pop(this.small))
	}else if this.large.Len() > this.small.Len(){
		heap.Push(this.small, heap.Pop(this.large))
	}

}


func (this *MedianFinder) FindMedian() float64 {

	if this.small.Len() > this.large.Len(){
		return float64((*this.small).Top())
	}else{
		return float64((*this.small).Top() + (*this.large).Top())/2.0
	}
}
