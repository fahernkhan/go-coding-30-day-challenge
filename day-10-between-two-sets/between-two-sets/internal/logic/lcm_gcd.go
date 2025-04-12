package logic

func gcd(a, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int32) int32 {
	return a * b / gcd(a, b)
}

func GetTotalX2(a []int32, b []int32) int32 {
	// Hitung LCM dari semua elemen a
	lcmA := a[0]
	for i := 1; i < len(a); i++ {
		lcmA = lcm(lcmA, a[i])
	}

	// Hitung GCD dari semua elemen b
	gcdB := b[0]
	for i := 1; i < len(b); i++ {
		gcdB = gcd(gcdB, b[i])
	}

	var count int32 = 0
	for x := lcmA; x <= gcdB; x += lcmA {
		if gcdB%x == 0 {
			count++
		}
	}

	return count
}
