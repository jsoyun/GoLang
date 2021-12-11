package dataStruct

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	h.list = append(h.list, v)
	//현재인덱스= 리스트길이-1 (길이는 인덱스보다 하나 큰거니까)
	idx := len(h.list) - 1
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] > h.list[parentIdx] {
			//현재노드랑 부모랑 자리 바꿔줌
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Print() {
	fmt.Println(h.list)
}

func (h *Heap) Pop() int {
	if len(h.list)==0 {
		return 0
	}
	top:=h.list[0]
	last := h.list[len(h.list)-1]
	//끝 하나 잘라냄
	h.list = h.list[:len(h.list)-1]
	//맨마지막에 있는 애를 맨 처음으로 올려
	h.list[0]=last
	//현재인덱스는 맨처음
	idx := 0
	//인덱스가 전체 길이보다 작은 경우
	for idx < len(h.list){
		swapIdx := -1
		leftIdx := idx*2+1
		if leftIdx >= len(h.list){
			break
		}
		//왼쪽이랑    현재값이랑 비교
		if h.list[leftIdx] > h.list[idx]{
			//오른쪽도 비교해봐야돼서 바로 바꾸지 않고 스왑인덱스를 왼쪽으로 표시
			swapIdx=leftIdx
		}
		// left := h.list[	leftIdx]
		rightIdx := idx*2+2
		if rightIdx < len(h.list){
			//오른쪽과 현재 비교했을때
		  if h.list[rightIdx]> h.list[idx]{
			  if swapIdx<0|| h.list[swapIdx] < h.list[rightIdx]{
				  swapIdx= rightIdx
			  }
		  }
		}
		//0보다 작다는 뜻은
	if swapIdx <0 {
			//바꿀애가 없음
			break
		}
		//스왑
		h.list[idx],h.list[swapIdx] = h.list[swapIdx],h.list[idx]
		idx= swapIdx
		
	}
	return top


}