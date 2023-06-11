package test

import (
	"fmt"
	"testing"
	"time"
)

func Test_time(T *testing.T) {
	t := time.Now()
	time.Sleep(time.Second)

	t3 := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.Now().Location()).Unix()
	fmt.Println(time.Since(t))
	fmt.Println(t.Sub(time.Now().Add(time.Hour)).Microseconds(), t3)

}
