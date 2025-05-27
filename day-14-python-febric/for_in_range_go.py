my_list = ['merah', 'hijau', 'biru', 'kuning']

# 1. Penggunaan Dasar (indeks mulai dari 0)
print("--- Indeks default (mulai dari 0) ---")
for index, value in enumerate(my_list):
    print(f"Index {index}: {value}")
# Output:
# Index 0: merah
# Index 1: hijau
# Index 2: biru
# Index 3: kuning

# 2. Memulai Indeks dari Angka Lain (misalnya dari 1)
print("\n--- Indeks mulai dari 1 ---")
for num, color in enumerate(my_list, start=1):
    print(f"Warna ke-{num}: {color}")
# Output:
# Warna ke-1: merah
# Warna ke-2: hijau
# Warna ke-3: biru
# Warna ke-4: kuning

# 3. Menggunakan dengan Tipe Iterable Lain (contoh: string)
print("\n--- Dengan String ---")
my_string = "PYTHON"
for i, char in enumerate(my_string):
    print(f"Karakter ke-{i}: {char}")
# Output:
# Karakter ke-0: P
# Karakter ke-1: Y
# Karakter ke-2: T
# Karakter ke-3: H
# Karakter ke-4: O
# Karakter ke-5: N

# 4. Mengabaikan Indeks atau Nilai (dengan underscore _)
print("\n--- Mengabaikan Indeks (jika hanya butuh nilai) ---")
for _, color in enumerate(my_list):
    print(f"Saya suka warna {color}")
# Output:
# Saya suka warna merah
# Saya suka warna hijau
# Saya suka warna biru
# Saya suka warna kuning