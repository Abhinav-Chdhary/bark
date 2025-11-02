-- BARK: Remove debug prints
function hello()
    print("Debug mode")
    
    -- BARK: Replace with proper logging
    local x = 42
    print(x)
end

-- Regular comment
function calculate(a, b)
    -- BARK temporary implementation
    return a + b
end

--[[ BARK: This entire function needs review
     before production deployment ]]
function testFunction()
    return "test"
end

hello()

