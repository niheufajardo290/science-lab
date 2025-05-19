# Importing necessary library
import pandas as pd

# Creating DataFrame
data = {
    'Name': ['Alice', 'Bob', 'Charlie'],
    'Age': [20, 30, 40],
    'Gender': ['Male', 'Female', 'Male']
}
df = pd.DataFrame(data)

# Displaying the DataFrame
print(df)
