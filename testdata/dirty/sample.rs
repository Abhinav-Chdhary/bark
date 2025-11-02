// BARK: Remove debug prints before release
fn main() {
    println!("Debug mode");
    
    // BARK: Replace with proper error handling
    let result = calculate(5, 3);
    println!("{}", result);
}

// Regular comment
fn calculate(a: i32, b: i32) -> i32 {
    /* BARK: This is a temporary implementation
       Need to add proper validation */
    a + b
}

// BARK needs optimization
fn process_data() {
    // Implementation
}

