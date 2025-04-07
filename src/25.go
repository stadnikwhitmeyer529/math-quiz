def is_perfect_square(x):
    """
    Check if x is a perfect square.
    
    A perfect square is an integer that can be expressed as the product of another integer with itself.
    For example, 4 is a perfect square since it's the square of 2 (2 * 2 = 4).
    
    Parameters:
    x (int): The number to check.
    
    Returns:
    bool: True if x is a perfect square, False otherwise.
    """
    s = int(x ** 0.5)  # Calculate the integer part of the square root
    return s * s == x

# Test cases
print(is_perfect_square(16))  # Should print: True (16 is a perfect square)
print(is_perfect_square(14))  # Should print: False (14 is not a perfect square)
