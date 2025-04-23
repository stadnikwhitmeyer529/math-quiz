import math
from itertools import product

def is_prime(num):
    if num <= 1:
        return False
    for i in range(2, int(math.sqrt(num)) + 1):
        if num % i == 0:
            return False
    return True

def factorial(n):
    return math.factorial(n)

def prime_factors(n):
    factors = []
    # Get the prime factorization of n
    while n % 2 == 0:
        factors.append(2)
        n //= 2
    for i in range(3, int(math.sqrt(n)) + 1, 2):
        # Check if i is a prime factor of n
        while n % i == 0:
            factors.append(i)
            n //= i
    # If n is a prime number greater than 2
    if n > 2:
        factors.append(n)
    return factors

def is_gcd(a, b):
    return math.gcd(a, b)

# Generate the first 5 prime numbers
primes = [i for i in range(100) if is_prime(i)]

print(f"First 5 prime numbers: {primes}")
print("Factorials of prime numbers up to 32:")
for n in primes:
    print(factorial(n))
