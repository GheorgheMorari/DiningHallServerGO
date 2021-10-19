package main

type Rating struct {
	values    []int
	maxSize  int
	orderNum int
	average  float32
	full      bool
}

func NewRating() *Rating {
	maxSize := 100
	return &Rating{maxSize: maxSize, values: make([]int, maxSize), full: false, average: 0}
}

func (r *Rating) addValue(rating int) {
	if !r.full && r.orderNum >= r.maxSize {
		r.full = true
	}

	r.values[r.orderNum % r.maxSize] = rating
	r.orderNum++

	numberOfReviews := 0
	if r.full {
		numberOfReviews = r.maxSize
	} else {
		numberOfReviews = r.orderNum
	}

	sum := 0
	for i := 0; i < numberOfReviews; i++ {
		sum += r.values[i]
	}

	r.average = float32(sum) / float32(numberOfReviews)
}

func (r *Rating) getAverage() float32 {
	return r.average
}

func (r *Rating) getNumOfOrders() int {
	return r.orderNum
}