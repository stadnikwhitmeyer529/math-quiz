def add(a: int, b: int) -> int:
    return a + b

def subtract(a: int, b: int) -> int:
    return a - b

def multiply(a: int, b: int) -> int:
    return a * b

def divide(a: int, b: int) -> float:
    if b != 0:
        return a / b
    else:
        return "Cannot divide by zero."

# Example usage:
result = add(10, 5)
print(result)
