package main

import (
    "container/heap"
    "strings"
)

func reorganizeString(s string) string {
    freqMap := make(map[rune]int)
    for _, r := range s {
        freqMap[r]++
    }
    maxHeap := &MaxHeap{}
    heap.Init(maxHeap)

    for c, freq := range freqMap {
        heap.Push(maxHeap, CharFreq{c, freq})
    }

    var res strings.Builder

    for maxHeap.Len() >= 2 {
        cf1 := heap.Pop(maxHeap).(CharFreq)
        cf2 := heap.Pop(maxHeap).(CharFreq)

        res.WriteRune(cf1.char)
        res.WriteRune(cf2.char)
        if cf1.count > 1 {
            heap.Push(maxHeap, CharFreq{cf1.char, cf1.count - 1})
        }
        if cf2.count > 1 {
            heap.Push(maxHeap, CharFreq{cf2.char, cf2.count - 1})
        }
    }

    if maxHeap.Len() > 0 {
        lastFreq := heap.Pop(maxHeap).(CharFreq)

        if lastFreq.count > 1 {
            return ""
        }

        res.WriteRune(lastFreq.char)
    }
    return res.String()
}

type CharFreq struct {
    char  rune
    count int
}
type MaxHeap []CharFreq

func (h MaxHeap) Less(i, j int) bool {
    return h[i].count > h[j].count
}

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(CharFreq))
}

func (h MaxHeap) Len() int {
    return len(h)
}
func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]

}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    out := old[n-1]
    *h = old[:n-1]
    return out
}
