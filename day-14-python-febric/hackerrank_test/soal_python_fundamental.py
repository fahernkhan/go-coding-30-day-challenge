
# ======================================
# ✅ Kumpulan Soal Python Fundamental
# ======================================

# SOAL 1: Fungsi menukar dua angka
def swap(a, b):
    return b, a

# SOAL 2: Mengecek apakah sebuah string adalah palindrome
def is_palindrome(s):
    return s == s[::-1]

# SOAL 3: Menjumlahkan semua elemen dalam list
def sum_list(lst):
    return sum(lst)

# SOAL 4: Cari nilai maksimum kedua dalam list
def second_max(lst):
    return sorted(set(lst))[-2]

# SOAL 5: Hitung jumlah karakter unik dalam string
def unique_chars(s):
    return len(set(s))
