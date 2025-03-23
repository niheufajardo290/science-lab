import pandas as pd
from sklearn.model_selection import train_test_split

def load_data(file_path):
    # Load data from file
    df = pd.read_csv(file_path)
    X = df.drop('target_column', axis=1).values  # Assuming target column is in 'y' position (index 0)
    y = df['target_column'].values  # Assuming target column name is 'y'
    
    return X, y

def split_data(X, y):
    """
    Split the data into training and testing sets
    """
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)
    return X_train, X_test, y_train, y_test

def evaluate_model(model, X_train, X_test, y_train, y_test):
    """
    Evaluate the model on the test set
    """
    predicted_y = model.predict(X_test)
    
    # Calculate mean squared error (MSE) and root mean squared error (RMSE)
    mse = mean_squared_error(y_test, predicted_y)
    rmse = math.sqrt(mean_squared_error(y_test, predicted_y))
    
    print(f"MSE: {mse}")
    print(f"RMSE: {rmse}")

def main(file_path):
    X_train, y_train = load_data(file_path)
    X_test, y_test = split_data(X_train, y_train)
    evaluate_model(LinearRegression(), X_train, X_test, y_train, y_test)

if __name__ == "__main__":
    # Provide the file path where your data is stored
    main("path_to_your_file.csv")
