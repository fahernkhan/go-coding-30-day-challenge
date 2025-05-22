# fabric_data_engineering_tutorial.py

import pandas as pd
import pyarrow as pa
import pyarrow.parquet as pq
from datetime import datetime
import os

# 1. Simulasi Data - Membuat DataFrame dummy
def create_sample_dataframe():
    data = {
        'customer_id': [101, 102, 103, 104],
        'name': ['Ali', 'Budi', 'Citra', 'Dewi'],
        'signup_date': pd.to_datetime(['2023-01-01', '2023-01-03', '2023-02-10', '2023-03-25']),
        'total_purchase': [200000, 150000, 320000, 270000],
        'is_active': [True, True, False, True]
    }
    df = pd.DataFrame(data)
    return df

# 2. Simpan DataFrame sebagai Parquet (biasa digunakan di Lakehouse)
def save_to_parquet(df, filename='customers.parquet'):
    table = pa.Table.from_pandas(df)
    pq.write_table(table, filename)
    print(f"✅ Data berhasil disimpan sebagai {filename}")

# 3. Baca kembali file Parquet
def read_from_parquet(filename='customers.parquet'):
    table = pq.read_table(filename)
    df = table.to_pandas()
    print(f"📦 Data berhasil dibaca kembali dari {filename}")
    return df

# 4. Tambah kolom baru dan transformasi
def enrich_dataframe(df):
    df['year'] = df['signup_date'].dt.year
    df['loyalty_tier'] = pd.cut(
        df['total_purchase'],
        bins=[0, 200000, 300000, float('inf')],
        labels=['Bronze', 'Silver', 'Gold']
    )
    return df

# 5. Simpan ulang hasil transformasi
def save_transformed_data(df, output_file='customers_enriched.parquet'):
    table = pa.Table.from_pandas(df)
    pq.write_table(table, output_file)
    print(f"🚀 Data transformasi disimpan sebagai {output_file}")

if __name__ == "__main__":
    print("🏗️ Belajar Data Engineering Microsoft Fabric")

    # Pipeline sederhana
    df_raw = create_sample_dataframe()
    save_to_parquet(df_raw)

    df_loaded = read_from_parquet()
    df_enriched = enrich_dataframe(df_loaded)
    print("\n📊 Data Setelah Transformasi:")
    print(df_enriched)

    save_transformed_data(df_enriched)
