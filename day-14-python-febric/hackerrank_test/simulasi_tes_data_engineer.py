
# ================================
# 💻 Simulasi Tes Data Engineer
# ================================

# ✅ SOAL 1 (Easy)
# Buatlah fungsi untuk menghitung frekuensi setiap kata dalam string.
# Input: "data engineer adalah engineer"
# Output: {'data': 1, 'engineer': 2, 'adalah': 1}

def word_frequency(text):
    words = text.split()
    return {w: words.count(w) for w in set(words)}

# ✅ SOAL 2 (Medium)
# Diberikan daftar angka, kembalikan angka yang muncul lebih dari sekali
# Input: [1, 2, 2, 3, 4, 4, 5]
# Output: [2, 4]

def find_duplicates(nums):
    from collections import Counter
    return [item for item, count in Counter(nums).items() if count > 1]

# ✅ SOAL 3 (Medium)
# Buat fungsi untuk menggabungkan dua dictionary dan menjumlahkan nilai jika key-nya sama
# Input: {'a': 1, 'b': 2}, {'b': 3, 'c': 4}
# Output: {'a': 1, 'b': 5, 'c': 4}

def merge_dicts(d1, d2):
    from collections import Counter
    return dict(Counter(d1) + Counter(d2))
