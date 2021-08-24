package sort

type MedianFinder struct {
	arr []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {

	return MedianFinder{arr: make([]int, 0)}
}

func (this *MedianFinder) AddNum(num int) {

	if len(this.arr) == 0 {
		this.arr = append(this.arr, num)
		return
	} else {
		for i, v := range this.arr {
			if v > num {
				this.arr = append(this.arr[0:i], append([]int{num}, this.arr[i:]...)...)
				return
			}
		}
	}
	this.arr = append(this.arr, num)

}

func (this *MedianFinder) FindMedian() float64 {

	l := len(this.arr)
	if l%2 == 0 {
		return float64(this.arr[l/2]+this.arr[l/2-1]) / 2.0
	} else {
		return float64(this.arr[l/2])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
