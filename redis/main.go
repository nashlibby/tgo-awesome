package main

import (
	"github.com/go-redis/redis/v7"
	"redis/usage"
)

func main() {
	r := usage.NewRedis()
	// // 设置key
	// r.SetKey("key1", "value1", 10*time.Minute)
	// r.SetNXKey("key1", "value1", 10*time.Minute)
	// r.SetNXKey("key3", "value1", 10*time.Minute)
	// // 获取key
	// r.GetKey("key1")
	// r.GetKey("key2")
	// // 判断key是否存在
	// r.Exist("key1")
	// r.Exist("key2")
	// r.Ttl("key1")
	// // 删除key
	// r.Del("key1", "key2")
	// // 自增1
	// r.Incr("key4")
	// // 自定义自增
	// r.IncrBy("key5", 5)
	// // 自减1
	// r.Incr("key6")
	// // 自定义自减
	// r.IncrBy("key7", 5)
	//
	// // LPush
	// r.LPush("list1", "a", "b", "c")
	// // LPushX
	// r.LPushX("list1", "a", "d", "e")
	// // LPop
	// r.LPop("list1")
	// // LSet
	// r.LSet("list1", 1, "ab")
	// LRange
	// r.LRange("list1", 0, -1)
	// // LLen
	// r.LLen("list1")
	// // sort
	// r.Sort("list1", &redis.Sort{Offset: 0, Count: 10, Order: "ASC", Alpha: true})

	// SAdd
	r.SAdd("s1", "a", "b", "c")
	r.SAdd("s2", "c", "d", "e")
	r.SAdd("s3", "a", "c")
	// SMembers
	r.SMembers("s1")
	// SCard
	r.SCard("s1")
	// SDiff
	r.SDiff("s1", "s2", "s3")
	// SInter
	r.SInter("s1", "s2", "s3")
	// SPop
	r.SPop("s1")
	// SRandMember
	r.SRandMember("s1")
	// SIsMember
	r.SIsMember("s1", "a")


}
