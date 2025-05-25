
# === 1. TIPE DATA DASAR ===
a = 10           # int
b = 3.14         # float
c = "data"       # string
d = True         # boolean

print(type(a), type(b), type(c), type(d))

# === 2. STRUKTUR DATA PYTHON ===
# List
data = [1, 2, 3]
data.append(4)
data.remove(2)
print("List:", data)

# Tuple
t = (1, 2, 3)
print("Tuple:", t[0])

# Dictionary
d = {'a': 1, 'b': 2}
d['c'] = 3
print("Dict:", d)

# Set
s = {1, 2, 3}
s.add(4)
s.discard(2)
print("Set:", s)

# === 3. PENGKONDISIAN & PERULANGAN ===
x = 10
if x > 5:
    print("Besar")
else:
    print("Kecil")

for i in range(3):
    print("For:", i)

i = 0
while i < 3:
    print("While:", i)
    i += 1

# === 4. FUNGSI ===
def tambah(a, b):
    return a + b

print("Fungsi:", tambah(3, 5))

# === 5. LIST COMPREHENSION ===
squares = [x**2 for x in range(5)]
print("Squares:", squares)

# === 6. STRING MANIPULASI ===
s = "Data Engineer"
print(s.lower())
print(s.split())
print("-".join(["data", "engineer"]))

# === 7. ERROR HANDLING ===
try:
    a = 1 / 0
except ZeroDivisionError:
    print("Tidak bisa bagi 0")

# === 8. BUILT-IN FUNCTIONS & LAMBDA ===
nums = [1, 2, 3, 4]
even = list(filter(lambda x: x % 2 == 0, nums))
doubled = list(map(lambda x: x * 2, nums))
print("Even:", even)
print("Doubled:", doubled)

# === 9. FILE HANDLING ===
# Contoh ini diasumsikan ada file data.txt (skip jika file tidak ada)
# with open("data.txt", "r") as f:
#     content = f.read()
#     print("File content:", content)

# === 10. COUNTER & COLLECTIONS ===
from collections import Counter
nums = [1, 2, 2, 3, 3, 3]
print("Counter:", Counter(nums))
