def matchingBraces(braces):
    hasil = []  # List untuk menyimpan hasil akhir ("YES" atau "NO")
    pasangan = {')': '(', '}': '{', ']': '['}  # Peta kurung tutup ke kurung buka

    for s in braces:  # Iterasi setiap string dalam list input
        stack = []     # Stack untuk menyimpan kurung buka
        seimbang = True  # Asumsikan string seimbang pada awalnya

        for karakter in s:
            if karakter in '({[':  # Jika karakter adalah kurung buka
                stack.append(karakter)
            elif karakter in ')}]':  # Jika karakter adalah kurung tutup
                # Jika stack kosong atau tidak cocok dengan pasangan terakhir, maka tidak seimbang
                if not stack or stack[-1] != pasangan[karakter]:
                    seimbang = False
                    break  # Hentikan pengecekan karena sudah tidak valid
                else:
                    stack.pop()  # Kurung cocok, hapus dari stack

        # Jika masih ada kurung buka tertinggal di stack, berarti tidak seimbang
        if stack:
            seimbang = False

        # Tambahkan hasil ke list berdasarkan nilai `seimbang`
        if seimbang:
            hasil.append("YES")
        else:
            hasil.append("NO")
    
    return hasil  # Kembalikan list hasil untuk semua input string

# Bagian utama program (main)
if __name__ == '__main__':
    t = int(input())  # Jumlah test case
    braces = [input() for _ in range(t)]  # List input dari user

    results = matchingBraces(braces)  # Jalankan fungsi utama

    # Cetak hasil satu per satu
    for res in results:
        print(res)
