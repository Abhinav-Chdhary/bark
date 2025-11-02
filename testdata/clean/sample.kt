package com.example

// Production-ready Kotlin code
fun main() {
    val result = calculate(5, 3)
    println("Result: $result")
}

// Calculate sum with validation
fun calculate(a: Int, b: Int): Int {
    /* Validated implementation with proper checks */
    require(a >= 0) { "First argument must be non-negative" }
    require(b >= 0) { "Second argument must be non-negative" }
    return a + b
}

// Optimized data processor
class DataProcessor(private val data: List<Int>) {
    
    fun process(): List<Int> {
        /* Efficient implementation using sequence */
        return data.asSequence()
            .filter { it > 0 }
            .map { it * 2 }
            .toList()
    }
}

