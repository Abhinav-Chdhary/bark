#!/usr/bin/env python3


# BARK: Remove debug prints before merging
# BARK
def calculate_sum(a, b):
    """Calculate the sum of two numbers."""
    print(f"Debug: a={a}, b={b}")  # BARK: Remove debug statement
    return a + b


def main():
    # Regular comment
    result = calculate_sum(5, 3)
    # BARK: Replace with proper logging
    # BARK! plain marker with exclamation
    print(result)


if __name__ == "__main__":
    main()
