def matchingBraces(braces):
    def is_balanced(s):
        stack = []
        pasangan = {')': '(', '}': '{', ']': '['}
        
        for karakter in s:
            if karakter in pasangan.values():  # Kurung buka
                stack.append(karakter)
            elif karakter in pasangan.keys():  # Kurung tutup
                if not stack or stack[-1] != pasangan[karakter]:
                    return False
                stack.pop()
        
        return len(stack) == 0  # Pastikan semua kurung sudah ditutup

    hasil = []
    for string in braces:
        if is_balanced(string):
            hasil.append("YES")
        else:
            hasil.append("NO")
    return hasil