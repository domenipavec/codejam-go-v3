package main

type Queue struct {
	data  []int
	front int
	back  int
}

func (q *Queue) inc(i int) int {
	return (i + 1) & (len(q.data) - 1)
}

func (q *Queue) dec(i int) int {
	return (i - 1) & (len(q.data) - 1)
}

func (q Queue) Len() int {
	return (q.back - q.front) & (len(q.data) - 1)
}

func (q *Queue) growIfNeeded() {
	if len(q.data) == 0 {
		q.data = make([]int, 64)
	}
	l := q.Len()
	if l < len(q.data)-1 {
		return
	}
	newData := make([]int, 4*len(q.data))
	if q.front < q.back {
		copy(newData, q.data[q.front:q.back])
	} else {
		n := copy(newData, q.data[q.front:])
		copy(newData[n:], q.data[:q.back])
	}
	q.data = newData
	q.front = 0
	q.back = l
}

func (q *Queue) Push(v int) {
	q.growIfNeeded()
	q.data[q.back] = v
	q.back = q.inc(q.back)
}

func (q *Queue) PushFront(v int) {
	q.growIfNeeded()
	q.front = q.dec(q.front)
	q.data[q.front] = v
}

func (q Queue) Front() int {
	return q.data[q.front]
}

func (q Queue) Back() int {
	return q.data[q.dec(q.back)]
}

func (q *Queue) Pop() int {
	v := q.Front()
	q.front = q.inc(q.front)
	return v
}

func (q *Queue) PopBack() int {
	q.back = q.dec(q.back)
	return q.data[q.back]
}
