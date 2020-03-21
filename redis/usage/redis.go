/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package usage

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

const (
	Addr     = "localhost:6379"
	Password = ""
	DB       = 0
)

type Redis struct {
	Client *redis.Client
}

// 构造
func NewRedis() *Redis{
	client := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	} else {
		return &Redis{Client: client}
	}
}

/**
 ** string
 */

// 设置key
func (r *Redis) SetKey(key string, value string, exp time.Duration) {
	err := r.Client.Set(key, value, exp).Err()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("set %v value %v\n", key, value)
	}
}

// 设置key 如果key存在 则不添加
func (r *Redis) SetNXKey(key string, value string, exp time.Duration) {
	result, err := r.Client.SetNX(key, value, exp).Result()
	if err != nil {
		fmt.Println(err.Error())
	} else if result {
		fmt.Printf("setnx %v is %t\n", key, result)
	} else {
		fmt.Printf("setnx %v is %t\n", key, result)
	}
}

// 获取key
func (r *Redis) GetKey(key string) {
	value, err := r.Client.Get(key).Result()
	if err == redis.Nil {
		fmt.Printf("%s does not exist\n", key)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("get %s value %s\n", key, value)
	}
}

// 判断key是否存在
func (r *Redis) Exist(key string) {
	result, _ := r.Client.Exists(key).Result()
	fmt.Printf("%s exist result is %v\n", key, result)
}

// 判断key生存时间
func (r *Redis) Ttl(key string) {
	result, err := r.Client.TTL(key).Result()
	fmt.Println("Ttl", result, err)
}

// 删除key
func (r *Redis) Del(keys ...string) {
	result, err := r.Client.Del(keys...).Result()
	fmt.Println("Del", result, err)
}

// 自增1
func (r *Redis) Incr(key string) {
	result, err := r.Client.Incr(key).Result()
	fmt.Println("Incr", result, err)
}

// 自定义自增
func (r *Redis) IncrBy(key string, num int64) {
	result, err := r.Client.IncrBy(key, num).Result()
	fmt.Println("IncrBy", result, err)
}

// 自减1
func (r *Redis) Decr(key string) {
	result, err := r.Client.Decr(key).Result()
	fmt.Println("Decr", result, err)
}

// 自定义自减
func (r *Redis) DecrBy(key string, num int64) {
	result, err := r.Client.DecrBy(key, num).Result()
	fmt.Println("DecrBy", result, err)
}

/**
 ** list
 */

// LPush RPush同
func (r *Redis) LPush(key string, values ...interface{}) {
	result, err := r.Client.LPush(key, values...).Result()
	fmt.Println("LPush", result, err)
}

// LPushX 如果value值存在 则不添加 RPushX同
func (r *Redis) LPushX(key string, values ...interface{}) {
	result, err := r.Client.LPushX(key, values...).Result()
	fmt.Println("LPushX", result, err)
}

// LPop RPop同
func (r *Redis) LPop(key string) {
	result, err := r.Client.LPop(key).Result()
	fmt.Println("LPop", result, err)
}

// 修改下标为index的list
func (r *Redis) LSet(key string, index int64, value interface{}) {
	result, err := r.Client.LSet(key, index, value).Result()
	fmt.Println("LSet", result, err)
}

// 获取list
func (r *Redis) LRange(key string, start, stop int64) {
	result, err := r.Client.LRange(key, start, stop).Result()
	fmt.Println("LRange", result, err)
}

// 获取list长度
func (r *Redis) LLen(key string) {
	result, err := r.Client.LLen(key).Result()
	fmt.Println("LLen", result, err)
}

// sort list LIMIT 0 2 ASC
func (r *Redis) Sort(key string, sort *redis.Sort) {
	result, err := r.Client.Sort(key, sort).Result()
	fmt.Println("Sort", result, err)
}

/**
 ** set
 */

// 添加元素
func (r *Redis) SAdd(key string, members ...interface{}) {
	result, err := r.Client.SAdd(key, members...).Result()
	fmt.Println("SAdd", result, err)
}

