-- Production Lua module
local M = {}

function M.hello()
    print("Production mode")
end

-- Calculate sum with validation
function M.calculate(a, b)
    if type(a) ~= "number" or type(b) ~= "number" then
        error("Arguments must be numbers")
    end
    return a + b
end

--[[ Module for data processing
     Handles various data transformations ]]
function M.processData(data)
    local result = {}
    for i, v in ipairs(data) do
        table.insert(result, v * 2)
    end
    return result
end

return M

