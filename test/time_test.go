package test

import (
	"fmt"
	"testing"
	"time"
)

func Test_time(T *testing.T) {
	t := time.Now()
	time.Sleep(time.Second)
	fmt.Println(time.Since(t))
	fmt.Println(t.Sub(time.Now().Add(time.Hour)).Microseconds())
}
