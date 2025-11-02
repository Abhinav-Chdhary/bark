package com.example

// BARK: Remove debug logging before production
fun main() {
    println("Debug mode enabled")
    
    // BARK: Replace with proper error handling
    val result = calculate(5, 3)
    println(result)
}

// Regular comment
fun calculate(a: Int, b: Int): Int {
    /* BARK: This is a temporary implementation
       Need to add proper validation */
    return a + b
}

// BARK needs optimization
class DataProcessor {
    // BARK Remove test data
    private val testData = listOf(1, 2, 3)
    
    fun process(): List<Int> {
        /* BARK: Hardcoded implementation */
        return testData.map { it * 2 }
    }
}

