package kit

// Stack nums
type Stack struct {
	nums []int
}

// Push element into stack
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

// Pop out last element
func (s *Stack) Pop() int {
	r := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return r
}

// Len return stack length
func (s *Stack) Len() int {
	return len(s.nums)
}

// Empty return bool
func (s *Stack) Empty() bool {
	return len(s.nums) == 0
}

// NewStack return new stack
func NewStack() *Stack {
	return &Stack{
		nums: []int{}}
}
