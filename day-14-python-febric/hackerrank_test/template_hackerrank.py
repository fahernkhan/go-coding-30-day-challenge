
# ===============================
# 📂 HackerRank Practice Template
# ===============================
# 🧪 Deskripsi Soal:
# Diberikan sebuah daftar angka, hitung jumlah angka genap di dalam daftar tersebut.

# 🔢 Contoh Input:
# [1, 2, 3, 4, 5, 6]

# ✅ Contoh Output:
# 3

def count_even_numbers(nums):
    # TODO: implementasikan fungsi ini
    return sum(1 for n in nums if n % 2 == 0)

# Tes
print(count_even_numbers([1, 2, 3, 4, 5, 6]))  # Output: 3
