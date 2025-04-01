import math
from fractions import Fraction

def calculate_math_quiz():
    # Generate a random problem and its solution
    problem = "What is 0.5 + 1.2?"
    solution = "3.7"
    
    # Calculate the result using the calculated solution
    answer = Fraction(math.fsum([problem, solution]), math.pow(10, 9)).limit_denominator()
    
    return f"Answer: {answer}"

# Generate and display a random go code for practice
print(calculate_math_quiz())
