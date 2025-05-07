# data_engineer_fundamental.py

import pandas as pd
import numpy as np

# 1. Membaca file CSV
def load_data(filepath):
    print(f"Loading data from {filepath}...")
    return pd.read_csv(filepath)

# 2. Menampilkan info dasar
def explore_data(df):
    print("\n--- Head of the Data ---")
    print(df.head())
    print("\n--- Info ---")
    print(df.info())
    print("\n--- Descriptive Stats ---")
    print(df.describe())

# 3. Membersihkan data
def clean_data(df):
    print("\nCleaning data...")
    # Contoh: Menghapus duplikat
    df = df.drop_duplicates()
    # Mengisi missing value pada kolom numerik
    for col in df.select_dtypes(include=np.number).columns:
        df[col] = df[col].fillna(df[col].median())
    # Mengisi missing value pada kolom object (string)
    for col in df.select_dtypes(include='object').columns:
        df[col] = df[col].fillna('unknown')
    return df

# 4. Transformasi data sederhana
def transform_data(df):
    print("\nTransforming data...")
    if 'amount' in df.columns:
        df['amount_usd'] = df['amount'] * 1.1  # Contoh konversi nilai
    return df

# 5. Menyimpan hasil
def save_data(df, output_path):
    print(f"\nSaving cleaned data to {output_path}...")
    df.to_csv(output_path, index=False)

# 6. Pipeline utama
def main():
    input_path = 'data/raw_data.csv'     # Ubah path sesuai kebutuhan
    output_path = 'data/cleaned_data.csv'

    df = load_data(input_path)
    explore_data(df)
    df = clean_data(df)
    df = transform_data(df)
    save_data(df, output_path)
    print("\nPipeline selesai.")

if __name__ == '__main__':
    main()
