package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 令牌桶限流
type TokenBucket struct {
	capacity  chan int      // 令牌桶容量
	rate      time.Duration // 令牌生成速率
	tokens    int           // 当前令牌数量
	timestamp time.Time     // 上次生成令牌的时间
	//mutex     sync.Mutex    // 互斥锁，保证并发安全
}

func NewTokenBucket(tokens int, rate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:  make(chan int, tokens),
		rate:      rate,
		tokens:    tokens,
		timestamp: time.Now(),
	}

}

/*
用channel 实现令牌桶
*/
var wg sync.WaitGroup

func (tb *TokenBucket) TakeToken(i int) bool {
	time.Sleep(time.Second)
	defer wg.Done()
	//defer wg.Done()
	//    记录时间
	now := time.Now()
	//经过的时间
	t := now.Sub(tb.timestamp)
	fmt.Println(tb.rate, t)
	//	 生成的令牌
	rate := int((tb.rate * t).Seconds())
	fmt.Println("rate", rate, "t", t, "tb", tb.tokens, "tb.capacity", len(tb.capacity))

	if rate > 0 {
		var mest int
		if rate+tb.tokens > cap(tb.capacity) && tb.tokens != cap(tb.capacity) {
			mest = rate + tb.tokens - len(tb.capacity)
			fmt.Println("mest", mest)
			for i := 0; i < mest; i++ {
				tb.capacity <- i
			}
		}
		if rate+tb.tokens < cap(tb.capacity) {
			//桶里面放入令牌
			for i := 0; i < rate; i++ {
				tb.capacity <- 1
				tb.timestamp = now
			}
		}

	}
	//  取令牌
	if len(tb.capacity) > 0 {
		<-tb.capacity
		fmt.Println("第", i, true)
		return true
	}
	fmt.Println("第", i, false)
	return false

}

func Test_limit(t *testing.T) {
	//   NewTokenBucket
	tb := NewTokenBucket(10, 1)
	go func() {
		for i := 0; i < tb.tokens; i++ {
			tb.capacity <- i
		}
	}()

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go tb.TakeToken(i)
	}

	//
	//if tb.TakeToken(){
	//	fmt.Println("请求成功")
	//}
	//fmt.Println("请求失败")
	defer wg.Wait()

}
