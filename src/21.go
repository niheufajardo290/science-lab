import numpy as np
from scipy.stats import ttest_ind

data = [10, 20, 30, 40, 50]
expected_data = data

def perform_ttest(data, expected_data):
    t_statistic, p_value = ttest_ind(data, expected_data)
    return t_statistic, p_value

t_statistic, p_value = perform_ttest([12.7, 23.8, 34.5], [10.5, 21.9, 36.0])
