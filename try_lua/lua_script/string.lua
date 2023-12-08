-- 简单设置一个值
redis.call("SET", KEYS[1], ARGV[1])
-- 不存在就设置
redis.call("SET", KEYS[1], ARGV[1], "NX")
-- 存在就设置
redis.call("SET", KEYS[2], ARGV[2], "XX")
-- 设置过期时间(秒)
redis.call("SET", KEYS[3], ARGV[3], "EX", tonumber(ARGV[3]))
-- 设置过期时间(毫秒)
redis.call("SET", KEYS[4], ARGV[4], "PX", tonumber(ARGV[3]))
-- 设置值，并保留当前的过期时间
redis.call("SET", KEYS[1], ARGV[1], "KEEPTTL")
