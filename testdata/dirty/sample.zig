const std = @import("std");

// BARK: Remove debug prints before release
pub fn main() !void {
    std.debug.print("Debug mode\n", .{});
    
    // BARK: Replace with proper error handling
    const result = calculate(5, 3);
    std.debug.print("{}\n", .{result});
}

// Regular comment
fn calculate(a: i32, b: i32) i32 {
    /// BARK: This is a temporary implementation
    return a + b;
}

// BARK needs optimization
fn processData() void {
    // Implementation
}

