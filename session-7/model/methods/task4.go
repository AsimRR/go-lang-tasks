package methods

type Student struct {
	Name   string
	Grades []int
}

func (s *Student) Average() float32 {
	length := len(s.Grades)
	if length == 0 {
		return 0
	}

	sum := 0
	for _, num := range s.Grades {
		sum += num
	}

	return float32(sum) / float32(length)
}
