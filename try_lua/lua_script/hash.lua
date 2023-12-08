-- 设置哈希值
-- 获取所有哈希字段及其值
local allFieldsAndValues = redis.call("HGETALL", "myHash")

for i, v in ipairs(allFieldsAndValues) do
    if i % 2 == 0 then
        print("key:", v)
    else
        print("value:", v)
    end
end

return allFieldsAndValues