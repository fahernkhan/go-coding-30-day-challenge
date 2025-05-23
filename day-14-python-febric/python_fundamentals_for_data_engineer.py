# ================================================
# 📘 Python Fundamentals for Data Engineer
# ================================================

# -------------------------
# 1. Variabel & Tipe Data
# -------------------------
name = "Fathur"
age = 23
is_engineer = True
skills = ["Python", "SQL", "Cloud"]
salary = 10_000_000

print(f"👤 {name}, usia {age}, gaji Rp{salary}")

# -------------------------
# 2. Struktur Data Dasar
# -------------------------
# List
languages = ["Python", "Java", "SQL"]
languages.append("Scala")

# Dictionary
employee = {
    "id": 101,
    "name": "Budi",
    "role": "Data Engineer"
}

# Tuple (immutable)
status_codes = (200, 404, 500)

# Set (unique)
unique_skills = set(["SQL", "SQL", "Python"])

# -------------------------
# 3. Control Flow
# -------------------------
for lang in languages:
    if lang == "Python":
        print("🔥 Favorite Language")
    else:
        print(f"Language: {lang}")

# -------------------------
# 4. Function
# -------------------------
def calculate_bonus(salary, percentage=10):
    return salary * (percentage / 100)

bonus = calculate_bonus(salary)
print("Bonus:", bonus)

# -------------------------
# 5. File Handling
# -------------------------
with open("output.txt", "w") as f:
    f.write("This is a log file for Data Engineer")

# -------------------------
# 6. Regex
# -------------------------
import re
text = "Email saya fathur@example.com"
email = re.findall(r'\b[\w.-]+?@\w+?\.\w+?\b', text)
print("📧 Email ditemukan:", email)

# -------------------------
# 7. JSON & API
# -------------------------
import json
data = {
    "name": "Data Engineer",
    "level": "Entry",
    "skills": skills
}
json_str = json.dumps(data, indent=2)
print("📦 JSON:", json_str)

# -------------------------
# 8. DateTime
# -------------------------
from datetime import datetime, timedelta
now = datetime.now()
yesterday = now - timedelta(days=1)
print("Today:", now.strftime("%Y-%m-%d"), "| Yesterday:", yesterday.strftime("%Y-%m-%d"))

# -------------------------
# 9. SQLite (Simulasi SQL)
# -------------------------
import sqlite3

conn = sqlite3.connect(":memory:")
cursor = conn.cursor()
cursor.execute("CREATE TABLE customers (id INTEGER, name TEXT, signup_date DATE)")
cursor.execute("INSERT INTO customers VALUES (1, 'Ali', '2024-01-01')")
cursor.execute("SELECT * FROM customers")
print("📊 SQL Result:", cursor.fetchall())
conn.close()

# -------------------------
# 10. Pandas for ETL
# -------------------------
import pandas as pd

df = pd.DataFrame({
    'customer_id': [1, 2],
    'name': ['Andi', 'Budi'],
    'amount': [200000, 300000]
})
df['amount_million'] = df['amount'] / 1_000_000
print("🧺 DataFrame:\n", df)

# -------------------------
# 11. Exception Handling
# -------------------------
try:
    x = 1 / 0
except ZeroDivisionError as e:
    print("⚠️ Error:", e)

# -------------------------
# 12. Logging
# -------------------------
import logging
logging.basicConfig(level=logging.INFO)
logging.info("🚀 ETL process started")

# -------------------------
# 13. Lambda, Map, Filter, Reduce
# -------------------------
from functools import reduce

numbers = [1, 2, 3, 4, 5]
squares = list(map(lambda x: x**2, numbers))
even = list(filter(lambda x: x % 2 == 0, numbers))
summed = reduce(lambda x, y: x + y, numbers)
print("Squares:", squares, "| Even:", even, "| Sum:", summed)

# -------------------------
# 14. OOP Dasar
# -------------------------
class DataEngineer:
    def __init__(self, name, skill):
        self.name = name
        self.skill = skill

    def introduce(self):
        print(f"Hi, I'm {self.name}, skilled in {', '.join(self.skill)}")

de = DataEngineer("Nanda", ["Python", "DBT", "Cloud"])
de.introduce()

# -------------------------
# 15. OS & Environment
# -------------------------
import os
os.environ["API_KEY"] = "123456"
print("🔐 API_KEY:", os.getenv("API_KEY"))

# -------------------------
# 16. Modularisasi (Simulasi)
# -------------------------
def run_etl():
    logging.info("ETL Started")
    df = pd.read_csv("your_data.csv") if os.path.exists("your_data.csv") else pd.DataFrame()
    logging.info("ETL Completed")

# Uncomment saat ada file
# run_etl()
