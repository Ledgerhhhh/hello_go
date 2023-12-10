package test

import (
	"com.ledger.goproject/myconfig"
	"com.ledger.goproject/try_redis/conn"
	"fmt"
	"testing"
	"time"
)

func init() {
	err := myconfig.InitGConfig()
	if err != nil {
		_ = fmt.Errorf("ttl_service error: %s\n", err)
	}
	err = conn.InitRedisConfig()
	if err != nil {
		_ = fmt.Errorf("redis error:%s", err)
	}
}

func TestTryString(t *testing.T) {
	// 设置一个字符串
	set := conn.RClient.Set("ledger", "hello redis", time.Second*10)
	t.Logf("set: %s", set.Val()) //"OK"
	// 获取一个字符串
	get := conn.RClient.Get("ledger")
	t.Logf("get: %s", get.Val()) //hello redis
	// 删除一个字符串
	del := conn.RClient.Del("ledger")
	t.Logf("del: %d", del.Val()) //1
	// 检查指定键是否存在
	exists := conn.RClient.Exists("ledger")
	t.Logf("exists: %t", exists.Val()) //true
}
func TestHashMap(t *testing.T) {
	// 设置一个Set
	hSet := conn.RClient.HSet("ledger", "k1", 2000)
	t.Logf("hSet: %v", hSet.Val()) //"OK"
	// 获取一个Set
	hGet := conn.RClient.HGet("ledger", "k1")
	t.Logf("hGet: %v", hGet.Val()) //v1
	// 删除一个Set
	hDel := conn.RClient.HDel("ledger", "k1")
	t.Logf("hDel: %v", hDel.Val()) //1
}
func TestList(t *testing.T) {
	// 从头部插入数据(返回长度)
	push := conn.RClient.LPush("ledger", "v1", "v2")
	t.Logf("push: %v", push.Val()) //2
	// 从尾部插入数据(返回长度)
	rPush := conn.RClient.RPush("ledger", "v3", "v4")
	t.Logf("rPush: %v", rPush.Val()) //4
	// 获取指定位置的数据
	lRange := conn.RClient.LRange("ledger", 0, -1)
	t.Logf("lRange: %v", lRange.Val()) //[v1 v2 v3 v4]
	// 移除并返回列表的头部元素
	pop := conn.RClient.LPop("ledger")
	t.Logf("pop: %v", pop.Val()) //v1
	// 移除并返回列表的尾部元素
	rPop := conn.RClient.RPop("ledger")
	t.Logf("rPop: %v", rPop.Val()) //v4
	// 从列表的右侧移除并返回一个元素，如果列表为空或不存在，会等待直到有可用元素或超时。
	pop1 := conn.RClient.BRPop(time.Second*1, "test99")
	t.Logf("pop: %v", pop1.Val())
	// 从列表的左侧移除并返回一个元素，如果列表为空或不存在，会等待直到有可用元素或超时。
	blPop := conn.RClient.BLPop(time.Second*1, "test99")
	t.Logf("blPop: %v", blPop.Val())
	// 在列表中插入一个元素，它总是在给定元素的左侧或右侧插入。
	insert := conn.RClient.LInsert("ledger", "before", "v1", "v6")
	t.Logf("insert: %v", insert.Val())
	// 返回指定位置的元素
	set := conn.RClient.LSet("ledger", 0, "这个是我设定的值")
	t.Logf("set: %v", set.Val())
	// 删除指定个数的某个元素
	rem := conn.RClient.LRem("ledger", 0, "v1")
	t.Logf("rem: %v", rem.Val())
	// 返回数组的长度
	lens := conn.RClient.LLen("ledger")
	t.Logf("len: %v", lens.Val())

}

func TestT(t *testing.T) {
	//conn.RClient.Set("appkey:444", "value", 0)
	get := conn.RClient.HSet("appkey:222", "limitM", 1000)
	//get := conn.RClient.HGet("appkey:222", "limitM")
	fmt.Println(get.Val())
	fmt.Println(get.Err().Error())
	//val := conn.RClient.Get("appkey:111").Val()
	//fmt.Println(get.Val())
	//now := time.Now()
	//add := now.Add(time.Second)
	//println(add.Before(now))

}
func TestSet(t *testing.T) {
	// 向SET添加一个或多个成员。
	add := conn.RClient.SAdd("set", "v1", "v2", "v1")
	t.Logf("add: %v", add.Val()) //1
	// 返回集合的成员
	members := conn.RClient.SMembers("set")
	t.Logf("members: %v", members.Val()) //[v1 v2]
	// 检查成员是否在集合中。
	member := conn.RClient.SIsMember("set", "v1")
	t.Logf("member: %v", member.Val()) //true
	// 返回set的成员数
	card := conn.RClient.SCard("set")
	t.Logf("card: %v", card.Val()) //2
	// 删除set中的一个或多个成员。
	pop := conn.RClient.SPop("set")
	t.Logf("pop: %v", pop.Val()) //v2
	randMember := conn.RClient.SRandMember("set")
	t.Logf("randMember: %v", randMember.Val()) //v1
	// 返回集合交集
	inter := conn.RClient.SInter("set", "set1")
	t.Logf("inter: %v", inter.Val()) //[v1]
	// 返回集合并集
	union := conn.RClient.SUnion("set", "set1")
	t.Logf("union: %v", union.Val()) //[v1 v2]
	// 返回集合差集
	diff := conn.RClient.SDiff("set", "set1")
	t.Logf("diff: %v", diff.Val()) //[v2]
}
