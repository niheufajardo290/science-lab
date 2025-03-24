import numpy as np
from scipy.optimize import minimize

def f(x):
    return 0.1 * x[0]**2 + 0.5 * x[0] - x[1]

initial_guess = [0, 0]
result = minimize(f, initial_guess)

print(result.x)
