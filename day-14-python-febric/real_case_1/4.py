def gainMaxValue(security_val, k):
    max_total = -float('inf')
    for m in range(k):
        group = []
        idx = m
        while idx < len(security_val):
            group.append(security_val[idx])
            idx += k
        current_sum = 0
        max_suffix = -float('inf')
        for num in reversed(group):
            current_sum += num
            if current_sum > max_suffix:
                max_suffix = current_sum
        if max_suffix > max_total:
            max_total = max_suffix
    return max_total