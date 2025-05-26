def gainMaxValue(security_val, k):
    # Inisialisasi 'max_total' dengan nilai negatif tak terhingga.
    # Ini adalah titik awal yang sangat kecil, memastikan bahwa setiap jumlah positif
    # atau bahkan jumlah negatif yang lebih besar dari ini akan langsung menggantikan 'max_total'.
    # Tujuannya adalah untuk menemukan nilai maksimum dari semua jalur yang mungkin.
    max_total = -float('inf')

    # Loop ini akan membuat 'k' grup atau 'jalur' yang berbeda.
    # Setiap grup dimulai dari indeks 'm' (dari 0 hingga k-1) dan melompat sejauh 'k' langkah.
    # Contoh: Jika k=2,
    #   - Saat m=0, grup akan berisi elemen di indeks 0, 2, 4, ...
    #   - Saat m=1, grup akan berisi elemen di indeks 1, 3, 5, ...
    for m in range(k):
        group = []  # List ini akan menyimpan nilai-nilai dari 'grup' atau 'jalur' saat ini.
        idx = m     # Indeks awal untuk mengumpulkan elemen di grup ini.

        # Loop ini mengumpulkan semua elemen yang termasuk dalam 'grup' saat ini.
        # Dimulai dari 'idx' dan terus melompat 'k' langkah,
        # sampai 'idx' mencapai atau melebihi panjang daftar 'security_val'.
        while idx < len(security_val):
            group.append(security_val[idx]) # Tambahkan nilai dari security_val[idx] ke 'group'.
            idx += k                        # Pindah ke indeks berikutnya dalam 'jalur' ini.

        # Setelah satu 'grup' terbentuk (misalnya, [3, -2, 9] untuk k=2, m=0),
        # kita akan mencari 'jumlah suffix maksimum' dalam 'grup' ini.
        # 'Jumlah suffix maksimum' adalah jumlah terbesar yang bisa didapat
        # dari sub-bagian 'grup' yang dimulai dari titik mana pun dan berakhir di akhir 'grup'.

        current_sum = 0  # Variabel ini akan menyimpan jumlah kumulatif saat kita bergerak mundur.
        # Inisialisasi 'max_suffix' dengan nilai negatif tak terhingga.
        # Ini memastikan bahwa jumlah suffix pertama yang dihitung (bahkan jika negatif)
        # akan menjadi nilai 'max_suffix' yang baru untuk grup ini.
        max_suffix = -float('inf')

        # Loop ini mengiterasi mundur melalui elemen-elemen 'group'.
        # Ini adalah cara efisien untuk menghitung jumlah suffix maksimum.
        for num in reversed(group): # 'reversed(group)' memproses elemen dari kanan ke kiri.
            current_sum += num      # Tambahkan elemen saat ini ke 'current_sum'.

            # Bandingkan 'current_sum' (jumlah suffix yang sedang dipertimbangkan)
            # dengan 'max_suffix' yang sudah ditemukan sejauh ini untuk grup ini.
            # Jika 'current_sum' lebih besar, perbarui 'max_suffix'.
            if current_sum > max_suffix:
                max_suffix = current_sum

        # Setelah semua elemen dalam 'grup' diproses dan 'max_suffix' terbaiknya ditemukan,
        # kita bandingkan 'max_suffix' ini dengan 'max_total' global.
        # 'max_total' menyimpan nilai maksimum yang ditemukan dari semua grup yang sudah diproses.
        if max_suffix > max_total:
            max_total = max_suffix # Jika 'max_suffix' dari grup ini lebih baik, perbarui 'max_total'.

    # Setelah semua 'k' grup telah diproses dan nilai 'max_suffix' terbaiknya
    # telah dibandingkan, 'max_total' akan berisi nilai maksimum keseluruhan
    # yang bisa didapatkan dari jalur mana pun.
    return max_total #kembalikan nilai maksimum dari semua jalur
