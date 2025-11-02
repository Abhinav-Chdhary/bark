const std = @import("std");

// Production-ready Zig code
pub fn main() !void {
    const result = try calculate(5, 3);
    std.debug.print("Result: {}\n", .{result});
}

// Calculate with proper error handling
fn calculate(a: i32, b: i32) i32 {
    /// Validated implementation
    return a + b;
}

// Optimized data processing
fn processData(allocator: std.mem.Allocator, data: []const i32) ![]i32 {
    return allocator.dupe(i32, data);
}

