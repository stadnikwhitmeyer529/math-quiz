import sympy as sp

def solve_quadratic_equation(a, b, c):
    try:
        # Calculate the discriminant
        discriminant = b**2 - 4*a*c
        
        # Find the two solutions using the quadratic formula
        x1 = (-b + sp.sqrt(discriminant)) / (2 * a)
        x2 = (-b - sp.sqrt(discriminant)) / (2 * a)
        
        return x1, x2
    except:
        return "Please provide valid coefficients."
