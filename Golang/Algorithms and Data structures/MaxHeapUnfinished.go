//Unfinished

//A heap is a data sructure that can be expressed as a complete tree
//heapify time complexity o(logn)

package main

//struct for maxheap has slice for array
type MaxHeap struct{
	array []int
	
}
//insert method add elements to the heap
func (h *MaxHeap) Insert(key int){
	h.array = append(h.array,key)
	//rearrange to maintain heap properties
	h.maxHeapifyUp(len(h.array)-1)
	
}
//extract method return largest key & removes it from heap
func (h *MaxHeap)maxHeapifyUp(index int){
	
	
}
func parent(i int)int{
	return (i-1)/2
}

//get the left child index
func left(i int)int{
	return 2*i + 1
}

//get the right child index
func right(i int)int{
	return 2*i + 1
}

//swap keys in the array
func (h *MaxHeap) swap(i1,i2){
	h.array[i1], h.array[i2] = h.array[i2], h.array[i2]
}



func main(){
	
	
	
	
	
	
}

