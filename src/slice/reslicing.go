package slice

import (
	"log"
)

func Reslicing() {
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	log.Println("s1", s1)
	log.Println("reSlice", reSlice)

	reSlice[0] = -100
	reSlice[1] = 123456

	log.Println("s1", s1)
	log.Println("reSlice", reSlice)

}
