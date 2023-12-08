-- 尝试使用 SET 命令设置键为 KEYS[1] 的值为 1，同时设置过期时间和不允许覆盖
local notexists = redis.call("SET", KEYS[1], 1, "NX", "EX", tonumber(ARGV[2]))

-- 如果 SET 成功，则说明键不存在，返回 1 表示设置成功
if notexists then
    return 1
end

-- 获取当前键的值，并将其转换为数字
local current = tonumber(redis.call("GET", KEYS[1]))

-- 如果当前值不存在，说明前面的 SET 操作失败，使用 INCR 命令设置新值为 1，并设置过期时间
if current == nil then
    local result = redis.call("INCR", KEYS[1])
    redis.call("EXPIRE", KEYS[1], tonumber(ARGV[2]))
    return result
end

-- 如果当前值大于等于设定的阈值，则返回 -1 表示已达到阈值
if current >= tonumber(ARGV[1]) then
    return -1
end

-- 如果当前值小于阈值，则使用 INCR 命令递增，并返回递增后的值
local result = redis.call("INCR", KEYS[1])
return result
