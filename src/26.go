def add(x, y):
    while True:
        if not isinstance(x, int) or not isinstance(y, int):
            print("Both parameters must be integers!")
            break
        else:
            return x + y

def subtract(x, y):
    while True:
        if not isinstance(x, int) or not isinstance(y, int):
            print("Both parameters must be integers!")
            break
        else:
            return x - y

def multiply(x, y):
    while True:
        if not isinstance(x, int) or not isinstance(y, int):
            print("Both parameters must be integers!")
            break
        else:
            return x * y

def divide(x, y):
    while True:
        if not isinstance(x, int) or not isinstance(y, int):
            print("Both parameters must be integers!")
            break
        else:
            return x / y
