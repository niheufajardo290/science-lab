import numpy as np
from scipy import stats

def calculate_statistic(data):
    """
    This function takes a dataset 'data' and calculates its statistical properties.
    
    Parameters:
    - data: A list or numpy array of numbers.

    Returns:
    - mean: The mean value of the dataset.
    - std_dev: The standard deviation of the dataset.
    - min_val: The minimum value in the dataset.
    - max_val: The maximum value in the dataset.
    """
    mean = np.mean(data)
    std_dev = np.std(data, ddof=1)  # Using ddof=1 to calculate unbiased sample variance
    min_val = np.min(data)
    max_val = np.max(data)

    return mean, std_dev, min_val, max_val

# Example usage:
data_points = [10, 20, 30, 40, 50]
mean, std_dev, min_val, max_val = calculate_statistic(data_points)
print(f"Mean: {mean}, Standard Deviation: {std_dev}, Minimum Value: {min_val}, Maximum Value: {max_val}")
