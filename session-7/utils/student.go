package utils

import (
	"fmt"
	"session-7/model/methods"
)

func CompareStudents(s1, s2 methods.Student) string {
	avg1 := s1.Average()
	avg2 := s2.Average()

	if avg1 > avg2 {
		return fmt.Sprintf("%s has a higher average grade.", s1.Name)
	} else if avg1 < avg2 {
		return fmt.Sprintf("%s has a higher average grade.", s2.Name)
	} else {
		return fmt.Sprintf("%s and %s have the same average grade.", s1.Name, s2.Name)
	}
}
