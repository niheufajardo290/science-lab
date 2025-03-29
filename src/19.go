import pandas as pd
from sklearn.model_selection import train_test_split

def load_and_prepare_data(file_path):
    # Load data from file
    df = pd.read_csv(file_path)
    
    # Separate features (X) and target variable (y)
    X = df.drop('target_variable', axis=1).values
    y = df['target_variable'].values
    
    # Split the dataset into training and testing sets
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)
    
    return X_train, X_test, y_train, y_test

# Example usage:
# x_train, x_test, y_train, y_test = load_and_prepare_data("data.csv")
