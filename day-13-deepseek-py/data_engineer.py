"""
data_pipeline.py
Script serbaguna untuk tugas Data Engineering harian dengan fitur:
- ETL Pipeline
- Data Validation
- Logging
- Error Handling
- Database Interaction
- Cloud Storage Integration
- Configuration Management
"""

# Import libraries umum
import os
import sys
import logging
import datetime as dt
from time import time
from typing import Dict, Any, Optional
from functools import wraps

# Import libraries data processing
import pandas as pd
import numpy as np
from pyspark.sql import SparkSession
from sqlalchemy import create_engine, exc

# Import cloud SDKs
import boto3
from google.cloud import storage

# Import data validation
from great_expectations import Dataset, GreatExpectations

# Konfigurasi
class Config:
    # Database
    DB_HOST = os.getenv('DB_HOST', 'localhost')
    DB_PORT = os.getenv('DB_PORT', '5432')
    DB_NAME = os.getenv('DB_NAME', 'data_warehouse')
    DB_USER = os.getenv('DB_USER', 'de_user')
    DB_PASSWORD = os.getenv('DB_PASSWORD', 'password')
    
    # AWS
    AWS_ACCESS_KEY = os.getenv('AWS_ACCESS_KEY')
    AWS_SECRET_KEY = os.getenv('AWS_SECRET_KEY')
    S3_BUCKET = os.getenv('S3_BUCKET', 'data-lake-bucket')
    
    # GCP
    GCP_PROJECT = os.getenv('GCP_PROJECT')
    GCS_BUCKET = os.getenv('GCS_BUCKET')
    
    # Spark
    SPARK_MASTER = os.getenv('SPARK_MASTER', 'local[*]')
    
    # Logging
    LOG_LEVEL = os.getenv('LOG_LEVEL', 'INFO').upper()

# Setup logging
logging.basicConfig(
    level=Config.LOG_LEVEL,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
    handlers=[
        logging.FileHandler('data_pipeline.log'),
        logging.StreamHandler()
    ]
)
logger = logging.getLogger(__name__)

# Custom Exceptions
class DataValidationError(Exception):
    """Exception for data validation failures"""
    pass

class PipelineExecutionError(Exception):
    """Exception for pipeline execution failures"""
    pass

# Decorators
def timer(func):
    """Decorator untuk mengukur waktu eksekusi fungsi"""
    @wraps(func)
    def wrapper(*args, **kwargs):
        start_time = time()
        result = func(*args, **kwargs)
        end_time = time()
        logger.info(f"Function {func.__name__} executed in {end_time - start_time:.2f} seconds")
        return result
    return wrapper

# Utility Functions
def create_spark_session(app_name: str = "DataPipeline") -> SparkSession:
    """Membuat Spark session dengan konfigurasi optimal untuk ETL"""
    return SparkSession.builder \
        .master(Config.SPARK_MASTER) \
        .appName(app_name) \
        .config("spark.sql.shuffle.partitions", "8") \
        .config("spark.driver.memory", "4g") \
        .config("spark.executor.memory", "8g") \
        .getOrCreate()

def get_db_connection(db_type: str = 'postgresql'):
    """Membuat koneksi database menggunakan SQLAlchemy"""
    engine = create_engine(
        f"{db_type}://{Config.DB_USER}:{Config.DB_PASSWORD}@{Config.DB_HOST}:{Config.DB_PORT}/{Config.DB_NAME}"
    )
    return engine.connect()

# Data Validation
def validate_dataframe(df: pd.DataFrame, expectations: Dict[str, Any]) -> bool:
    """Validasi dataframe menggunakan Great Expectations"""
    ge_df = Dataset.from_pandas(df)
    for expectation_type, params in expectations.items():
        getattr(ge_df, expectation_type)(**params)
    
    validation_result = ge_df.validate()
    if not validation_result["success"]:
        logger.error("Data validation failed:")
        for error in validation_result["results"]:
            if not error["success"]:
                logger.error(f" - {error['expectation_config']['expectation_type']}")
        raise DataValidationError("Data validation failed")
    return True

# ETL Components
class DataPipeline:
    def __init__(self):
        self.spark = create_spark_session()
        self.engine = get_db_connection()
        self.s3 = boto3.client(
            's3',
            aws_access_key_id=Config.AWS_ACCESS_KEY,
            aws_secret_access_key=Config.AWS_SECRET_KEY
        )
        
    @timer
    def extract_from_s3(self, key: str) -> pd.DataFrame:
        """Ekstrak data dari S3"""
        try:
            obj = self.s3.get_object(Bucket=Config.S3_BUCKET, Key=key)
            df = pd.read_parquet(obj['Body'])
            logger.info(f"Successfully extracted {key} from S3")
            return df
        except Exception as e:
            logger.error(f"Error extracting data from S3: {str(e)}")
            raise PipelineExecutionError("S3 extraction failed")

    @timer
    def transform_data(self, df: pd.DataFrame) -> pd.DataFrame:
        """Transformasi data dasar"""
        # Contoh transformasi
        df['created_at'] = pd.to_datetime(df['created_at'])
        df = df.drop_duplicates()
        df = df.dropna(subset=['user_id'])
        return df

    @timer
    def load_to_dwh(self, df: pd.DataFrame, table_name: str):
        """Load data ke data warehouse"""
        try:
            df.to_sql(
                name=table_name,
                con=self.engine,
                if_exists='append',
                index=False
            )
            logger.info(f"Successfully loaded data to {table_name}")
        except exc.SQLAlchemyError as e:
            logger.error(f"Database error: {str(e)}")
            raise PipelineExecutionError("Database load failed")

    def run_pipeline(self, s3_key: str, table_name: str):
        """Jalankan full pipeline"""
        try:
            # ETL Process
            raw_data = self.extract_from_s3(s3_key)
            transformed_data = self.transform_data(raw_data)
            
            # Data Validation
            validate_dataframe(transformed_data, {
                "expect_column_to_exist": {"column": "user_id"},
                "expect_column_values_to_not_be_null": {"column": "user_id"}
            })
            
            self.load_to_dwh(transformed_data, table_name)
            
        except Exception as e:
            logger.error(f"Pipeline failed: {str(e)}")
            raise

# Contoh penggunaan
if __name__ == "__main__":
    pipeline = DataPipeline()
    
    try:
        pipeline.run_pipeline(
            s3_key="raw_data/user_activity/2023-10-01.parquet",
            table_name="user_activity"
        )
    except PipelineExecutionError:
        sys.exit(1)
    except Exception as e:
        logger.error(f"Unexpected error: {str(e)}")
        sys.exit(1)