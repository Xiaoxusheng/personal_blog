package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_blog/models"
	"strconv"
	"sync"
	"time"
)

// 记录ip次数
func LimitIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.RemoteIP()
		UseAgent := c.GetHeader("User-Agent")
		//查询是否为封禁ip
		b := models.GetIp(ip)
		if b {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "IP访问次数过于频繁，已经封禁",
			})
			return
		}
		//插入
		if f := models.InsertIP(&models.Ip{Ip: ip, Time: strconv.FormatInt(time.Now().Unix(), 10), UseAgent: UseAgent}); f != nil {
			c.Abort()
			panic(f.Error())
		}

		//查询次数
		num := models.GetIpNumber()
		if num > 60 {
			//加黑名单
			models.InsertIpbyBans(ip, strconv.FormatInt(time.Now().Unix(), 10))
		}
		c.Next()
	}

}

// 令牌桶限流
type TokenBucket struct {
	capacity  int           // 令牌桶容量
	rate      time.Duration // 令牌生成速率
	tokens    int           // 当前令牌数量
	timestamp time.Time     // 上次生成令牌的时间
	mutex     sync.Mutex    // 互斥锁，保证并发安全
}

func NewTokenBucket(capacity int, rate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:  capacity,
		rate:      rate,
		tokens:    capacity,
		timestamp: time.Now(),
	}
}

func (tb *TokenBucket) TakeToken() bool {
	//上锁
	tb.mutex.Lock()
	//解锁
	defer tb.mutex.Unlock()
	now := time.Now()
	// 计算时间间隔，生成令牌数量
	diff := now.Sub(tb.timestamp)

	//令牌数目
	count := int(diff / tb.rate)

	if count > 0 {
		tb.tokens += count
		//如果当前令牌数量>当前令牌桶容量,赋值令牌桶最大容量
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
		//更新生成令牌的时间
		tb.timestamp = now
	}
	// 判断令牌数量是否足够
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

// LimitIP TODO IP限流中间件
func RequetLimite() gin.HandlerFunc {
	return func(c *gin.Context) {
		tb := NewTokenBucket(10, time.Second)
		if !tb.TakeToken() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