// 获取set
func (r *Redis) SMembers(key string) {
	result, err := r.Client.SMembers(key).Result()
	fmt.Println("SMembers", result, err)
}

// 获取set元素数量
func (r *Redis) SCard(key string) {
	result, err := r.Client.SCard(key).Result()
	fmt.Println("SCard", result, err)
}

// 差集 返回第一个set里面有的，其它set里面没有的元素
func (r *Redis) SDiff(keys ...string) {
	result, err := r.Client.SDiff(keys...).Result()
	fmt.Println("SDiff", result, err)
}

// 交集 返回所有set交集
func (r *Redis) SInter(keys ...string) {
	result, err := r.Client.SInter(keys...).Result()
	fmt.Println("SInter", result, err)
}

// 随机取出集合中某个元素并且删除 适用于抽奖
func (r *Redis) SPop(key string) {
	result, err := r.Client.SPop(key).Result()
	fmt.Println("SPop", result, err)
}

// 随机取出集合中某个元素不删除 适用于重复抽奖
func (r *Redis) SRandMember(key string) {
	result, err := r.Client.SRandMember(key).Result()
	fmt.Println("SRandMember", result, err)
}

// 判断元素是否在集合中
func (r *Redis) SIsMember(key string, member interface{}) {
	result, err := r.Client.SIsMember(key, member).Result()
	fmt.Println("SIsMember", result, err)
}

/**
 ** zset
 */

// 添加元素
func (r *Redis) ZAdd(key string,  members ...*redis.Z) {
	result, err := r.Client.ZAdd(key, members...).Result()
	fmt.Println("ZAdd", result, err)
}

// 获取zset 按分数正序排序
func (r *Redis) ZRange(key string, start, stop int64) {
	result, err := r.Client.ZRange(key, start, stop).Result()
	fmt.Println("ZRange", result, err)
}

// 获取zset 按分数倒序排序
func (r *Redis) ZRevRange(key string, start, stop int64) {
	result, err := r.Client.ZRevRange(key, start, stop).Result()
	fmt.Println("ZRevRange", result, err)
}

// 获取zset元素数量
func (r *Redis) ZCard(key string) {
	result, err := r.Client.ZCard(key).Result()
	fmt.Println("ZCard", result, err)
}

/**
 ** Hash
 */

// 添加值
func (r *Redis) HSet(key string, values ...interface{}) {
	result, err := r.Client.HSet(key, values...).Result()
	fmt.Println("HSet", result, err)
}

// 根据field获取值
func (r *Redis) HGet(key, field string) {
	result, err := r.Client.HGet(key, field).Result()
	fmt.Println("HGet", result, err)
}

// 获取所有field和value
func (r *Redis) HGetAll(key string) {
	result, err := r.Client.HGetAll(key).Result()
	fmt.Println("HGetAll", result, err)
}

// 返回元素个数
func (r *Redis) HLen(key string) {
	result, err := r.Client.HLen(key).Result()
	fmt.Println("HLen", result, err)
}

// 返回所有field
func (r *Redis) HKeys(key string) {
	result, err := r.Client.HKeys(key).Result()
	fmt.Println("HKeys", result, err)
}

// 返回所有values
func (r *Redis) HVals(key string) {
	result, err := r.Client.HVals(key).Result()
	fmt.Println("HVals", result, err)
}

// 删除field
func (r *Redis) HDel(key, field string) {
	result, err := r.Client.HDel(key, field).Result()
	fmt.Println("HDel", result, err)
}

// field是否存在
func (r *Redis) HExists(key, field string) {
	result, err := r.Client.HExists(key, field).Result()
	fmt.Println("HExists", result, err)
}

// 整数自增
func (r *Redis) HIncrBy(key, field string, incr int64) {
	result, err := r.Client.HIncrBy(key, field, incr).Result()
	fmt.Println("HIncrBy", result, err)
}
