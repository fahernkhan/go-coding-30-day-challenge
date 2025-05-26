#!/bin/python3

import os # Modul os untuk berinteraksi dengan sistem operasi, khususnya untuk path file
import sys # Modul sys untuk berinteraksi dengan interpreter Python, tidak terlalu digunakan di sini

#
# Fungsi untuk memeriksa keseimbangan tanda kurung dalam array string.
#
# Fungsi ini diharapkan mengembalikan sebuah array string (STRING_ARRAY)
# yang berisi 'YES' atau 'NO' untuk setiap string input.
#
# Fungsi ini menerima parameter:
#   braces: STRING_ARRAY - sebuah array dari string yang berisi tanda kurung.
#
def matchingBraces(braces):
    # 'hasil' akan menyimpan daftar string 'YES' atau 'NO' untuk setiap input
    hasil = []
    
    # 'pasangan' adalah kamus (dictionary) yang memetakan tanda kurung penutup
    # ke tanda kurung pembuka yang sesuai. Ini memudahkan pencocokan.
    # Contoh: ')' cocok dengan '(', '}' dengan '{', dst.
    pasangan = {
        ')': '(', 
        '}': '{', 
        ']': '['
    }
    
    # Iterasi (loop) melalui setiap string yang ada di dalam array 'braces'
    for s in braces:
        # 'stack' adalah list yang akan digunakan sebagai tumpukan (stack).
        # Kita akan menyimpan tanda kurung pembuka di sini.
        # Konsep stack adalah LIFO (Last-In, First-Out) - yang terakhir masuk, pertama keluar.
        stack = []
        
        # 'seimbang' adalah variabel boolean yang melacak apakah string saat ini
        # masih dianggap seimbang. Kita asumsikan seimbang di awal.
        seimbang = True
        
        # Iterasi melalui setiap karakter dalam string 's' saat ini
        for karakter in s:
            # Jika karakter adalah tanda kurung pembuka ('(', '{', '['):
            if karakter in '({[':
                # Masukkan (push) karakter ini ke dalam stack.
                # Ini berarti kita menemukan tanda kurung pembuka yang perlu ditutup nanti.
                stack.append(karakter)
            # Jika karakter adalah tanda kurung penutup (')', '}', ']'):
            elif karakter in ')}]':
                # Ada dua kondisi yang membuat string tidak seimbang saat menemukan tanda kurung penutup:
                # 1. 'not stack': Stack kosong. Ini berarti kita menemukan tanda kurung penutup
                #    tanpa tanda kurung pembuka yang sesuai sebelumnya. Contoh: `)(`
                # 2. 'stack[-1] != pasangan[karakter]': Tanda kurung pembuka di bagian atas stack
                #    (stack[-1] adalah elemen terakhir yang ditambahkan) tidak cocok
                #    dengan tanda kurung penutup saat ini berdasarkan kamus 'pasangan'.
                #    Contoh: `([)]` - stack berisi `([`, lalu ditemukan `]`, tapi di stack teratas adalah `[` bukan `(`.
                if not stack or stack[-1] != pasangan[karakter]:
                    # Jika salah satu kondisi di atas terpenuhi, string tidak seimbang.
                    seimbang = False
                    # Kita bisa berhenti memeriksa string ini karena sudah dipastikan tidak seimbang.
                    break 
                else:
                    # Jika stack tidak kosong DAN tanda kurung di atas stack cocok,
                    # itu berarti kita berhasil menemukan pasangan untuk tanda kurung pembuka.
                    # Keluarkan (pop) tanda kurung pembuka dari stack.
                    stack.pop()
        
        # Setelah selesai memeriksa semua karakter dalam string:
        # Kita perlu memastikan bahwa stack kosong. Jika stack tidak kosong,
        # itu berarti ada tanda kurung pembuka yang tidak pernah ditutup.
        # Contoh: `({)` atau `[`
        if stack:
            seimbang = False
        
        # Berdasarkan nilai akhir dari 'seimbang', tambahkan 'YES' atau 'NO' ke daftar hasil.
        if seimbang:
            hasil.append("YES")
        else:
            hasil.append("NO")
    
    # Kembalikan daftar hasil akhir
    return hasil

# Bagian di bawah ini adalah kode standar untuk membaca input dan menulis output
# sesuai dengan format yang biasa digunakan di platform coding competitive (seperti HackerRank).
if __name__ == '__main__':
    # Membuka file output berdasarkan variabel lingkungan 'OUTPUT_PATH'.
    # Ini adalah cara platform mengekspektasikan Anda menulis hasil.
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    # Membaca jumlah string yang akan diuji (contoh: 2)
    braces_count = int(input().strip())

    # Membuat list kosong untuk menyimpan string-string tanda kurung yang akan diuji
    braces = []

    # Loop untuk membaca setiap string tanda kurung dari input
    for _ in range(braces_count):
        braces_item = input() # Membaca satu baris input
        braces.append(braces_item) # Menambahkan string ke list 'braces'

    # Memanggil fungsi 'matchingBraces' dengan list string yang sudah dibaca
    result = matchingBraces(braces)

    # Menulis setiap hasil (YES/NO) ke file output.
    # '\n'.join(result) akan menggabungkan semua elemen di 'result' dengan baris baru di antaranya.
    fptr.write('\n'.join(result))
    fptr.write('\n') # Menambahkan baris baru di akhir file, sesuai kebiasaan

    # Menutup file output setelah semua selesai ditulis
    fptr.close()