package kit

// Queue store nums int
type Queue struct {
	nums []int
}

// Push n to Queue
func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

// Pop first num
func (q *Queue) Pop() int {
	r := q.nums[0]
	q.nums = q.nums[1:]
	return r
}

// Len return number of queue lenght
func (q *Queue) Len() int {
	return len(q.nums)
}

// Empty return bool
func (q *Queue) Empty() bool {
	return len(q.nums) == 0
}

// NewQueue return *Queue
func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}
